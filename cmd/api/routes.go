package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"karen.assignment/shoup/internal/models"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *mux.Router, userRepo *models.UserRepository, itemRepo *models.ItemRepository, orderRepo *models.OrderRepository) {
	router.HandleFunc("/v1/signup", signUpHandler(userRepo)).Methods("POST")
	router.HandleFunc("/v1/login", loginHandler(userRepo)).Methods("POST")
	router.HandleFunc("/v1/items", getItemsHandler(itemRepo)).Methods("GET")
	router.HandleFunc("/v1/order", func(w http.ResponseWriter, r *http.Request) {
		AuthMiddleware(placeOrderHandler(orderRepo, itemRepo)).ServeHTTP(w, r)
	}).Methods("POST")
}
