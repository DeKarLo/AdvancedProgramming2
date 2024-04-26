package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"karen.assignment/shoup/internal/models"
)

func main() {
	router := mux.NewRouter()
	db, err := sql.Open("postgres", "postgres://postgres:dek@123455@localhost:5432/golab?sslmode=disable")

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	userRepo := &models.UserRepository{DB: db}
	itemRepo := &models.ItemRepository{DB: db}
	orderRepo := &models.OrderRepository{DB: db}

	RegisterRoutes(router, userRepo, itemRepo, orderRepo)

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
