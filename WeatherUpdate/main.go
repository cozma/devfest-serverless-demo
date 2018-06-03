package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/gommon/log"
	"github.com/ramsgoli/Golang-OpenWeatherMap"
)

type WeatherRequest struct {
	APIKey string `json:"api_key"`
	Location string `json:"location"`
}

func WeatherHandler(request WeatherRequest) (interface{}, error) {
	owm := openweathermap.OpenWeatherMap{
		API_KEY: request.APIKey,
	}

	if w, err := owm.CurrentWeatherFromCity(request.Location); err != nil {
		log.Error(err)
		return nil, err
	} else {
		log.Info("Current Weather: ", w)
		return w, nil
	}
}

func main() {
	lambda.Start(WeatherHandler)
}