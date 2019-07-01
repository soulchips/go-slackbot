package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Krognol/go-wolfram"
	"github.com/christianrondeau/go-wit"
	"github.com/nlopes/slack"
)

var (
	slackClient   = slack.New(os.Getenv("SLACK_ACCESS_KEY"))
	witClient     = wit.NewClient(os.Getenv("WIT_AI_ACCESS_KEY"))
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
)

func main() {
	fmt.Println("starting connection...")

	// Creates a connection to slack using the bot's access_key
	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Connection counter:", ev.ConnectionCount)
			
		case *slack.MessageEvent:
			info := rtm.GetInfo()
			prefix := fmt.Sprintf("<@%s> ", info.User.ID)

			if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
				handleMessage(ev)
			}
		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
		}
	}
}
