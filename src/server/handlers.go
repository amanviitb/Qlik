package server

import (
	"net/http"
	"strings"

	"github.com/amanviitb/Qlik/src/data"
	"github.com/gorilla/mux"
)

// handleGetMessages is the handler for GET request to fetch all messages
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

func (s *server) handlePostMessage() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		message := new(data.Message)
		err := data.FromJSON(message, req.Body)
		if err != nil {
			s.logger.Error("Unable to deserialize message", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		rw.WriteHeader(http.StatusCreated)
	}
}

func (s *server) handleGetSingleMessage() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// parse the request to fetch the id from the URI
		pathVars := mux.Vars(req)
		messageID, ok := pathVars["id"]
		if !ok {
			http.Error(rw, "Unknown query parameter passed", http.StatusBadRequest)
			return
		}
		s.logger.Info(messageID)
	}
}

func (s *server) handleDeleteMessage() http.HandlerFunc {
	// check to see if the id exists, if not send http.StatusNotFound
	// else delete the resource and if error occurs return http.StatusInternalServerError
	// else return http.Deleted status
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}

func checkIfPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	var left, right = 0, len(s) - 1
	for left <= right {
		if strings.ToLower(string(s[right])) < string('a') || strings.ToLower(string(s[right])) > string('z') {
			right--
		} else if strings.ToLower(string(s[left])) < string('a') || strings.ToLower(string(s[left])) > string('z') {
			left++
		} else if strings.ToLower(string(s[right])) != strings.ToLower(string(s[left])) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
