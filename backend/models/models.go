
package models

import (
	"database/sql"
)


type Product struct {
	Id          int32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	Images      []string `json:"images"`
	Sold 		bool `json:"sold"`
}

type Order struct {
	Id int32
	CreatedAt string
	LastName string
	FirstName string
	Email 	string
	ProductIds []int32
	TotalCost float64
	Paid bool
	PaymentId string
}

type CheckoutRequest struct {
	ProductIds []int32 `json:"productIds"`
	LastName string `json:"lastName"`
	FirstName string `json:"firstName"`
	Email string `json:"email"`
}

type CheckoutResponse struct {
	OrderID   int32   `json:"orderId"`
	TotalCost float64 `json:"total"`
	Message   string  `json:"message"`
}

type Migration struct {
	ID int32
	Name string
	Up func(db *sql.DB) error
	Down func(db *sql.DB) error
}
