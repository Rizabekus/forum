package controllers

import (
	"fmt"
	"net/http"
)

func (controllers *Controllers) Googleauth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controllers.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	// Parse the form values from the request
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form values:", err)
		w.WriteHeader(http.StatusBadRequest) // Return a bad request response
		return
	}

	// Access the user data
	id := r.FormValue("id")
	name := r.FormValue("name")
	imageUrl := r.FormValue("imageUrl")
	email := r.FormValue("email")

	// Use the user data as needed
	fmt.Printf("Received data from Google Sign-In:\nID: %s\nName: %s\nImageURL: %s\nEmail: %s\n",
		id, name, imageUrl, email)

	// You can now process the data, authenticate the user, or perform other actions as needed.

	// Send a response to the client, indicating success or failure, if necessary.
}
