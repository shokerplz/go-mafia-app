package main

import (
	"encoding/json"
	"log"
	"mafia-app/api"
	"net/http"
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
	http.HandleFunc("/mafia", mainHander)
	http.HandleFunc("/get-user-id", api.GetUserID)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
