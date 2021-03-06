package slack

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//SlackAPIResponse Slack API Response object
type SlackAPIResponse struct {
	Channel   string  `json:"channel"`
	TimeStamp string  `json:"ts"`
	Message   Message `json:"message"`
}

//SlackAction properties
type SlackAction struct {
	CallBackId      string  `json:"callback_id"`
	OriginalMessage Message `json:"original_message"`
	User            User    `json:"user"`
	Channel         User    `json:"channel"`
	MessageTS       string  `json:"message_ts"`
}

//User properties
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//Message properties
type Message struct {
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	ImageURL    string       `json:"image_url,omitempty"`
	Username    string       `json:"username,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

//Attachment properties
type Attachment struct {
	Fields   []AttachmentFields `json:"fields"`
	ImageURL string             `json:"image_url,omitempty"`
}

//AttachmentFields properties
type AttachmentFields struct {
	Title string `json:"title"`
	Short bool   `json:"short"`
}

func HandleMadeCoffeeEvent(channel string, username string) (string, error) {
	itemBytes, err := json.Marshal(Message{
		Channel:   "#" + channel,
		Text:      "Coffee made by " + username,
		IconEmoji: ":star2:",
	})
	fmt.Printf("itemJSON: %v\n", string(itemBytes))
	if err != nil {
		return "", err
	}
	return string(itemBytes), nil
}

func HandleIOTEvent() (string, error) {
	itemBytes, err := json.Marshal(Message{
		Channel:   "#richmondcoffee",
		Text:      "FRESH COFFEE!! - ",
		IconEmoji: ":coffee:",
	})
	fmt.Printf("itemJSON: %v\n", string(itemBytes))
	if err != nil {
		return "", err
	}
	return string(itemBytes), nil
}

func HandleLeaderBoard(channel string, leaders []string) (string, error) {
	leaderBoardText := "*Coffee All Stars* \n "
	for index, leader := range leaders {
		leaderBoardText += "```" + strconv.Itoa(index+1) + ". " + leader + "```\n"

	}
	itemBytes, err := json.Marshal(Message{
		Channel:   "#" + channel,
		Text:      leaderBoardText,
		IconEmoji: ":star2:",
	})
	fmt.Printf("itemJSON: %v\n", string(itemBytes))
	if err != nil {
		return "", err
	}
	return string(itemBytes), nil
}
