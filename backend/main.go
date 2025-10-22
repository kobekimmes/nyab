
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"
)


// POST /api/checkout
func postCheckout(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

}


// GET /api/products
func getProducts(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	
}