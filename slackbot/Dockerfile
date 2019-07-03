FROM golang:1.12.1

WORKDIR /go/src/app
COPY . .

RUN go get github.com/pilu/fresh
RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "fresh" ]