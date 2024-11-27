package main

import (
	"fmt"
	"net/http"
)

//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello World"))
//	})
//	http.ListenAndServe(":8080", nil)
//}

type ExampleHandler struct{}

func (h *ExampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World (whispered from inside the ExampleHandler) :D This is http.Handle()")
}

func main() {
	myHandler := &ExampleHandler{}
	http.Handle("/", myHandler)
	err := http.ListenAndServe(":8080", &ExampleHandler{})
	if err != nil {
		return
	}
}
