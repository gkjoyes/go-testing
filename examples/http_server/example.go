package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	// test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Client")
	}))
	defer ts.Close()

	// reqeust to test server
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	// read response body
	greeting, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// display
	fmt.Printf("%s", greeting)
}
