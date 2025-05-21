package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://secretchat-ocgsqd14x-vamshi1188s-projects.vercel.app") // <-- change this
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func encodeMessage(msg string) string {
	encoded := ""
	for _, ch := range msg {
		encoded += fmt.Sprintf("%03d", int(ch))
	}
	return encoded
}

func decodeMessage(code string) string {
	decoded := ""
	for i := 0; i < len(code); i += 3 {
		part := code[i : i+3]
		num, err := strconv.Atoi(part)
		if err == nil {
			decoded += string(rune(num))
		}
	}
	return decoded
}

func handleEncode(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	msg := r.FormValue("message")
	encoded := encodeMessage(msg)
	fmt.Fprint(w, encoded)
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	code := r.FormValue("code")
	decoded := decodeMessage(code)
	fmt.Fprint(w, decoded)
}

func main() {
	http.HandleFunc("/encode", handleEncode)
	http.HandleFunc("/decode", handleDecode)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
