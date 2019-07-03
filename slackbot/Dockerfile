FROM golang:1.12.1

WORKDIR /go/src/app
COPY . .

RUN go get github.com/pilu/fresh
RUN go get github.com/Krognol/go-wolfram
RUN go get github.com/christianrondeau/go-wit
RUN go get github.com/joho/godotenv
RUN go get github.com/nlopes/slack
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/mongo/options

CMD [ "fresh" ]