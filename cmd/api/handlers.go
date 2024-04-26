package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"karen.assignment/shoup/internal/models"
)

func signUpHandler(userRepo *models.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		if requestBody.Username == "" || requestBody.Email == "" || requestBody.Password == "" {
			http.Error(w, "Username, email, and password are required", http.StatusBadRequest)
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		user := &models.User{
			Username:     requestBody.Username,
			Email:        requestBody.Email,
			PasswordHash: string(passwordHash),
		}

		err = userRepo.InsertUser(user)
		if err != nil {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User registered successfully"))
	}
}

func loginHandler(userRepo *models.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		if requestBody.Username == "" || requestBody.Password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		user, err := userRepo.GetUserByUsername(requestBody.Username)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestBody.Password)); err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte("aboba"))
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

func getItemsHandler(itemRepo *models.ItemRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		items, err := itemRepo.GetItems()
		if err != nil {
			http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func placeOrderHandler(orderRepo *models.OrderRepository, itemRepo *models.ItemRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Items []struct {
				ItemID   int `json:"item_id"`
				Quantity int `json:"quantity"`
			} `json:"items"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		if len(requestBody.Items) == 0 {
			http.Error(w, "Order must contain at least one item", http.StatusBadRequest)
			return
		}

		totalPrice := 0.0
		orderItems := []*models.OrderItem{}

		for _, item := range requestBody.Items {
			dbItem, err := itemRepo.GetItemByID(item.ItemID)
			if err != nil {
				http.Error(w, "Failed to fetch item", http.StatusInternalServerError)
				return
			}

			if dbItem == nil {
				http.Error(w, "Item not found", http.StatusBadRequest)
				return
			}

			subtotal := float64(item.Quantity) * dbItem.Price
			totalPrice += subtotal

			orderItem := &models.OrderItem{
				ItemID:   item.ItemID,
				Quantity: item.Quantity,
			}
			orderItems = append(orderItems, orderItem)
		}
		userID := r.Context().Value("userID").(int)

		order := &models.Order{
			Total:  totalPrice,
			UserID: userID,
		}
		orderID, err := orderRepo.InsertOrder(order)
		if err != nil {
			http.Error(w, "Failed to place order", http.StatusInternalServerError)
			return
		}

		for _, orderItem := range orderItems {
			orderItem.OrderID = orderID
			if err := orderRepo.InsertOrderItem(orderItem); err != nil {
				http.Error(w, "Failed to place order", http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Order placed successfully"))
	}
}
