import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, 
		port, 
		user, 
		password, 
		dbname
	)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL")
}

func getProducts() {

}


func getOrderHistory() {

}