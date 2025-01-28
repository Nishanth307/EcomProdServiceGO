package repository

import (
	// Go Internal Packages
	"context"
	"fmt"
	"time"

	// External Packages
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Client, error) {
	fmt.Println("Starting MongoDB Database Connection")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	// Ping the database to verify the connection
	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println("Error pinging MongoDB:", err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}