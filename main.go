package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	weatherAPI := WeatherAPIClient(os.Getenv("OWM_API_KEY"))
	twitterAPI := TwitterAPIClient(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_API_TOKEN"),
		os.Getenv("TWITTER_API_SECRET"),
	)

	giphyAPI := GiphyAPIClient(os.Getenv("GIPHY_API_KEY"))

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
