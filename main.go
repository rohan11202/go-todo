package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"p4nda/go-todo/models"
	"p4nda/go-todo/routes"
)

func main() {

	models.Initiate()

	router := mux.NewRouter()
	routes.RegisterTodoRoutes(router)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
