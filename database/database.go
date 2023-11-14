package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var DB *mongo.Database

func init() {
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	DB = client.Database("todo")
}

func Disconnect() {
	if err := DB.Client().Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
