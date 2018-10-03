package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	// handler
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	// sample request
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		log.Fatal(err)
	}

	// new recorder
	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
}
