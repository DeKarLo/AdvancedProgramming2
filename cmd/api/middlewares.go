package main

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("aboba"), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		userID, ok := claims["user_id"].(float64)

		if !ok {
			http.Error(w, "User ID not found in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", int(userID))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
