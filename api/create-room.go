package api

import (
	"encoding/json"
	"fmt"
	"mafia-app/tools"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func createRoomValidator(w http.ResponseWriter, data url.Values) bool {
	if _, ok := data["user_id"]; !ok {
		http.Error(w, "User id not found", http.StatusBadRequest)
		return false
	} else {
		userID, err := strconv.Atoi(data["user_id"][0])
		if err != nil {
			http.Error(w, "User id can only be integer in range 0-999", http.StatusBadRequest)
			return false
		}
		if !tools.InIntArray(userID, Players) {
			http.Error(w, "Player not found", http.StatusBadRequest)
			return false
		} else {
			if ok, player := getPlayerByID(userID, &PlayersInfo); ok {
				if player.InRoom == true {
					http.Error(w, "You are already in room", http.StatusBadRequest)
					return false
				}
			}
		}
	}
	if _, ok := data["users"]; !ok {
		http.Error(w, "Max users not set", http.StatusBadRequest)
		return false
	} else {
		_, err := strconv.Atoi(data["users"][0])
		if err != nil {
			http.Error(w, "Max users can only be integer in range 0-128", http.StatusBadRequest)
			return false
		}
	}
	return true
}

// CreateRoom creates new game room
func CreateRoom(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "POST":
		data := request.URL.Query()
		if !createRoomValidator(w, data) {
			return
		}
		var room roomJSON
		maxUsers, _ := strconv.Atoi(data["users"][0])
		i := 0
		for {
			i++
			roomID := rand.Intn(100)
			if ok, _ := getRoomByID(roomID, &Rooms); !ok {
				room.ID = roomID
				room.MaxUsers = maxUsers
				room.State = "joining"
				Rooms = append(Rooms, room)
				break
			}
			if i >= 100 {
				http.Error(w, "503 Game is not available", http.StatusServiceUnavailable)
				return
			}
		}
		resp, err := json.Marshal(&room)
		if err == nil {
			w.Write(resp)
		}
		return
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Only POST method is allowed")
	}
}
