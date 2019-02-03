package made

import (
	"encoding/json"
	"fmt"
)

//Message properties
type Message struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	IconEmoji   string       `json:"icon_emoji"`
	ImageURL    string       `json:"image_url"`
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
