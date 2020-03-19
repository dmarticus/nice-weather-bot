package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kyokomi/emoji"
)

func main() {
	weatherAPI := WeatherAPIClient(os.Getenv("OWM_API_KEY"))
	twitterAPI := TwitterAPIClient(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_API_TOKEN"),
		os.Getenv("TWITTER_API_SECRET"),
	)

	for {
		forecast := weatherAPI.Now()
		forecastString := "It's " + fmt.Sprintf("%.0f", forecast.Temperature) + " degrees Fahrenheit in Nice currently."
		detailsString := "\nProof: " + Location.OpenWeatherURL
		fmt.Println(forecastString + detailsString)
		if forecast.Temperature >= 69.00 && forecast.Temperature < 70.00 {
			twitterAPI.PostTweet(forecastString + emoji.Sprint(" Niiice :fire: :fire: fire:") + detailsString)
		}
		time.Sleep(2 * time.Hour)
	}
}
