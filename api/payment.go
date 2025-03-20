package handler

import (
	"net/http"
	"ride-reservation/handlers"
)

func HandlerPayment(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/payment", handlers.PaymentHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}