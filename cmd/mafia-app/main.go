package main

import (
	"encoding/json"
	"log"
	"mafia-app/api"
	"math/rand"
	"net/http"
	"time"
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
	http.HandleFunc("/mafia", mainHander)
	http.HandleFunc("/get-user-id", api.GetUserID)
	http.HandleFunc("/create", api.CreateRoom)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
