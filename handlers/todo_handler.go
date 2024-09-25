package handlers

import (
	"encoding/json"
	"net/http"
	"p4nda/go-todo/config"
	"p4nda/go-todo/models"
	"time"

	"github.com/gorilla/mux"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.GetDB().Create(&todo).Error; err != nil {
		http.Error(w, "Error creating todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	if err := config.GetDB().Find(&todos).Error; err != nil {
		http.Error(w, "Error fetching todos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var todo models.Todo
	if err := config.GetDB().First(&todo, id).Error; err != nil {
		http.Error(w, "Todo not found: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var todo models.Todo
	if err := config.GetDB().First(&todo, id).Error; err != nil {
		http.Error(w, "Todo not found: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.GetDB().Save(&todo).Error; err != nil {
		http.Error(w, "Error updating todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := config.GetDB().Delete(&models.Todo{}, id).Error; err != nil {
		http.Error(w, "Error deleting todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetTodosPassedDeadline(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	if err := config.GetDB().Where("deadline < ?", time.Now()).Find(&todos).Error; err != nil {
		http.Error(w, "Error fetching todos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
