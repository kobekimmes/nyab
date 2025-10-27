package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kobekimmes/nyab/backend/models"

	"github.com/lib/pq"
)

var Db *sql.DB

func DbInit() {

	host     := os.Getenv("PG_HOST")
    port     := os.Getenv("PG_PORT")
    user     := os.Getenv("PG_USER")
    password := os.Getenv("PG_PASSWORD")
    dbname   := os.Getenv("PG_DB")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL")
}

func logUpdate(transaction *sql.Tx, table string, record int32, action string, oldData any, newData any) (error) {
	old, _ := json.Marshal(oldData)
	new, _ := json.Marshal(newData)

	query := `
		INSERT INTO updates(table_name, record, action_method, old, new)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := transaction.Exec(query, table, record, action, old, new)
	return err
}

// Attempts to create a new product and store in the "products" table, returns id of new product
func CreateProduct(p models.Product, transaction *sql.Tx) (int32, error) {
	var id int32
	query := `
		INSERT INTO products (name, description, price, discount, images, sold)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	err := transaction.QueryRow(query, p.Name, p.Description, p.Price, p.Discount, pq.Array(p.Images), p.Sold).Scan(&id)
	if logErr := logUpdate(transaction, "products", id, "CREATE", nil, p); logErr != nil {
		return -1, logErr
	}
	
	if err != nil {
		return -1, err
	}
	return id, nil
}

func CreateOrder(o models.Order, transaction *sql.Tx) (int32, error) {
	var id int32
	query := `
		INSERT INTO orders (total_cost, created_at, product_ids, first_name, last_name, email, paid, payment_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	err := transaction.QueryRow(query, o.TotalCost, o.CreatedAt, pq.Array(o.ProductIds), o.FirstName, o.LastName, o.Email, o.Paid, o.PaymentId).Scan(&id)
	if log_err := logUpdate(transaction, "orders", id, "CREATE", nil, o); log_err != nil {
		return -1, log_err
	}
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ReadProduct(id int32) (*models.Product, error) {
	var p models.Product
	var images pq.StringArray
	query := `
		SELECT id, name, description, price, discount, images, sold 
		FROM products
		WHERE id=$1`

	
	err := Db.QueryRow(query, id).Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Discount, &images, &p.Sold)

	if err != nil {
		return nil, err
	}
	p.Images = images
	return &p, nil

}

func ReadOrder(id int32) (*models.Order, error) {
	var o models.Order
	var productIds pq.Int32Array
	query := `
		SELECT id, product_ids, total_cost, first_name, last_name, email, created_at 
		FROM orders 
		WHERE id=$1`
	err := Db.QueryRow(query, id).Scan(&o.Id, &productIds, &o.TotalCost, &o.FirstName, &o.LastName, &o.Email, &o.CreatedAt)
	if err != nil {
		return nil, err
	}
	o.ProductIds = productIds
	return &o, nil
}


func UpdateProduct(old models.Product, new models.Product, transaction *sql.Tx) (error) {
	if old.Id != new.Id {
		panic("Invalid operation")
	}
	query := `
		UPDATE products
		SET name=$1, description=$2, price=$3, discount=$4, images=$5, sold=$6
		WHERE id=$7`
	_, err := transaction.Exec(query, new.Name, new.Description, new.Price, new.Discount, pq.Array(new.Images), new.Sold, old.Id)
	if logErr := logUpdate(transaction, "products", old.Id, "UPDATE", old, new); logErr != nil {
		return logErr
	}

	return err
}

func UpdateOrder() {

}

func DeleteProduct(p models.Product, transaction *sql.Tx) error {
    query := `
        DELETE FROM products
        WHERE id=$1
    `
    _, err := transaction.Exec(query, p.Id)
    if logErr := logUpdate(transaction, "products", p.Id, "DELETE", p, nil); logErr != nil {
        return logErr
    }

    return err
}

func DeleteOrder() {

}

func GetProducts() ([]models.Product, error) {

	productRows, err := Db.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, err
	}
	defer productRows.Close()

	var products []models.Product

	for productRows.Next() {
		var p models.Product
		var images pq.StringArray

		if err := productRows.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Discount, &images, &p.Sold); err != nil {
			return nil, err
		}
		p.Images = images
		products = append(products, p)

	}
	return products, nil
}

func GetOrderHistory() {

}