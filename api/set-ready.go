package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func setReadyValidator(w http.ResponseWriter, data url.Values) (rCode bool) {
	rCode = true
	if _, ok := data["id"]; !ok {
		http.Error(w, "User id not found", http.StatusBadRequest)
		rCode = false
	} else {
		userID, err := strconv.Atoi(data["id"][0])
		if err != nil {
			http.Error(w, "User id can only be integer in range 0-999", http.StatusBadRequest)
			rCode = false
		}
		ok, player := getPlayerByID(userID, &PlayersInfo)
		if !ok {
			http.Error(w, "Player not found", http.StatusBadRequest)
			rCode = false
		} else if !player.InRoom {
			http.Error(w, "You are not in room", http.StatusBadRequest)
			rCode = false
		}
	}
	return
}

// SetReady allows user to change his readiness status
func SetReady(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "POST":
		data := request.URL.Query()
		if !setReadyValidator(w, data) {
			return
		}
		userID, _ := strconv.Atoi(data["id"][0])
		_, player := getPlayerByID(userID, &PlayersInfo)
		_, playerInRoom := getPlayerInRoom(userID, player.RoomID, &Rooms)
		playerInRoom.Ready = true
		fmt.Fprintf(w, "{'success':'true'}")
		return
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Only POST method is allowed")
	}
}
