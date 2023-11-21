package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func NewTodo(title string) Todo {
	return Todo{
		Title:  title,
		IsDone: false,
	}
}

var todos []Todo

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(todos)
	default:
		http.Error(w, "Unsupported method", 400)
		return
	}
}

func main() {
	todos = append(todos, NewTodo("Learn Microsoft Excel"))
	todos = append(todos, NewTodo("Finish homework"))

	fileServer := http.FileServer(http.Dir("./app")) // New code
    http.Handle("/", fileServer) // New code

	http.HandleFunc("/todos", handleGetTodos)

	http.ListenAndServe(":3000", nil)
}
