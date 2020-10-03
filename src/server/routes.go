package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	// HdrRequestID is request header used for tracing requests
	HdrRequestID = "Request-Tracing-ID"
)

func (s *server) RegisterRoutes() {
	s.router.HandleFunc("/messages", s.handleGetMessages()).Methods(http.MethodGet)
	s.router.HandleFunc("/messages", s.handlePostMessage()).Methods(http.MethodPost)
	s.router.HandleFunc("/messages/{id}", s.handleGetSingleMessage()).Methods(http.MethodGet)
	s.router.HandleFunc("/messages/{id}", s.handleDeleteMessage()).Methods(http.MethodDelete)
}

// Logging middleware logs the incoming requests
func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now().UTC()
			defer func() {
				requestID, ok := r.Context().Value(HdrRequestID).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Printf("%s: %s  Method: %s URL: %s RemoteAddr: %s UserAgent: %s Latency: %v ", HdrRequestID, requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Since(t))
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// Tracing middleware adds a TracingID to each requests
func Tracing(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = fmt.Sprintf("%d", time.Now().UnixNano())
			}
			ctx := context.WithValue(r.Context(), HdrRequestID, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
