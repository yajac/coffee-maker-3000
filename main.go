package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yajac/coffee-maker-3000/slack"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
// "https://hooks.slack.com/services/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"
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Request: %v\n", request)

	iotResponse, err := slack.HandleIOTEvent()

	fmt.Printf("IOTResponse: %v\n", iotResponse)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(iotResponse),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
