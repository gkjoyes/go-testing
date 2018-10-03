package main

import (
	"log"
	"net/http"

	"github.com/g-kutty/go-testing/examples/coverage/handlers"
)

func main() {

	// initialize routers
	handlers.Routes()

	// listen and serve
	log.Println("listener: Started: Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
