package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yajac/coffee-maker-3000/dynamodb"
	"github.com/yajac/coffee-maker-3000/slack"
	"net/url"
	"os"
	"sort"
	"strconv"
)

//User coffee usage
type User struct {
	user   string
	coffee int
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Request: %v\n", request)

	key := os.Getenv("SlackKey")

	fmt.Printf("SlackKey: %v\n", key)

	fmt.Printf("Request Body: %v\n", request.Body)

	values, err := url.ParseQuery(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Printf("Values: %v\n", values)

	jsonResponse := ""

	if values["command"] != nil {
		command := values["command"][0]

		fmt.Printf("Command: %v\n", command)

		username := values["user_name"][0]
		text := values["text"][0]
		channel := values["channel_name"][0]

		fmt.Printf("Request Values: %v %v\n", text, username)

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

		if command == "/coffeeleader" {
			userMap, dbErr := dynamodb.GetUsers()
			fmt.Printf("UserMap: %v\n", userMap)
			if dbErr != nil {
				fmt.Printf("DB Error: %v\n", dbErr)
				return events.APIGatewayProxyResponse{}, dbErr
			}
			fmt.Printf("UserMap: %v\n", userMap)
			userList := GetUserText(userMap)
			response, slackErr := slack.HandleLeaderBoard(channel, userList)

			jsonResponse = string(response)
			if slackErr != nil {
				fmt.Printf("Slack Error: %v\n", slackErr)
				return events.APIGatewayProxyResponse{}, slackErr
			}
		}
	}

	if values["payload"] != nil {
		payload := values["payload"][0]
		fmt.Printf("Payload: %v\n", payload)
		jsonResponse, err = handleAction(payload)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
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

//GetUserText for users get list of text
func GetUserText(userMap map[string]int) []string {
	users := OrderUserMap(userMap)
	fmt.Printf("Users: %v\n", users)
	var userList []string
	for _, user := range users {
		fmt.Printf("UserList user: %v\n", user)
		userList = append(userList, fmt.Sprintf("%-12v", user.user)+strconv.Itoa(user.coffee))
	}
	fmt.Printf("UserList: %v\n", userList)
	return userList
}

//OrderUserMap takes in a user map of names and values and returns an ordered list of User objects
func OrderUserMap(userMap map[string]int) []User {
	fmt.Printf("UserMap Before: %v\n", userMap)

	var users []User

	for k := range userMap {
		user := User{k, userMap[k]}
		users = append(users, user)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].coffee > users[j].coffee
	})

	if len(users) > 10 {
		users = users[:10]
	}
	fmt.Printf("UserMap After: %v\n", userMap)
	return users
}

func handleAction(payload string) (string, error) {
	fmt.Printf("Payload: %v\n", payload)

	var slackAction slack.SlackAction
	err := json.Unmarshal([]byte(payload), &slackAction)
	if err != nil {
		return "", err
	}
	fmt.Printf("SlackAction: %v\n", slackAction)

	//dbErr := dynamodb.UpdateLastCoffee(slackAction.User.Name)
	//if dbErr != nil {
	//	return "", dbErr
	//}

	itemBytes, err := json.Marshal(slack.Message{
		Channel: "#coffee",
		Text:    "FRESH COFFEE!! UPDATE ",
	})
	if err != nil {
		return "", err
	}
	return string(itemBytes), nil
}

func main() {
	lambda.Start(Handler)
}
