package routes

import (
	"github.com/gorilla/mux"
	"p4nda/go-todo/handlers"
)

var RegisterTodoRoutes = func(router *mux.Router) {
    router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
    router.HandleFunc("/todos", handlers.GetAllTodos).Methods("GET")
    router.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")
    router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
    router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
    router.HandleFunc("/todos/past-deadline", handlers.GetTodosPassedDeadline).Methods("GET")
}