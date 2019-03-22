package main

import (
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

type Forecast struct {
	Humidity     int
	High         float64
	Low          float64
	WindSpeed    float64
	Clouds       int
	Weather      string
	WeatherLevel string
	WeatherId    int
	Temperature  float64
}

type WeatherApi struct {
	Location string
}

type LocationInfo struct {
	City           string
	OpenWeatherURL string
}

// set specific location here
// TODO make this in some sort of config file
var Location = LocationInfo{
	City:           "Nice",
	OpenWeatherURL: "https://openweathermap.org/city/6454924",
}

func WeatherAPIClient(apiKey string) WeatherApi {
	os.Setenv("OWM_API_KEY", apiKey) // owm requires this env var to function

	return WeatherApi{
		Location: Location.City,
	}
}

func (self WeatherApi) Now() Forecast {
	weather, err := owm.NewCurrent("F", "EN")
	if err != nil {
		log.Fatalln(err)
	}

	weather.CurrentByName(self.Location)

	return Forecast{
		Humidity:     weather.Main.Humidity,
		High:         weather.Main.TempMax,
		Temperature:  weather.Main.Temp,
		Low:          weather.Main.TempMin,
		WindSpeed:    weather.Wind.Speed,
		Clouds:       weather.Clouds.All,
		Weather:      weather.Weather[0].Main,
		WeatherLevel: weather.Weather[0].Description,
		WeatherId:    weather.Weather[0].ID,
	}
}
