package server

import (
	"net/http"
	"strings"

	"github.com/amanviitb/Qlik/src/data"
	"github.com/gorilla/mux"
)

// handleGetMessages is the handler for GET request to fetch all messages
// GET /messages
func (s *server) handleGetMessages() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		s.logger.Info(req.Method, "Get All Messages")
		messages := data.GetMessages()
		err := data.ToJSON(messages, rw)
		if err != nil {
			s.logger.Error("Unable to serializing message", err)
		}
	}
}

// POST /messages
func (s *server) handlePostMessage() http.HandlerFunc {
	// add logic to generate an ID, create a message object and assign the ID to it
	return func(rw http.ResponseWriter, req *http.Request) {
		message := new(data.Message)
		err := data.FromJSON(message, req.Body)
		if err != nil {
			s.logger.Error("Unable to deserialize message", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		data.AddMessage(message)
		rw.WriteHeader(http.StatusCreated)
	}
}

// GET /messages/{id}
func (s *server) handleGetSingleMessage() http.HandlerFunc {
	// a separate response for message
	type response struct {
		MessageText  string `json:"messageText"`
		IsPalindrome bool   `json:"isPalindrome"`
	}

	return func(rw http.ResponseWriter, req *http.Request) {
		// parse the request to fetch the id from the URI
		pathVars := mux.Vars(req)
		messageID := pathVars["id"]

		// assuming each message gets its unique ID
		err := data.DeleteMessageWithID(messageID)

		switch err {
		case nil:
		case data.ErrMessageNotFound:
			s.logger.Error("Unable to fetch message", "error: ", err)
			http.Error(rw, "No message found with the given ID", http.StatusNotFound)
			return
		default:
			s.logger.Error("Unable to fetch message", "error: ", err)
			http.Error(rw, "", http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusNoContent)
	}
}

// DELETE /messages/{id}
func (s *server) handleDeleteMessage() http.HandlerFunc {
	// check to see if the id exists, if not send http.StatusNotFound
	// else delete the resource and if error occurs return http.StatusInternalServerError
	// else return http.Deleted status
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}

// checkIfPalindrome checks to see if the given string is a palindrome
func checkIfPalindrome(s string) bool {
	// an empty string or a string with a single character is always a palindrome
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	// two indices to keep track of each character from left and right
	var left, right = 0, len(s) - 1
	for left <= right {
		if strings.ToLower(string(s[right])) < string('a') ||
			strings.ToLower(string(s[right])) > string('z') { // jump over the non-alphabet character from right
			right--
		} else if strings.ToLower(string(s[left])) < string('a') ||
			strings.ToLower(string(s[left])) > string('z') { // jump over the non-alphabet character from left
			left++
		} else if strings.ToLower(string(s[right])) !=
			strings.ToLower(string(s[left])) { // characters at both ends don't match
			return false
		} else { // characters at both ends match
			left++
			right--
		}
	}
	return true
}
