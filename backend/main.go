package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const CLIENT_ID = "6df37e2ce30640edbc0b73a5098a7fdc"
const CLIENT_SECRET = "b81634579fb34d49b6def119d8189b72"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", searchHandler).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "OPTIONS"})

	fmt.Println("Server running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query param 'q'", http.StatusBadRequest)
		return
	}

	token, err := getAppAccessToken()
	if err != nil {
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		log.Println("Token error:", err)
		return
	}

	searchURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=10", url.QueryEscape(query))
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "Spotify search failed", http.StatusInternalServerError)
		log.Println("Spotify search error:", err)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

func getAppAccessToken() (string, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(CLIENT_ID + ":" + CLIENT_SECRET))
	body := strings.NewReader("grant_type=client_credentials")

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return "", fmt.Errorf("token request failed: %v", err)
	}
	defer resp.Body.Close()

	var res struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	return res.AccessToken, nil
}
