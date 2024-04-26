package api

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *mux.Router) {
	// Define routes
	router.HandleFunc("/signup", signUpHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/items", getItemsHandler).Methods("GET")
	router.HandleFunc("/order", placeOrderHandler).Methods("POST")
}
