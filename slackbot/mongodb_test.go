package main

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	testDatabase   = "test_database"
	testCollection = "test_statuses"
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
	user := newUser()
	user.UserID = "TestID"
	user.Name = "Ash Ketchum"
	user.LastStatus = "WFH"

	t.Run("Should be able write a new userStatus instance to the collection", func(t *testing.T) {
		res, err := user.createNew(testDatabase, testCollection)
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
		result, err := getUserInfo(userID, testDatabase, testCollection)
		if err != nil {
			t.Errorf("Unable to read from db, Error: %v\n", err)
		}

		if result.Name != "Ash Ketchum" {
			t.Errorf("Expecting to find Ash Ketchum but instead found: %v\n", result.Name)
		}
	})
}

func TestUpdateUserInfo(t *testing.T) {
	t.Run("Should be able to update user's status", func(t *testing.T) {

	})

	t.Run("Should be able to update user's basic info", func(t *testing.T) {

	})
}

func TestDatabaseCleanUp(t *testing.T) {
	// clean up testing database and confirm empty
	collectionResult := client.Database(testDatabase).Collection(testCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collectionResult.Drop(ctx)
	if err != nil {
		t.Errorf("Unable to drop collection, Error: %v\n", err)
	}
}
