# Golang Slackbot
Other Names: JeffBot, CheckinBot, RoboRando

## Description
This project gives a slack bot the ability to answer general questions and it keeps track of a user's working status (In Office, Sick, etc). 

For simplicity, we'll refer to the bot as Jeffbot below.

## How it works
Jeff bot is able to determine the context of a user's message based on what is said. For example: Jeffbot knows to respond to a greeting with a hello, vs a request to change the user's status to sick, vs being asked who is the queen of england. The intent is to be able to speak to Jeffbot using natural language rather than just strict commands.

It recognizes natural language by leveraging Wit.ai's machine learning API to determine the intent of the message. 

#### Examples
Screenshots coming soon
Sample messages you can send
- hola
- What time is it in Tokyo?
- What is the price of Bitcoin in USD?
- I'm working from home today

## How to run locally
#### Install Golang
- Install [Golang](https://golang.org/doc/install)
- clone this repo
- get [slack bot api token](https://get.slack.help/hc/en-us/articles/215770388-Create-and-regenerate-API-tokens#-bot-user-tokens)
- Sign up for wit.ai and [get an access key](https://wit.ai/docs/quickstart)
- get [wolfram token](https://www.wolframalpha.com/tour/)
- install [mongodb](https://docs.mongodb.com/manual/installation/)
- You will need set the following as environment variables:
```
SLACK_ACCESS_KEY=<your_key_here>
WIT_AI_ACCESS_KEY=<your_key_here>
WOLFRAM_APP_ID=<your_app_id_here>
DEPLOYMENT=<your_environment_name>
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_DATABASE=slack
MONGO_COLLECTION=statuses
```

The then run the following from within the __slackbot directory__:
```
$ go run main.go slack.go mongodb.go
```
#### Or...

Jeffbot can also be run locally with docker using __docker compose__. A docker-compose.yml file is provided. First ensure [docker is installed](https://docs.docker.com/) and running. Create a .env file with the environment variables listed above.

Then run:
```
$ docker-compose up -d --build
```
This command will will deploys containers for mongodb and golang and run the project for us.

## How to test
From within the __slackbot directory__
```
$ go test
```

## How to build
From within the __slackbot directory__
```
$ go build
```

## Bot abilities
- Answer direct questions about facts
- Keep track of a user's work status
- Respond to greetings

## How to train your bot
Please see here for documentation on training wit.ai's text recognition. 
My account already has

#### Features ToDo:
- __Tell jokes__: Currently if you say to JeffBot something like "tell me a joke" or "I want to laugh" it will respond to say it's working on the jokes feature.
- __Track user checkins over time__: Jeffbot keeps track of user's most recent status but it may be useful to have a history of status changes.
- __Request checkins from a user__: It would be nice if JeffBot would remind you to checkin at a certain time.