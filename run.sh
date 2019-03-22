#!/bin/bash

go get ./... &&
    gofmt -w *.go &&
    go build . &&
    export $(cat .env | xargs) &&
    ./nice-weather-bot
