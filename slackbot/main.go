package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Krognol/go-wolfram"
	"github.com/christianrondeau/go-wit"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoHost						= os.Getenv("MONGO_HOST")
	mongoPort						= os.Getenv("MONGO_PORT")
	slackClient        	= slack.New(os.Getenv("SLACK_ACCESS_KEY"))
	witClient          	= wit.NewClient(os.Getenv("WIT_AI_ACCESS_KEY"))
	wolframClient      	= &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
	clientOptions      	= options.Client().ApplyURI("mongodb://" + mongoHost + ":" + mongoPort)
	ctx, _             	= context.WithTimeout(context.Background(), 10*time.Second)
	client, mongoError 	= mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	slackBotIDString		string
	slackDatabase				= os.Getenv("MONGO_DATABASE")
	statusCollection		= "statuses"
)

func main() {
	// Loads environment variables from .env if present
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("starting connections...")

	// Check if connected to MongoDB
	if mongoError != nil {
		log.Fatal(mongoError)
	}

	// Check MongoDB connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Creates a connection to slack using the bot's access_key
	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Slack connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			info := rtm.GetInfo()
			prefix := fmt.Sprintf("<@%s> ", info.User.ID)
			slackBotIDString = prefix

			if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
				handleMessage(ev)
			}
		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Fatal("Invalid credentials")
		}
	}
}
