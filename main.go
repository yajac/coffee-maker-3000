package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yajac/coffee-maker-3000/dynamodb"
	"github.com/yajac/coffee-maker-3000/slack"
	"net/url"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Request: %v\n", request)

	fmt.Printf("Request Body: %v\n", request.Body)

	values, err := url.ParseQuery(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Printf("Values: %v\n", values)

	command := values["command"][0]

	fmt.Printf("Command: %v\n", command)

	username := values["user_name"][0]
	text := values["text"][0]
	channel := values["channel_name"][0]

	fmt.Printf("Request Values: %v %v\n", text, username)

	var jsonResponse string

	if command == "/madecoffee" {
		dbErr := dynamodb.UpdateLastCoffee(username)
		if dbErr != nil {
			fmt.Printf("DB Error: %v\n", dbErr)
			return events.APIGatewayProxyResponse{}, dbErr
		}

		response, slackErr := slack.HandleMadeCoffeeEvent(channel, username)
		jsonResponse = string(response)

		if slackErr != nil {
			fmt.Printf("Slack Error: %v\n", slackErr)
			return events.APIGatewayProxyResponse{}, slackErr
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResponse),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
