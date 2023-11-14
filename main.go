package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Restaurant struct {
	Name         string
	RestaurantId string        `bson:"restaurant_id,omitempty"`
	Cuisine      string        `bson:"cuisine,omitempty"`
	Address      interface{}   `bson:"address,omitempty"`
	Borough      string        `bson:"borough,omitempty"`
	Grades       []interface{} `bson:"grades,omitempty"`
}

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	fmt.Println(os.Getenv("MONGODB_URI"))
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_restaurants").Collection("restaurants")
	newRestaurant := Restaurant{Name: "8282", Cuisine: "Korean"}
	result, err := coll.InsertOne(context.TODO(), newRestaurant)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
