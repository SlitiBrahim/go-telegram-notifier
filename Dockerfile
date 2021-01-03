FROM golang:1.15 AS build
WORKDIR /go/src/go-telegram-notifier
COPY . .
# install dependencies
RUN go get -v ./...
RUN go build -o telegram-notifier .

CMD ["./telegram-notifier"]