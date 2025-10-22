package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Product struct {
	ID          int32
	Name        string
	Description string
	Price       float64
	Discount    float64
	Images      []string
	Sold 		bool
}

type Order struct {
	ID int32
	CreatedAt string
	LastName string
	FirstName string
	Email 	string
	ProductIds []int32
	TotalCost float64
}


func initDB() {

	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, failed to initialize database connection")
    }

	host     := os.Getenv("PG_HOST")
    port     := os.Getenv("PG_PORT")
    user     := os.Getenv("PG_USER")
    password := os.Getenv("PG_PASSWORD")
    dbname   := os.Getenv("PG_DB")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL")
}

func logUpdate(transaction *sql.Tx, table string, record int32, action string, old_data interface{}, new_data interface{}) (error) {
	old, _ := json.Marshal(old_data)
	new, _ := json.Marshal(new_data)

	query := `
		INSERT INTO updates(table, record, action, old, new)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := transaction.Exec(query, table, record, action, old, new)
	return err
}

// Attempts to create a new product and store in the "products" table, returns id of new product
func createProduct(p Product, transaction *sql.Tx) (int32, error) {
	var id int32
	query := `
		INSERT INTO products (name, description, price, discount, images, sold)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	err := transaction.QueryRow(query, p.Name, p.Description, p.Price, p.Discount, pq.Array(p.Images), p.Sold).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createOrder(o Order, transaction *sql.Tx) (int32, error) {
	var id int32
	query := `
		INSERT INTO orders (total_cost, created_at, product_ids, first_name, last_name, email)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	err := transaction.QueryRow(query, o.TotalCost, o.CreatedAt, pq.Array(o.ProductIds), o.FirstName, o.LastName, o.Email).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func readProduct(id int32) (*Product, error) {
	var p Product
	var images pq.StringArray
	query := `
		SELECT id, name, description, price, discount, images, sold 
		FROM products
		WHERE id=$1`

	
	err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Discount, &images, &p.Sold)

	if err != nil {
		return nil, err
	}
	p.Images = images
	return &p, nil

}

func readOrder(id int32) (*Order, error) {
	var o Order
	var product_ids pq.Int32Array
	query := `
		SELECT id, product_ids, total_cost, first_name, last_name, email, created_at 
		FROM orders 
		WHERE id=$1`
	err := db.QueryRow(query, id).Scan(&o.ID, &product_ids, &o.TotalCost, &o.FirstName, &o.LastName, &o.Email, &o.CreatedAt)
	if err != nil {
		return nil, err
	}
	o.ProductIds = product_ids
	return &o, nil
}


func updateProduct(p Product, transaction *sql.Tx) (error) {
	query := `
		UPDATE products
		SET name=$1, description=$2, price=$3, discount=$4, images=$5, sold=$6
		WHERE id=$7`
	_, err := transaction.Exec(query, p.Name, p.Description, p.Price, p.Discount, pq.Array(p.Images), p.Sold, p.ID)
	return err
}

func updateOrder() {

}

func deleteProduct() {

}

func deleteOrder() {

}

func getProducts() ([]Product, error) {

	product_rows, err := db.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, err
	}
	defer product_rows.Close()

	var products []Product

	for product_rows.Next() {
		var p Product
		var images pq.StringArray

		if err := product_rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Discount, &images, &p.Sold); err != nil {
			return nil, err
		}
		p.Images = images
		products = append(products, p)

	}
	return products, nil
}

func getOrderHistory() {

}