package main

import (
	"context"
	"testing"
	"time"
)

var (
	// Define database and collection for testing
	database   = "test_database"
	collection = "test_statuses"
)

func TestMongoConnection(t *testing.T) {
	// Check if connected to MongoDB
	if mongoError != nil {
		t.Errorf("Unable to connect to mongodb, Error: %v", mongoError)
	}

	// Check MongoDB connection
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		t.Errorf("Unable to ping mongodb, Error: %v", err)
	}
}

func TestWriteStatusToDatabase(t *testing.T) {
	user := Status{
		"TestID",
		"Ash Ketchum",
		"WFH",
		time.Now().UTC(),
		[]Checkin{
			Checkin{
				time.Now().Format("Jan-02-2006"),
				time.Now().UTC(),
				"WFH",
				"Hi I'm working from home today",
			},
		},
	}

	res, err := user.writeToDB(database, collection)
	if err != nil {
		t.Errorf("Unable to write to db, Error: %v", err)
	}

	if res.InsertedID == nil {
		t.Errorf("Expected insert response to have an InsertID but no ID was returned")
	}
}
