package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amanviitb/Qlik/src/logger"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	logger logger.Logger
}

var _ Server = (*server)(nil)

// NewServer returns an instance of server configured with logger and router
func NewServer() Server {
	return &server{
		router: mux.NewRouter(),
		logger: logger.GetLogger(),
	}
}

func (srv *server) Start() {
	srv.RegisterRoutes()
	s := http.Server{
		Addr:    ":9091",                                    // configure the bind address
		Handler: Tracing()(Logging(srv.logger)(srv.router)), // set the default handler
		// ErrorLog:     Error,             // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		srv.logger.Info("Starting server on port 9091")

		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			srv.logger.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	srv.logger.Info("Gracefully shutting down the server....", sig)

	// gracefully shutdown the server, waiting max 30 seconds for active connections to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
