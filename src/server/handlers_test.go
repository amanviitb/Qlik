package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/amanviitb/Qlik/src/logger"
	"github.com/gorilla/mux"
)

func Test_server_handleDeleteMessage(t *testing.T) {
	type fields struct {
		router *mux.Router
		logger logger.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				router: tt.fields.router,
				logger: tt.fields.logger,
			}
			if got := s.handleDeleteMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.handleDeleteMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_handleGetSingleMessage(t *testing.T) {
	type fields struct {
		router *mux.Router
		logger logger.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				router: tt.fields.router,
				logger: tt.fields.logger,
			}
			if got := s.handleGetSingleMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.handleGetSingleMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_handlePostMessage(t *testing.T) {
	var buffer = []byte(`{"text": "Amore, roma", "sender":"Steve"}`)
	s := server{router: mux.NewRouter(), logger: logger.GetLogger()}
	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewBuffer(buffer))
	w := httptest.NewRecorder()
	handler := s.handlePostMessage()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Error("error")
	}
}

func Test_server_handleGetMessages(t *testing.T) {
	s := server{router: mux.NewRouter(), logger: logger.GetLogger()}
	req := httptest.NewRequest(http.MethodGet, "/messages", nil)
	w := httptest.NewRecorder()
	handler := s.handleGetMessages()
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Error("error occured")
	}
}
