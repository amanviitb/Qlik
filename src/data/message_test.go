package data

import (
	"reflect"
	"testing"
)

func TestAddMessage(t *testing.T) {
	messages = Messages{&Message{ID: 1, Text: "Message 1", Sender: "S"}}
	currentLength := len(messages)
	defer func() {
		messages = Messages{}
	}()
	type args struct {
		msg *Message
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Add a new message to the Messages collection",
			args: args{msg: &Message{Text: "New Message", Sender: "T"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddMessage(tt.args.msg)
			if len(messages) != currentLength+1 {
				t.Error("error")
			}
		})
	}
}

func TestGetMessageByID(t *testing.T) {
	messages = Messages{&Message{ID: 1, Text: "Message 1", Sender: "S"}}
	defer func() {
		messages = Messages{}
	}()
	type args struct {
		messageID int
	}
	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		{
			name:    "return message if found and nil as error",
			args:    args{messageID: 1},
			want:    messages[0],
			wantErr: false,
		},
		{
			name:    "return nil if no found and ErrMessageNotFound as error",
			args:    args{messageID: 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMessageByID(tt.args.messageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMessageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMessageByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMessageWithID(t *testing.T) {
	messages = Messages{&Message{ID: 1, Text: "Message 1", Sender: "S"}}
	defer func() {
		messages = Messages{}
	}()
	type args struct {
		messageID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "length of messages collection should decrease by one",
			args:    args{messageID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteMessageWithID(tt.args.messageID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMessageWithID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(messages) != 0 {
				t.Error("error")
			}
		})
	}
}
