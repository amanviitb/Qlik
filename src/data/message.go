package data

// Message is the structure for a message
type Message struct {
	ID     string 		`json:"id"`
	Text   string 		`json:"text"`
	Sender string 		`json:"sender"`
}

// Messages is the collection of all messages
type Messages []*Message

func GetMessages() Messages {
	return Messages{
		&Message{ID:"123", Text: "Olo", Sender: "Su"}
	}
}