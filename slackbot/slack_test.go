package main

import (
	"context"
	"testing"
	"time"
)

func TestUpdateStatus(t *testing.T) {
	user := newUser()
	user.UserID = "WRENEWID"
	user.Name = "WRENEWID"

	t.Run("Should be able to checkin a new user", func(t *testing.T) {
		res := userCheckIn("user.UserID", "user.Name", "IO", testDatabase, testCollection)

		if !res {
			t.Errorf("Unable to checkin new user")
		}
	})

	t.Run("Should be able checkin an existing user", func(t *testing.T) {

	})
}

func TestDataCleanUp(t *testing.T) {
	// clean up testing database and confirm empty
	collectionResult := client.Database(testDatabase).Collection(testCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collectionResult.Drop(ctx)
	if err != nil {
		t.Errorf("Unable to drop collection, Error: %v\n", err)
	}
}
