package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Status of a user along with their checkin history
type Status struct {
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

func (toBeStored Status) writeToDB(db string, collectionName string) (*mongo.InsertOneResult, error) {
	collection := client.Database(db).Collection(collectionName)

	insertResult, err := collection.InsertOne(context.TODO(), toBeStored)
	if err != nil {
		fmt.Println(err)
		return insertResult, err
	}

	fmt.Println("Inserted a single document: ", insertResult)
	return insertResult, nil
}

func readFromDB() {

}
