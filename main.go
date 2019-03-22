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
		gif, _ := giphyAPI.GetRandomGif(forecast.WeatherLevel)
		testText := "and can I get a giphy: "
		fmt.Println(testText + gif.URL)
		twitterAPI.PostTweet(testText + gif.URL)
		time.Sleep(2 * time.Hour)
		fmt.Println(twitterAPI)
	}
}
