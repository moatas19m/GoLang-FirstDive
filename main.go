package main

import (
	"fmt"
	"net/http"
)

// ExampleHandler demonstrates a custom handler
type ExampleHandler struct{}

// ServeHTTP implements the http.Handler interface
func (h *ExampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World (whispered from inside the ExampleHandler) :D This is http.Handle()")
}

// Simpler handler to be used by HandleFunc
func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World, this is http.HandleFunc()")
}

func main() {
	// Custom handler using http.Handle
	myCustomHandler := &ExampleHandler{}
	http.Handle("/custom", myCustomHandler)

	// Simple handler using http.HandleFunc
	http.HandleFunc("/simple", simpleHandler)

	// Start the server
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
