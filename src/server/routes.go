package server

import (
	"context"
	"net/http"
	"time"

	"github.com/amanviitb/Qlik/src/logger"
	"github.com/google/uuid"
)

const (
	tracingID = "Request-Tracing-ID"
)

type RequestTracing string

// RegisterRoutes register the endpoints for the service to receive requests on
func (s *server) RegisterRoutes() {
	s.router.HandleFunc("/messages", s.handleGetMessages()).Methods(http.MethodGet)
	s.router.HandleFunc("/messages", s.handlePostMessage()).Methods(http.MethodPost)
	s.router.HandleFunc("/messages/{id}", s.handleGetSingleMessage()).Methods(http.MethodGet)
	s.router.HandleFunc("/messages/{id}", s.handleDeleteMessage()).Methods(http.MethodDelete)
}

// Logging middleware logs all the incoming requests
func Logging(l logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now().UTC()
			defer func() {
				requestID, ok := r.Context().Value(RequestTracing(tracingID)).(string)
				if !ok {
					requestID = "unknown"
				}
				l.Info("%s: %s  Method: %s URL: %s RemoteAddr: %s UserAgent: %s Latency: %v ", tracingID, requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Since(t))
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// Tracing middleware adds a TracingID to each request
func Tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				// requestID = fmt.Sprintf("%d", time.Now().UnixNano())
				requestID = uuid.New().String()
			}
			ctx := context.WithValue(r.Context(), RequestTracing(tracingID), requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
