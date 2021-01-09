package main

import (
	"encoding/json"
	"log"
	"mafia-app/api"
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type responseJSON struct {
	Title   string
	Payload string
}

func mainHander(w http.ResponseWriter, request *http.Request) {
	out := &responseJSON{Title: "Main", Payload: "Project Root"}
	resp, err := json.Marshal(out)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/mafia", mainHander)
	mux.HandleFunc("/get-user-id", api.GetUserID)
	mux.HandleFunc("/create", api.CreateRoom)
	mux.HandleFunc("/join", api.JoinRoom)
	mux.HandleFunc("/ready", api.SetReady)
	mux.HandleFunc("/status", api.GetStatus)
	handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatal(err)
	}
}
