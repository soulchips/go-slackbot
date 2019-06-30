package main

import (
	"fmt"
	"log"
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
	fmt.Println("starting connectin...")

	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			info := rtm.GetInfo()
			prefix := fmt.Sprintf("<@%s> ", info.User.ID)

			if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
				handleMessage(ev)
			}
		}
	}
}

func handleMessage(ev *slack.MessageEvent) {
	res, err:= witClient.Message(ev.Msg.Text)
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

	// Using the entity key, we determine the appropriate response
	switch entityKey {
	case "greetings":
		slackClient.PostMessage(ev.Channel, slack.MsgOptionText("Hi there!", false), slack.MsgOptionAsUser(true), slack.MsgOptionUser(ev.User))
		return

	case "wolfram_search_query":
		res, err := wolframClient.GetSpokentAnswerQuery(entity.Value.(string), wolfram.Metric, 1000)
		fmt.Printf("res: %v \n", res)
		fmt.Printf("err: %v \n", err)

		if err == nil {
			slackClient.PostMessage(ev.Channel, slack.MsgOptionText(res, false), slack.MsgOptionAsUser(true), slack.MsgOptionUser(ev.User))
			return
		}
		
		log.Printf("unable to get data from wolfram: %v", err)


	default:
		slackClient.PostMessage(ev.Channel, slack.MsgOptionText("I don't understand -\\__(0_0)__/-", false), slack.MsgOptionAsUser(true), slack.MsgOptionUser(ev.User))
	}
}
