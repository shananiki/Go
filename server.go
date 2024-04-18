package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the route for serving index.html
	http.HandleFunc("/go/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the index.html file
		http.ServeFile(w, r, "index.html")
	})

	// Start the server
	port := ":3005" // You can change the port if needed
	fmt.Printf("Server is running on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
