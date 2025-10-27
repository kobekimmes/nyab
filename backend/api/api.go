package api

import (
	"os"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/kobekimmes/nyab/backend/db"
	"github.com/kobekimmes/nyab/backend/models"

	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func HandleCheckout(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var orderRequest models.CheckoutRequest
	if err := json.NewDecoder(req.Body).Decode(&orderRequest); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}	

	tx, err := db.Db.Begin()
	if err != nil {
		http.Error(res, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var total float64
	for _, id := range orderRequest.ProductIds {
		product, err := db.ReadProduct(id)
		if err != nil {
			http.Error(res, fmt.Sprintf("Product not found: %s", err.Error()), http.StatusBadRequest)
			return
		}
		if product.Sold {
			http.Error(res, "Product unavailable", http.StatusConflict)
			return
		}
		oldProduct := *product
		product.Sold = true

		if err := db.UpdateProduct(oldProduct, *product, tx); err != nil {
			http.Error(res, "Failed to update product", http.StatusInternalServerError)
			return
		}

		total += product.Price * (1 - product.Discount)
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	paymentIntentParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(total * 100)),
		Currency: stripe.String("usd"),
		ReceiptEmail: stripe.String(orderRequest.Email),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	intent, err := paymentintent.New(paymentIntentParams)
	if err != nil {
		http.Error(res, fmt.Sprintf("Payment failed: %v", err), http.StatusPaymentRequired)
		return
	}

	order := models.Order {
		ProductIds: orderRequest.ProductIds,
		TotalCost: total,
		LastName: orderRequest.LastName,
		FirstName: orderRequest.FirstName,
		Email: orderRequest.Email,
		Paid: true,
		PaymentId: intent.ID,
	}

	record, err := db.CreateOrder(order, tx)
	if err != nil {
		http.Error(res, "Failed to create order", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(res, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	orderResponse := models.CheckoutResponse {
		OrderID: record,
		TotalCost: total,
		Message: "Thank you for your purchase!",
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(orderResponse)
}

func HandleProducts(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }	

	products, err := db.GetProducts()
	if err != nil {
		http.Error(res, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(products)
}


