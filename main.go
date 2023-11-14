package main

import (
	"context"
	"fmt"
	"github.com/sidma64/todo-go/database"
	"log"
)

type Todo struct {
	Title       string
	Description string
	isDone      bool
}

type User struct {
	Name     string
	Email    string
	Password string
	Todos    []Todo
}

func main() {
	defer database.Disconnect()
	todos := []Todo{
		{
			Title:  "Clean your room",
			isDone: false,
		},
		{
			Title:  "Finish a programming project",
			isDone: false,
		},
		{
			Title:       "Do the homework from math class",
			Description: "The details are on the paper in my backpack",
			isDone:      true,
		},
	}
	_ = todos
	newUser := User{
		Name:     "sidma",
		Email:    "toprak.code@gmail.com",
		Password: "1323",
		Todos:    todos,
	}
	coll := database.DB.Collection("users")

	id, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
