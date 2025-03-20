package main

import (
	"fmt"
	"net/http"
	"ride-reservation/handlers"
)

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ride", handlers.Rides)
	mux.HandleFunc("/payment", handlers.PaymentHandler)

	fmt.Println("Server running on 5000")
	http.ListenAndServe(":5000", enableCORS(mux))
}
