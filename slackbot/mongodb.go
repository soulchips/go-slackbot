package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserStatus of a user along with their checkin history
type UserStatus struct {
	UserID        string    `bson:"userID"`
	Name          string    `bson:"name"`
	LastStatus    string    `bson:"last_status"`
	LastUpdate    time.Time `bson:"last_update"`
	DailyCheckins []Checkin `bson:"checkins"`
}

// Checkin object. Database will track daily check-ins for each user
type Checkin struct {
	Date           string    `bson:"checkin_date"`
	TimeStamp      time.Time `bson:"checkin_timestamp"`
	Status         string    `bson:"status"`
	CheckinMessage string    `bson:"message"`
}

// TODO: remove database and collection from args, set defaults and retrieve from .env

// creates a new instance of a user's status record
func (toBeStored UserStatus) createNew(database string, collection string) (*mongo.InsertOneResult, error) {
	collectionResult := client.Database(database).Collection(collection)
	insertResult, err := collectionResult.InsertOne(context.Background(), toBeStored)

	if err != nil {
		fmt.Println(err)
		return insertResult, err
	}

	return insertResult, nil
}

// Retrieve the mose recent data for the user
func getUserStatus(userID string, database string, collection string) (status UserStatus, err error) {
	collectionResult := client.Database(database).Collection(collection)
	filter := bson.D{primitive.E{Key: "userID", Value: userID}}
	result := UserStatus{}

	err = collectionResult.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (toBeStored UserStatus) userCheckin(db string, collectionName string) {

}
