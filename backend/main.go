
package main

import (
	"log"
	"net/http"
	"strings"
)


func main() {
    
    initDB();

    http.HandleFunc("/api/checkout", handleCheckout)
    http.HandleFunc("/api/products", handleProducts)

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, "/api/") {
			http.NotFound(res, req)
			return
		}
		http.ServeFile(res, req, "./frontend/build/index.html")
	})

    log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}