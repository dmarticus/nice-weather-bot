## Nice Weather Bot

A Twitter bot that'll tweet every time the weather in Nice, France is 69 degrees Fahrenheit.  Nice.

## Building

1. Install [Go](https://golang.org/doc/install) and set up your `GOPATH` correctly.  
2. Set up a [twitter API account](https://developer.twitter.com/en/application/use-case) and an [OpenWeather API](https://openweathermap.org/appid) account to connect to the Twitter and OpenWeather APIs.
3. Create a `.env` file at your root directory that contains the following keys and their respective values from the Twitter API and Open Weather API accounts:
    * `TWITTER_CONSUMER_KEY`
    * `TWITTER_CONSUMER_SECRET`
    * `TWITTER_API_TOKEN`
    * `TWITTER_API_SECRET`
    * `OWM_API_KEY`
4. Execute `./build.sh`

## Running

Execute `./run.sh`
