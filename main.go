package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func (user *User) Save(db *sql.DB) (id int64, err error) {
	if _, err = GetUserByEmail(user.Email, db); !errors.Is(err, sql.ErrNoRows) {
		return 0, errors.New("email taken")
	}
	if _, err = GetUserByName(user.Name, db); !errors.Is(err, sql.ErrNoRows) {
		return 0, errors.New("username taken")
	}

	result, err := db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println("hello")
		return 0, err
	}
	id, err = result.LastInsertId()
	return id, err
}

func GetUserByName(name string, db *sql.DB) (user User, err error) {
	err = db.QueryRow("SELECT name, email, password FROM user WHERE name=?", name).Scan(&user.Name, &user.Email, &user.Password)
	return user, err
}

func GetUserByEmail(email string, db *sql.DB) (user User, err error) {
	err = db.QueryRow("SELECT name, email, password FROM user WHERE email=?", email).Scan(&user.Name, &user.Email, &user.Password)
	return user, err
}

func LoginWithEmail(email string, password string) () {

}

func main() {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	newUser := User{
		Name:     "red",
		Email:    "red@example.com",
		Password: "132325",
	}
	newUserId, err := newUser.Save(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUserId)
}
