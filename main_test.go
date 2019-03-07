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
