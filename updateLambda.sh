#!/bin/bash

GOOS=linux go build -o main WeatherUpdate/main.go
zip DevFestDemoLambda.zip main
aws lambda update-function-code --region us-east-1 --function-name "devfest-serverless-demo" --zip-file fileb://DevFestDemoLambda.zip