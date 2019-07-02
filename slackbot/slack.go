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


// Get user's info from slack
func getSlackUserInfo(userID string) (slack.User, error) {
	user, err := slackClient.GetUserInfo(userID)
	if err != nil {
		log.Printf("%s\n", err)
		return *user, err
	}

	return *user, err
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
	log.Printf("message sent: %v", msgText)

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
			log.Printf("entity: %v\n", entity)
			log.Printf("EntityKey: %v\n", entityKey)
			log.Printf("Confidence: %v\n", entity.Confidence)
			log.Println()

			if entity.Confidence > topEntity.Confidence && entity.Confidence > minimumConfidence {
				topEntityKey = entityKey
				topEntity = entity
			}
		}
	}

	log.Printf("topEntity: %v\n", topEntity)
	log.Printf("topEntityKey: %v\n", topEntityKey)
	log.Println()

	replyToUser(ev, topEntityKey, topEntity)
}


// Replies to user based on the type of message received
func replyToUser(ev *slack.MessageEvent, entityKey string, entity wit.MessageEntity) {
	log.Printf("username: %v", ev.User)

	user := newUser()
	user.UserID = ev.User
	// user.Name = ev

	switch entityKey {
	case "greetings":
		sendMessage("Hi there!", ev.Channel)
		return

	case "funny":
		sendMessage("I can't tell jokes yet, but I'm working on it!", ev.Channel)
		return

	case "work_working_from_home":
		// userCheckIn(ev.User, ev.Username, "IO", slackDatabase, statusCollection)
		sendMessage("I've upated your status to working remotely. Hope you have a good day!", ev.Channel)
		return

	case "work_out_of_office":
		sendMessage("I've upated your status to out of office. Hope you have a good day!", ev.Channel)
		return

	case "work_in_office":
		sendMessage("I'm glad you're here! I've updated your status to in office.", ev.Channel)
		return

	case "work_sick":
		sendMessage("I'm sorry to hear that. I hope you feel better soon. I've updated your status to sick today.", ev.Channel)
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
