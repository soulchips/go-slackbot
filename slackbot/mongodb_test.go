package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoConnection(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check if connected to MongoDB
	if err != nil {
		t.Errorf("Unable to connect to mongodb, Error: %v", err)
	}

	// Check MongoDB connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Errorf("Unable to ping mongodb, Error: %v", err)
	}
}
