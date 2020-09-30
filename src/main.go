package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var (
	// Info logger is to log Info type logs
	Info *log.Logger
	// Warning logger is to log Warning type logs
	Warning *log.Logger
	// Error logger is to log Error type logs
	Error *log.Logger
)

func InitLogger(infoHandle, warningHandle, errorHandle io.Writer) {
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// ph := handlers.NewProducts(l, v, db)
	// create a new serve mux and register the handlers
	InitLogger(os.Stdout, os.Stdout, os.Stderr)
	r := mux.NewRouter()
	router := r.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/messages", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Messages Exist"))
	}).Methods(http.MethodGet)

	// handlers for API
	// getR := sm.Methods(http.MethodGet).Subrouter()
	// getR.HandleFunc("/products", ph.ListAll).Queries("currency", "{[A-Z]{3}}")
	// getR.HandleFunc("/products", ph.ListAll)

	// getR.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle).Queries("currency", "{[A-Z]{3}}")
	// getR.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	// putR := sm.Methods(http.MethodPut).Subrouter()
	// putR.HandleFunc("/products", ph.Update)
	// putR.Use(ph.MiddlewareValidateProduct)

	// postR := sm.Methods(http.MethodPost).Subrouter()
	// postR.HandleFunc("/products", ph.Create)
	// postR.Use(ph.MiddlewareValidateProduct)

	// deleteR := sm.Methods(http.MethodDelete).Subrouter()
	// deleteR.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	// handler for documentation
	// opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	// sh := middleware.Redoc(opts, nil)

	// getR.Handle("/docs", sh)
	// getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	// ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         ":9091",           // configure the bind address
		Handler:      r,                 // set the default handler
		ErrorLog:     Error,             // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		Info.Println("Starting server on port 9091")

		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			Error.Println("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	Info.Println("Gracefully shutting down the server....", sig)

	// gracefully shutdown the server, waiting max 30 seconds for active connections to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
