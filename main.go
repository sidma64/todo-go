package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"net/http"
	"mime"
)

func init () {
    _ = mime.AddExtensionType(".js", "text/javascript")
}

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

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

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

	http.ListenAndServe(":3000", router())
}

type User struct {
	Name     string
	Email    string
	Password string
}

var users []User

func init() {
	users = append(users, User{
		"sidma",
		"toprak.code@gmail.com",
		"1323",
	})
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

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

func router() http.Handler {
	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("app")))
	r.Post("/login", handleLogIn)
	r.Get("/secret", secret)
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fgasdfa"))
	})
	return r
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
