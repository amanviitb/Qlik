package data

import "errors"

// ErrMessageNotFound is the error that is returned when there is not matching message
var ErrMessageNotFound = errors.New("no message found with the given ID")

// Message is the structure for a message
type Message struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Sender string `json:"sender"`
	Time   string `json:"-"`
}

// Messages is the collection of all messages
type Messages []*Message

var messages Messages = Messages{
	&Message{ID: "123", Text: "Olo", Sender: "Stu"},
	&Message{ID: "1122", Text: "Hello", Sender: "Sam"},
}

// GetMessages returns all the messages
func GetMessages() Messages {
	return messages
}

// AddMessage adds the message to the list of messages
func AddMessage(msg *Message) {
	messages = append(messages, msg)
}

// GetMessageByID returns a message for a given messageID
func GetMessageByID(messageID string) (*Message, error) {
	for i := range messages {
		if messages[i].ID == messageID {
			return messages[i], nil
		}
	}
	return nil, ErrMessageNotFound
}

// DeleteMessageWithID deletes a message with the given ID
func DeleteMessageWithID(messageID string) error {
	var indexToDelete = -1
	for i := range messages {
		if messageID == messages[i].ID {
			indexToDelete = i
			break
		}
	}
	// no message with the given ID was found
	if indexToDelete == -1 {
		return ErrMessageNotFound
	}
	messages = append(messages[:indexToDelete], messages[indexToDelete+1])
	return nil
}

// improve logging
// add versioning to the APIs like /api/v1
// add some validation to the fields of a message
// dockerfile and creating image instructions
// test cases
// palindromic string - done
// observed”(Monitoring/Traceability/metrics)
