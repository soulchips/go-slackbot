package main

import (
	"strings"
	"fmt"
	"log"
	"os"

	"github.com/christianrondeau/go-wit"
	"github.com/nlopes/slack"
)

var (
	slackClient = slack.New(os.Getenv("SLACK_ACCESS_KEY"))
	witClient = wit.NewClient(os.Getenv("WIT_AI_ACCESS_KEY"))
)


func main() {
	fmt.Println("starting connectin...")

	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					// rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy!?!?", ev.Channel))
					handleMessage(ev)
				}
		}
	}
}


func handleMessage(ev *slack.MessageEvent) {
	response, error := witClient.Message(ev.Msg.Text)
	if error != nil {
		log.Printf("Unable to connect to Wit.ai. Error: %v", error)
		return
	}

	var (
		topEntity wit.MessageEntity
		topEntityKey string
		minimumConfidence = 0.5
	)

	for entityKey, entityList := range response.Entities {
		for _, entity := range entityList {
			if entity.Confidence > topEntity.Confidence && entity.Confidence > minimumConfidence {
				topEntityKey = entityKey
				topEntity = entity
			}
		}

		// fmt.Printf("topEntity [0] %v \n", entityList[0])
	}
	// fmt.Printf("%v \n", response)
	// fmt.Printf("topEntityKey %v \n", topEntityKey)
	// fmt.Printf("topEntity %v \n", topEntity)

	replyToUser(ev, topEntityKey, topEntity)
}

func replyToUser(ev *slack.MessageEvent, entityKey string, entity wit.MessageEntity) {
	// rtm := slackClient.NewRTM()
	// go rtm.ManageConnection()

	// rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy!?!?", ev.Channel))

	// fmt.Printf("%v \n", ev)
	// fmt.Printf("%v \n", ev.User)
	// fmt.Printf("%v \n", ev.Channel)
	// fmt.Printf("%v \n", ev.Text)

	slackClient.PostMessage(ev.Channel, slack.MsgOptionText("Hello World!", false))

	// switch entityKey {
	// case "greetings": 
	// 	slackClient.PostMessage("UL014TWUQ", slack.MsgOptionText("Hello world", false))
	// 	// rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy!?!?", ev.Channel))
	// 	return
	// }

	// slackClient.PostMessage("UL014TWUQ", slack.MsgOptionText("I dont understand", false), slack.MsgOptionAsUser(true), slack.MsgOptionUser(ev.User))

	fmt.Println("...done")
}
