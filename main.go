package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/google/uuid"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type Todo struct {
	ID     uint32  `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func NewTodo(title string) Todo {
	return Todo{
		ID: uuid.New().ID(),
		Title:  title,
		IsDone: false,
	}
}

type User struct {
	Name     string
	Email    string
	Password string
}

var users []User
var todos []Todo

func main() {
	todos = append(todos, NewTodo("Learn Microsoft Excel"))
	todos = append(todos, NewTodo("Finish homework"))
	users = append(users, User{
		"sidma",
		"toprak.code@gmail.com",
		"1323",
	})

	http.ListenAndServe(":3000", router())
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("app")))

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", handleLogIn)
		r.Get("/secret", secret)
		r.HandleFunc("/todos", handleTodos)
	})
	return r
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(todos)
	default:
		http.Error(w, "Unsupported method", 400)
		return
	}
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-key")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func handleLogIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		loginForm := struct {
			Email    string
			Password string
		}{}
		err := json.NewDecoder(r.Body).Decode(&loginForm)
		if err != nil {
			http.Error(w, "JSON is invalid", http.StatusBadRequest)
		}
		valid := validateUser(loginForm.Email, loginForm.Password)
		if !valid {
			http.Error(w, "Invalid login information", http.StatusUnauthorized)
			return
		}
		session, _ := store.Get(r, "session-key")
		session.Values["authenticated"] = true
		session.Save(r, w)
		w.Write([]byte("Done"))

	default:
		http.Error(w, "Unsupported method", 400)
	}
}

func validateUser(usernameOrEmail string, password string) bool {
	for _, user := range users {
		if !(user.Name == usernameOrEmail || user.Email == usernameOrEmail) {
			continue
		}
		if user.Password == password {
			return true
		}
		break
	}
	return false
}
