package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Length int `json:"length"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	word := strings.TrimPrefix(
		r.URL.Path,
		"/strings/length/",
	)

	result := Response{
		Length: len([]rune(word)),
	}

	json.NewEncoder(w).Encode(result)
}

func main() {

	http.HandleFunc(
		"/strings/length/",
		handler,
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
