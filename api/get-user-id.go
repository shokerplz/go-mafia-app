package api

import (
	"encoding/json"
	"fmt"
	"mafia-app/tools"
	"math/rand"
	"net/http"
)

type userIdJSON struct {
	USER_ID string
}

// GetUserID generates new user id
func GetUserID(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i := 0
	var userID int
	for {
		userID = rand.Intn(1000)
		i++
		if !tools.InIntArray(userID, Players) {
			Players = append(Players, userID)
			break
		}
		if i >= 1000 {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Game server can not handle your request")
			return
		}
	}
	out := &userIdJSON{USER_ID: fmt.Sprint(userID)}
	resp, err := json.Marshal(out)
	if err == nil {
		w.Write(resp)
	}
}
