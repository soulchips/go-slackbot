package main

import (
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

// Saves user checking data, creates a new user if user isnt found
func userCheckIn(userID string, username string, status string, database string, collection string) bool {
	// get userinfo from db
	user, err := getUserInfo(userID, database, collection)

	// if no info, create new user
	if(err != nil || user.UserID == "") {
		user = newUser()
		user.UserID = userID
		user.Name = username
	}

	user.LastStatus = status

	updateResult, err := user.update(database, collection)

	if(err != nil || updateResult.UpsertedID == nil) {
		return false
	}
	return true
}

// Checks for message events directed to the slackbot
func handleMessage(ev *slack.MessageEvent) {

	// remove slackbot id from msg string
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
			if entity.Confidence > topEntity.Confidence && entity.Confidence > minimumConfidence {
				topEntityKey = entityKey
				topEntity = entity
			}
		}
	}

	replyToUser(ev, topEntityKey, topEntity)
}

// Replies to user based on the type of message received
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
