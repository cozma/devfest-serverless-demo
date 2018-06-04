package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/gommon/log"
	"github.com/ramsgoli/Golang-OpenWeatherMap"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/aws"
)

type WeatherRequest struct {
	SNSArn string `json:"snsarn"`
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
		return PublishSNS(w.Weather[0], request.SNSArn), nil
	}
}

func PublishSNS(weather openweathermap.Weather, snsArn string) interface{} {
	svc := sns.New(session.New())
	message := "Good morning! There's currently " + weather.Description + " in your area."
	params := &sns.PublishInput{
		Message: aws.String(message),
		TopicArn: aws.String(snsArn),
	}

	if _, err := svc.Publish(params); err != nil {
		return err
	} else {
		return message
	}
}

func main() {
	lambda.Start(WeatherHandler)
}