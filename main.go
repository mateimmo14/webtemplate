package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new ServeMux to handle request
	mux := http.NewServeMux()
	// Handle a GET request to / with a handleRoot function declared in handlers.go
	mux.HandleFunc("GET /", handleRoot)
	// Handle a POST request to  / (or anything that is not already handled) with a handlePost function also declared in handlers.go
	mux.HandleFunc("POST /", handlePost)
	// Handle the images as a FileServer
	mux.Handle("GET /images/", http.StripPrefix("/images", http.FileServer(http.Dir("resources/images"))))
	// Listen for any requests on port 8080 and print any errors (http.ListenAndServe returns an error)
	log.Println(http.ListenAndServe(":8080", mux))
}
