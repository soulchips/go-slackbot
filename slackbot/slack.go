package main

import (
	"log"

	"github.com/Krognol/go-wolfram"
	"github.com/christianrondeau/go-wit"
	"github.com/nlopes/slack"
)

// Sends a text message to a slack channel by the provided user
func sendMessage(msg string, userID string, channelID string) {
	slackClient.PostMessage(channelID, slack.MsgOptionText(msg, false), slack.MsgOptionAsUser(true), slack.MsgOptionUser(userID))
}

// Checks for message events directed to the slackbot

func handleMessage(ev *slack.MessageEvent) {
	res, err := witClient.Message(ev.Msg.Text)
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

func replyToUser(ev *slack.MessageEvent, entityKey string, entity wit.MessageEntity) {
	switch entityKey {
	case "greetings":
		sendMessage("Hi there!", ev.User, ev.Channel)
		return

	case "wolfram_search_query":
		res, err := wolframClient.GetSpokentAnswerQuery(entity.Value.(string), wolfram.Metric, 1000)

		if err == nil && res != "Wolfram Alpha did not understand your input" {
			sendMessage(res, ev.User, ev.Channel)
			return
		}

		log.Printf("unable to get data from wolfram: %v", err)
	}

	sendMessage("I don't understand -\\_(0_0)_/-", ev.User, ev.Channel)
}
