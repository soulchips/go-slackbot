package main

import (
	"github.com/nlopes/slack"
)

// Sends a text message to a slack channel by the provided user
func sendMessage(msg string, userID string, channelID string) {
	slackClient.PostMessage(channelID, slack.MsgOptionText(msg, false), slack.MsgOptionAsUser(true), slack.MsgOptionUser(userID))
}
