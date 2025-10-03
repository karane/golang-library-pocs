package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"text"`
}

type UUIDResponse struct {
	UUID string `json:"uuid"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Message{Text: "(Karane) Hello, Docker + Go + Air + Gorilla/Mux + UUID! 🚀"})
}

func uuidHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	json.NewEncoder(w).Encode(UUIDResponse{UUID: id.String()})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods("GET")
	r.HandleFunc("/uuid", uuidHandler).Methods("GET")

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
