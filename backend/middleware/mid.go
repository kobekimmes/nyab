package middleware

import (
	"net/http"
	"log"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req*http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if req.Method == http.MethodOptions {
			res.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(res, req)
	})
}


func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Incoming request %s %s from %s", req.Method, req.URL.Path, req.RemoteAddr)
		next.ServeHTTP(res, req)
	})
}

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if !limiter.Allow() {
			http.Error(res, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(res, req)
	})
}



