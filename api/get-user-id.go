package api

import (
	"encoding/json"
	"fmt"
	"mafia-app/tools"
	"math/rand"
	"net/http"
)

// GetUserID generates new user id
func GetUserID(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "GET":
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	case "POST":
		i := 0
		var userID int
		for {
			userID = rand.Intn(1000)
			i++
			if !tools.InIntArray(userID, Players) {
				Players = append(Players, userID)
				PlayersInfo = append(PlayersInfo, playerInfoJSON{ID: userID})
				break
			}
			if i >= 1000 {
				http.Error(w, "503 Game is not available", http.StatusServiceUnavailable)
				return
			}
		}
		out := &userIDJSON{UserID: userID}
		resp, err := json.Marshal(out)
		if err == nil {
			w.Write(resp)
		}
		return
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Only POST method is allowed")
	}

}
