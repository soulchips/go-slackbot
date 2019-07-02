package main

import (
	"context"
	"testing"
	"time"
)

func TestUpdateStatus(t *testing.T) {
	user := newUser()
	user.UserID = "WRENEWID"
	user.Name = "Test Name"

	t.Run("Should be able to checkin a new user", func(t *testing.T) {
		err := userCheckIn(user.UserID, "IO", testDatabase, testCollection)

		if err != nil {
			t.Errorf("Unable to checkin new user. Error: %v", err)
		}
	})

	t.Run("Should be able checkin an existing user", func(t *testing.T) {

	})
}

func TestGetUserDataFromSlack(t *testing.T) {
	// this test requires a valid slack userID
	userID := "UL1S30423"

	t.Run("Should be able to retrieve user data from slack", func(t *testing.T) {
		slackUser, err := getSlackUserInfo(userID)

		if err != nil {
			t.Errorf("Unable to get user data from slack. Error: %v\n", err)
		}

		if slackUser.ID != userID {
			t.Errorf("Unable to get correct user data from slack. Expected: %v but got %v\n", userID, slackUser.ID)
		}
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
