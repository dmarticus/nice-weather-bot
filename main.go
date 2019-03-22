package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	weatherAPI := NewWeatherApi(os.Getenv("OWM_API_KEY"))
	twitterAPI := NewTwitterService(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_API_TOKEN"),
		os.Getenv("TWITTER_API_SECRET"),
	)

	for {
		forecast := weatherAPI.Now()
		fmt.Println(time.Now())
		fmt.Println(forecast.Weather)
		testText := "sup family I'm alive"
		fmt.Println(testText)
		twitterAPI.PostTweet(testText)
		time.Sleep(2 * time.Hour)
		fmt.Println(twitterAPI)
	}
}
