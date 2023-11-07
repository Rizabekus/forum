package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (controllers *Controllers) Googleauth(w http.ResponseWriter, r *http.Request) {
	clientID := "432111077942-8gih8tkb7o1pt672t11280rdev9jk1cd.apps.googleusercontent.com"
	redirectURI := "http://localhost:8000/googlecallback"
	url := fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=email", clientID, redirectURI)
	http.Redirect(w, r, url, http.StatusFound)
}

func (controllers *Controllers) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	clientID := "432111077942-8gih8tkb7o1pt672t11280rdev9jk1cd.apps.googleusercontent.com"
	redirectURI := "http://localhost:8000/googlecallback"
	clientSecret := "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"
	code := r.URL.Query().Get("code")

	// Exchange the code for an access token
	tokenURL := "https://accounts.google.com/o/oauth2/token"
	data := fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code", code, clientID, clientSecret, redirectURI)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tokenResponse map[string]interface{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken := tokenResponse["access_token"].(string)

	// Now you can use the accessToken to make authenticated requests to Google APIs
	// You can use the accessToken in the Authorization header or as a query parameter in your requests.
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		// Handle the error, e.g., log it or return an error response.
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			// Handle the error, e.g., log it or return an error response.
		}

		// Convert the response body to a string and print it
		responseContent := string(body)
		fmt.Println(responseContent)
	} else {
		// Handle non-200 status codes as needed
	}
}
