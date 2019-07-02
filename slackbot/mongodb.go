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
type User struct {
	UserID        string    `bson:"userID"`
	Name          string    `bson:"name"`
	LastStatus    string    `bson:"last_status"`
	LastUpdate    time.Time `bson:"last_update"`
	CreatedAt			time.Time `bson:"created_at"`
}

func newUser() *User{
	user := new(User)
	user.LastUpdate = time.Now().UTC()
	user.CreatedAt = time.Now().UTC()

	return user
}


// creates a new instance of a user's status record
func (toBeStored User) createNew(database string, collection string) (*mongo.InsertOneResult, error) {
	collectionResult := client.Database(database).Collection(collection)
	insertResult, err := collectionResult.InsertOne(context.Background(), toBeStored)

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

// Updates the user's status
func (toBeStored User) userCheckin(db string, collectionName string) {

}
