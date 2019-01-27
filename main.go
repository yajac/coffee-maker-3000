package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//Message properties
type Message struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	IconEmoji   string       `json:"icon_emoji"`
	ImageUrl    string       `json:"image_url"`
	Attachments []Attachment `json:"attachments"`
}

//Attachment properties
type Attachment struct {
	Fields []AttachmentFields `json:"fields"`
}

//AttachmentFields properties
type AttachmentFields struct {
	Title string `json:"title"`
	Short bool   `json:"short"`
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
// "https://hooks.slack.com/services/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"
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Request: %v\n", request)

	itemBytes, err := json.Marshal(Message{
		Channel:   "#richmondcoffee",
		Text:      "FRESH COFFEE!! - ",
		IconEmoji: ":coffee:",
	})

	fmt.Printf("itemJSON: %v\n", string(itemBytes))

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(itemBytes),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
