package main

import (
	"context"
	"fmt"
	"time"
)

type status struct {
	userID        string
	name          string
	lastStatus    string
	lastUpdate    time.Time
	dailyCheckins []checkin
}

type checkin struct {
	date           string
	timeStamp      string
	status         string
	checkinMessage string
}

func (toBeStored status) writeToDB(db string, collectionName string) bool {
	collection := client.Database(db).Collection(collectionName)

	insertResult, err := collection.InsertOne(context.TODO(), toBeStored)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

func readFromDB() {

}
