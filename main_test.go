package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	if testing.Short() {
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

}

func TestOrderUserMap(t *testing.T) {
	type args struct {
		userMap map[string]int
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"Order Map 1 In Order", args{map[string]int{"user4": 4, "user3": 3, "testuser": 2, "imcewan": 1}},
			[]User{{"user4", 4}, {"user3", 3}, {"testuser", 2}, {"imcewan", 1}},
		},
		{"Order Map 2 Out Order", args{map[string]int{"user3": 3, "imcewan": 1, "user4": 4, "testuser": 2}},
			[]User{{"user4", 4}, {"user3", 3}, {"testuser", 2}, {"imcewan", 1}},
		},
		{"Order Map 3 Too Long", args{map[string]int{"user3": 3, "imcewan": 1, "user5": 5, "user6": 6, "user7": 7, "user4": 4, "user8": 8, "user9": 9, "user10": 10, "extraUser": 0, "testuser": 2}},
			[]User{{"user10", 10}, {"user9", 9}, {"user8", 8}, {"user7", 7}, {"user6", 6}, {"user5", 5}, {"user4", 4}, {"user3", 3}, {"testuser", 2}, {"imcewan", 1}},
		},
		{"Order Map 4 Same", args{map[string]int{"user3": 3, "3user3": 3, "testuser": 2, "imcewan": 1}},
			[]User{{"user3", 3}, {"3user3", 3}, {"testuser", 2}, {"imcewan", 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderUserMap(tt.args.userMap)
			if !testEq(got, tt.want) {
				t.Errorf("OrderUserMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testEq(a, b []User) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		fmt.Printf("Compare Nil Fail: %v  %v\n", a, b)
		return false
	}

	if len(a) != len(b) {
		fmt.Printf("Compare Length Fail: %v  %v\n", a, b)
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			fmt.Printf("Compare Item: %v  %v\n", a[i], b[i])
			return false
		}
	}

	return true
}

func TestGetUserText(t *testing.T) {
	type args struct {
		userMap map[string]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Get User Text 1", args{map[string]int{"Ian": 9, "Test": 8, "Bah": 7, "imcewan": 1}}, []string{"Ian         9", "Test        8", "Bah         7", "imcewan     1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserText(tt.args.userMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleAction(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test 1", args{"{\"type\": \"interactive_message\",\"actions\": [{\"name\": \"madeCoffee\",\"type\": \"button\",\"value\": \"madecoffee\"}],\"callback_id\": \"made_coffee_button\",\"team\": {\"id\": \"T20SR8Z88\",\"domain\": \"veracity-group\"},\"channel\": {\"id\": \"G4SUM84ES\",\"name\": \"privategroup\"},\"user\": {\"id\": \"U4QV2NNLD\",\"name\": \"imcewan\"},\"action_ts\": \"1553653180.218368\",\"message_ts\": \"1553640191.001300\",\"attachment_id\": \"1\",\"token\": \"wT0pZ9TvmWNUgwBt5G21UTPA\",\"is_app_unfurl\": false,\"original_message\": {\"type\": \"message\",\"subtype\": \"bot_message\",\"text\": \"Coffee Started Cool Beans 2\",\"ts\": \"1553640191.001300\",\"username\": \"CoffeeMaker3000\",\"icons\": {\"emoji\": \":coffee:\",\"image_64\": \"https://a.slack-edge.com/37d58/img/emoji_2017_12_06/apple/2615.png\"        },\"bot_id\": \"BFKFXHQRL\",\"attachments\": [{\"image_url\": \"https://media2.giphy.com/media/KpB7H0EPBWdDW/giphy.gif\",\"fields\": [{\"title\": \"Coffee is Ready\",\"value\": \"\",\"short\": false}]},{\"callback_id\": \"made_coffee_button\",\"fallback\": \"Error\",\"id\": 1,\"actions\": [{\"id\": \"1\",  \"name\": \"madeCoffee\",\"text\": \"I made the Coffee!\",\"type\": \"button\",\"value\": \"madecoffee\",\"style\": \"\"}]}]},\"response_url\": \"https://hooks.slack.com/actions/T20SR8Z88/586845881424/KYPy4wa5PNudR7lPSklnyeuL\",\"trigger_id\": \"589231921588.68909305280.c35b382fab7ce582243832e964f1e7ff\"}"}, "{\"channel\":\"G4SUM84ES\",\"text\":\"Coffee Started Cool Beans 2\",\"attachments\":[{\"fields\":[{\"title\":\"Coffee is Ready\",\"short\":false}],\"image_url\":\"https://media2.giphy.com/media/KpB7H0EPBWdDW/giphy.gif\"}]}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handleAction(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("handleAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
