package test_rest

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from the dummy server"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/metrics", handler)
	http.ListenAndServe(":8080", nil)
}
