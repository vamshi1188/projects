package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/instagram"
)

var (
	clientID     = "YOUR_CLIENT_ID"
	clientSecret = "YOUR_CLIENT_SECRET"
	redirectURL  = "http://localhost:8080/callback"
	scopes       = []string{"user_profile", "user_media"}

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     instagram.Endpoint,
	}
)

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/callback", handleCallback)
	fmt.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := oauth2Config.Client(context.Background(), token)
	response, err := client.Get("https://graph.instagram.com/me?fields=id,username,media_count,account_type")
	if err != nil {
		http.Error(w, "Failed to get user profile: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Print the response body to standard output
	fmt.Fprintf(w, "Profile Information: ")
	if _, err := os.Stdout.ReadFrom(response.Body); err != nil {
		http.Error(w, "Failed to read response body: "+err.Error(), http.StatusInternalServerError)
	}
}
