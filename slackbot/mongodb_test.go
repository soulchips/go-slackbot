package main

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// Define database and collection for testing
	database   = "test_database"
	collection = "test_statuses"
)

func TestMongoConnection(t *testing.T) {
	t.Run("Should be able to connect to mongodb", func(t *testing.T) {
		if mongoError != nil {
			t.Errorf("Unable to connect to mongodb, Error: %v\n", mongoError)
		}
	})

	t.Run("Should be able to ping mongodb", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := client.Ping(ctx, readpref.Primary())
		if err != nil {
			t.Errorf("Unable to ping mongodb, Error: %v\n", err)
		}
	})
}

func TestAddNewUser(t *testing.T) {
	user := UserStatus{
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

	t.Run("Should be able write a new userStatus instance to the collection", func(t *testing.T) {
		res, err := user.createNew(database, collection)
		if err != nil {
			t.Errorf("Unable to write to db, Error: %v\n", err)
		}

		if res.InsertedID == nil {
			t.Errorf("Expected insert response to have an InsertID but no ID was returned\n")
		}
	})
}

func TestGetUserInfo(t *testing.T) {
	userID := "TestID"

	t.Run("Should be able to search for userinfo by ID and get the correct data", func(t *testing.T) {
		result, err := getUserStatus(userID, database, collection)
		if err != nil {
			t.Errorf("Unable to read from db, Error: %v\n", err)
		}

		if result.Name != "Ash Ketchum" {
			t.Errorf("Expecting to find Ash Ketchum but instead found: %v\n", result.Name)
		}
	})
}
