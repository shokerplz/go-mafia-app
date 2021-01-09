package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func getStatusValidator(w http.ResponseWriter, data url.Values) (rCode bool) {
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

// GetStatus Allows user to see check status in real time
func GetStatus(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "GET":
		data := request.URL.Query()
		if !getStatusValidator(w, data) {
			return
		}
		userID, _ := strconv.Atoi(data["id"][0])
		_, player := getPlayerByID(userID, &PlayersInfo)
		_, room := getRoomByID(player.RoomID, &Rooms)
		resp, err := json.Marshal(&room)
		if err == nil {
			w.Write(resp)
		}
		return
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Only GET method is allowed")
	}

}
