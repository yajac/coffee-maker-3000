package main

import (
	"./slack"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

	command := values["command"]

	fmt.Printf("Command: %v\n", command)

	username := values["user_name"][0]
	text := values["text"][0]
	channel := values["channel_name"][0]

	fmt.Printf("Request Values: %v %v\n", text, username)

	response, slackErr := slack.HandleMadeCoffeeEvent(channel, username)

	if slackErr != nil {
		return events.APIGatewayProxyResponse{}, slackErr
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(response),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
