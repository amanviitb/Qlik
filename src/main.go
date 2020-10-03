package main

import (
	"github.com/amanviitb/Qlik/src/server"
)

func main() {
	// we could also load from config files
	// set up all the dependencies of the server by calling NewServer
	s := server.NewServer()
	// start the server
	s.Start()
}
