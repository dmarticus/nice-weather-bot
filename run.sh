#!/bin/bash

./build.sh && export $(cat .env | xargs) && ./nice-weather-bot
