package main

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Post struct {
	Text      string    `json:"text" bson:"text"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
}

var posts *mgo.Collection


func writeToDB() {

}

func readFromDB() {
	
}