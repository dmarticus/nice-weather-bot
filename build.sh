#!/bin/bash

go get ./... &&
    gofmt -w *.go &&
    go build .