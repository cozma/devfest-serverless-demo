# DevFest Serverless Demo
> This application system will send an SMS notification on a daily basis
> to the desired subscription member with the daily weather

![Version 1.0](https://img.shields.io/badge/Version-1.0-yellow.svg)


![Example SNS](https://github.com/cozma/devfest-serverless-demo/blob/master/examplesns.gif)

## Setup Requirements

  - CloudWatch Rule Setup
  - SNS Topic

The application requires [you create a CloudWatch Rule](https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#rules) that will need to make a request to the Lambda on a daily basis. This is easily configurable by setting your event source to schedule on a 1 day fixed rate rotation and attaching the ARN of the current version of the Lambda Script. You also want to change the Configure Input to "Constant (JSON text)" and input the request JSON (sample below) for your Pager Duty Application.

### Pager Duty Lambda (PROD East) : 
> Name: devfest-serverless-demo

Sample Request:
```
{
  "snsarn": "arn:aws:sns:us-east-1:344449400151:WeatherSNS",
  "api_key": "cd3aceac328507a6c8d90506036899ec",
  "location": "20001,us"
}
```
