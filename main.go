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
		fmt.Println(forecast.Temperature)
		gif, _ := giphyAPI.GetRandomGif(forecast.WeatherLevel)
		testText := "and can I get a location: "
		locationUrl := Location.OpenWeatherURL
		fmt.Println(testText + gif.URL)
		if forecast.Temperature >= 55.00 && forecast.Temperature < 70.00 {
			twitterAPI.PostTweet(testText + locationUrl)
		}
		time.Sleep(2 * time.Hour)
		fmt.Println(twitterAPI)
	}
}
