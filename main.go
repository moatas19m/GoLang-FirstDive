package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 1. Serve static files from the "/static" route
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	// 2. RedirectHandler
	// Redirect users visiting "/google" to "https://www.google.com"
	http.Handle("/google", http.RedirectHandler("https://www.google.com", http.StatusMovedPermanently))

	// 3. NotFoundHandler
	// Use the default NotFoundHandler for "/notfound"
	http.Handle("/notfound", http.NotFoundHandler())

	// 4. TimeoutHandler
	// Wrap a slow handler with a timeout of 2 seconds
	slowHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second) // Simulate a slow response
		w.Write([]byte("This response took too long!"))
	})
	http.Handle("/slow", http.TimeoutHandler(slowHandler, 2*time.Second, "Request timed out"))

	// 5. Custom catch-all fallback for unmatched routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/notfound", http.StatusSeeOther)
	})

	// Start the server on port 8080
	fmt.Println("Starting server on :9000...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
