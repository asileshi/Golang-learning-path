package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UserCollection *mongo.Collection
	TaskCollection *mongo.Collection
)

func ConnectDB() error {
	// Set up MongoDB connection options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	// Set the user and task collections
	UserCollection = client.Database("task_manager").Collection("users")
	TaskCollection = client.Database("task_manager").Collection("tasks")

	return nil
}