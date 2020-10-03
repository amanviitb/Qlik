package data

import (
	"time"
)

// Message is the structure for a message
type Message struct {
	ID     string    `json:"id"`
	Text   string    `json:"text"`
	Sender string    `json:"sender"`
	Time   time.Time `json:"time,omitempty"`
}

// Messages is the collection of all messages
type Messages []*Message

var messages Messages = Messages{
	&Message{ID: "123", Text: "Olo", Sender: "Su"},
	&Message{ID: "1122", Text: "Hello", Sender: "Sam"},
}

// GetMessages returns all the messages
func GetMessages() Messages {
	return messages
}

// improve logging
// add versioning to the APIs like /api/v1
// add some validation to the fields of a message
// dockerfile and creating image
// test cases
// palindromic string
// observed”(Monitoring/Traceability/metrics)
