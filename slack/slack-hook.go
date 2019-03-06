package slack

import (
	"encoding/json"
	"fmt"
)

//Message properties
type Message struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	ImageURL    string       `json:"image_url,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
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
	for _, leader := range leaders {
		leaderBoardText += "```" + leader + "````\n"

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
