package handler

import (
	"net/http"
	"ride-reservation/handlers"
)

func HandlerRide(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/ride", handlers.Rides)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}