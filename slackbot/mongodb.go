package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User - struct defining user attributes
type User struct {
	UserID     string    `bson:"userID"`
	Name       string    `bson:"name"`
	LastStatus string    `bson:"last_status"`
	LastUpdate time.Time `bson:"last_update"`
	CreatedAt  time.Time `bson:"created_at"`
}

// Creates a new User with default values
func newUser() User {
	user := new(User)
	user.LastUpdate = time.Now().UTC()
	user.CreatedAt = time.Now().UTC()

	return *user
}

// creates a new instance of a user's status record
func (user User) createNew(database string, collection string) (*mongo.InsertOneResult, error) {
	collectionResult := client.Database(database).Collection(collection)
	insertResult, err := collectionResult.InsertOne(context.Background(), user)

	if err != nil {
		fmt.Println(err)
		return insertResult, err
	}

	return insertResult, nil
}

// Retrieves the most recent data for the user
func getUserInfo(userID string, database string, collection string) (User, error) {
	collectionResult := client.Database(database).Collection(collection)

	filter := bson.D{primitive.E{Key: "userID", Value: userID}}
	result := User{}

	err := collectionResult.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

// Returns a user's current status
func getUserStatus(userID string, database string, collection string) (string, error) {
	user, err := getUserInfo(userID, database, collection)
	status := ""

	if err == nil {
		switch user.LastStatus {
		case "OOO":
			status = "out of office"
		case "WFH":
			status = "working from home"
		case "IO":
			status = "in office"
		case "SICK":
			status = "sick"
		}
	}

	return status, err
}

// Updates the user's info
func (user User) update(database string, collection string) (*mongo.UpdateResult, error) {
	collectionResult := client.Database(database).Collection(collection)

	filter := bson.D{primitive.E{Key: "userID", Value: user.UserID}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: user.Name},
			primitive.E{Key: "last_status", Value: user.LastStatus},
			primitive.E{Key: "last_update", Value: time.Now().UTC()},
		}},
	}
	updateOption := new(options.UpdateOptions)
	upsertValue := true
	updateOption.Upsert = &upsertValue
	// bson.D{primitive.E{Key: "upsert", Value: true }}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateResult, err := collectionResult.UpdateOne(ctx, filter, update, updateOption)

	return updateResult, err
}

// Updates the user's status
func (user User) checkin(status string, database string, collection string) {
	user.LastStatus = status
	user.update(database, collection)
}
