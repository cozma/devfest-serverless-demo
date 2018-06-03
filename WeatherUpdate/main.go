package WeatherUpdate

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ZachtimusPrime/go-wunderground/wunderground"
	"github.com/labstack/gommon/log"
)

type WeatherRequest struct {
	APIKey string
	State string
	City string
}

func WeatherHandler(request WeatherRequest) (interface{}, error) {
	// Client to call the weather underground API
	weatherClient := wunderground.NewClient(nil, request.State, request.City, request.APIKey)

	// Get Current Weather
	if curWeather, err := weatherClient.GetWeather(); err != nil {
		log.Info("Could not get current weather: ", err)
		return nil, err
	} else {
		return curWeather, nil
	}
}

func main() {
	lambda.Start(WeatherHandler)
}