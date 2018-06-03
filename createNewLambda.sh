#!/bin/bash

GOOS=linux go build -o main WeatherUpdate/main.go
zip DevFestDemoLambda.zip main
aws lambda create-function --function-name "devfest-serverless-demo"  --runtime "go1.x" --role "arn:aws:iam::344449400151:role/Lambda" --handler "main" --description "Send Current Weather" --zip-file fileb://DevFestDemoLambda.zip