package main

import (
	"io"
	"log"
	"net/http"
)

/*
Handler function declaration should be of the form func(w http.ResponseWriter, r *http.Request)
*/
// Our handleRoot function will serve the resources/html/index.html file that you can fill with anything you want and print out information about the client who sent the request
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Print data about the user
	log.Printf("IP and Port: %v\n", r.RemoteAddr)
	// Serve the file
	http.ServeFile(w, r, "resources/html/index.html")

}

// Our handlePost function will handle any http POST requests by printing them to the console
func handlePost(w http.ResponseWriter, r *http.Request) {
	// Transform the response body into a []bytes
	responseBytes, err := io.ReadAll(r.Body)
	// Handle errors by printing them, sending an http.Error and returning early
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	// Turning the responseBytes into a string
	responseString := string(responseBytes)
	// Printing responseString
	log.Println(responseString)
	// Writing a 200 OK header response
	w.WriteHeader(http.StatusOK) //http.StatusOK is equal to 200
	return
}
