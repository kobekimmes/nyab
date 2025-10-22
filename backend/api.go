package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func handleCheckout(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var order Order
	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}	

	tx, err := db.Begin()
	if err != nil {
		http.Error(res, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var total float64
	for _, id := range order.ProductIds {
		product, err := readProduct(id)
		if err != nil {
			http.Error(res, fmt.Sprintf("Product not found: %s", err.Error()), http.StatusBadRequest)
			return
		}
		if product.Sold {
			http.Error(res, "Product unavailable", http.StatusConflict)
			return
		}
		old_product := *product
		product.Sold = true

		if err := updateProduct(*product, tx); err != nil {
			http.Error(res, "Failed to update product", http.StatusInternalServerError)
			return
		}

		if err := logUpdate(tx, "products", product.ID, "UPDATE", old_product, product); err != nil {
			http.Error(res, "Failed to log product update", http.StatusInternalServerError)
			return
		}

		total += product.Price * (1 - product.Discount)
	}

	order.TotalCost = total
	record, err := createOrder(order, tx)
	if err != nil {
		http.Error(res, "Failed to create order", http.StatusInternalServerError)
		return
	}

	if err := logUpdate(tx, "orders", record, "INSERT", nil, order); err != nil {
		http.Error(res, "Failed to log order", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(res, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(struct {
		OrderID int32 	`json:"order_id"`
		Total   float64 `json:"total"`
		Message string  `json:"message"`
	}{
		OrderID: order.ID,
		Total: order.TotalCost,
		Message: "",
	})
}

func handleProducts(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }	

	products, err := getProducts()
	if err != nil {
		http.Error(res, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(products)
}


