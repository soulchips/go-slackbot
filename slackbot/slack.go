package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Krognol/go-wolfram"
	"github.com/christianrondeau/go-wit"
	"github.com/nlopes/slack"
)

// Sends a text message to a slack channel by the provided user
func sendMessage(msg string, channelID string) {
	slackClient.PostMessage(channelID, slack.MsgOptionText(msg, false), slack.MsgOptionAsUser(true))
}

func userCheckIn(userID string, status string) {
	// get userinfo from db

	// if no info, create new user

	// update user's status

	// save status
}

// Checks for message events directed to the slackbot

func handleMessage(ev *slack.MessageEvent) {
	fmt.Printf("Message received: %v\n", ev.Msg.Text)
	fmt.Println()
 
	msgText := strings.Replace(ev.Msg.Text, slackBotIDString, "", 1)

	res, err := witClient.Message(msgText)
	if err != nil {
		log.Printf("Unable to connect to Wit.ai. Error: %v", err)
		return
	}

	var (
		topEntity         wit.MessageEntity
		topEntityKey      string
		minimumConfidence = 0.5
	)

	for entityKey, entityList := range res.Entities {
		for _, entity := range entityList {
			fmt.Printf("Entity: %v\n", entity)
			fmt.Printf("EntityKey: %v\n", entityKey)
			fmt.Printf("Confidence: %v\n", entity.Confidence)
			fmt.Println()

			if entity.Confidence > topEntity.Confidence && entity.Confidence > minimumConfidence {
				topEntityKey = entityKey
				topEntity = entity
			}
		}
	}

	fmt.Printf("Top topEntity: %v\n", topEntity)
	fmt.Printf("Top topEntityKey: %v\n", topEntityKey)
	fmt.Printf("User ID: %v\n", ev.User)

	replyToUser(ev, topEntityKey, topEntity)
}

func replyToUser(ev *slack.MessageEvent, entityKey string, entity wit.MessageEntity) {
	switch entityKey {
	case "greetings":
		sendMessage("Hi there!", ev.Channel)
		return

	case "joke":
		sendMessage("I can't tell jokes yet, but I'm working on it!", ev.Channel)
		return

	case "wolfram_search_query":
		res, err := wolframClient.GetSpokentAnswerQuery(entity.Value.(string), wolfram.Metric, 1000)

		if err == nil && res != "Wolfram Alpha did not understand your input" {
			sendMessage(res, ev.Channel)
			return
		}

		log.Printf("unable to get data from wolfram: %v", err)
	}

	sendMessage("I don't understand -\\_(0_0)_/-", ev.Channel)
}

// determine if ev.User is slackbot id || user id is msg sender's id
// remove slack id from string. create string "<id> " to search by
