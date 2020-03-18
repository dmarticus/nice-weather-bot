package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterService struct {
	Api *anaconda.TwitterApi
}

func TwitterAPIClient(consumerKey string, consumerSecret string, apiToken string, apiSecret string) TwitterService {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(apiToken, apiSecret)

	return TwitterService{
		Api: api,
	}
}

func (self TwitterService) PostTweet(text string) {
	_, err := self.Api.PostTweet(text, nil)
	if err != nil {
		fmt.Printf("there was an error posting the tweet: %s\n", err)
	}
}
