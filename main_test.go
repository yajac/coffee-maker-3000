package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Body: "token=gIkuvaNzQIHg97ATvDxqgjtO&team_id=T0001&team_domain=example&enterprise_id=E0001&enterprise_name=Globular%20Construct%20Inc&channel_id=C2147483705&channel_name=coffee&user_id=U2147483697&user_name=TestUser&command=/madecoffee&text=SINGLE&response_url=https://hooks.slack.com/commands/1234/5678&trigger_id=13345224609.738474920.8088930838d88f008e0",
	}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "{\"channel\":\"#coffee\",\"text\":\"Coffee made by TestUser\",\"icon_emoji\":\":star2:\"}",
	}

	response, err := Handler(request)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedResponse.Headers, response.Headers)
	assert.Equal(t, expectedResponse.Body, response.Body)

}
