package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func joinRoomValidator(w http.ResponseWriter, data url.Values) (rCode bool) {
	rCode = true
	if _, ok := data["id"]; !ok {
		http.Error(w, "Room id not found", http.StatusBadRequest)
		rCode = false
	} else {
		roomID, err := strconv.Atoi(data["id"][0])
		if err != nil {
			http.Error(w, "Room id can only be integer in range 0-100", http.StatusBadRequest)
			rCode = false
		}
		ok, room := getRoomByID(roomID, &Rooms)
		if !ok {
			http.Error(w, "There is no such room", http.StatusBadRequest)
			rCode = false
		} else if room.MaxUsers <= len(room.Users) {
			http.Error(w, "Room is full", http.StatusBadRequest)
			rCode = false
		}
		if _, ok := data["user_id"]; !ok {
			http.Error(w, "User id not found", http.StatusBadRequest)
			rCode = false
		} else {
			userID, err := strconv.Atoi(data["user_id"][0])
			if err != nil {
				http.Error(w, "User id can only be integer in range 0-999", http.StatusBadRequest)
				rCode = false
			}
			ok, user := getPlayerByID(userID, &PlayersInfo)
			if !ok {
				http.Error(w, "Player not found", http.StatusBadRequest)
				rCode = false
			} else if user.InRoom {
				http.Error(w, "You are already in room", http.StatusBadRequest)
				rCode = false
			}
		}
	}
	return
}

// JoinRoom allows user to join game room
func JoinRoom(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "POST":
		data := request.URL.Query()
		if !joinRoomValidator(w, data) {
			return
		}
		userID, _ := strconv.Atoi(data["user_id"][0])
		roomID, _ := strconv.Atoi(data["id"][0])
		playerInRoom := playerInRoomJSON{
			ID:           userID,
			Alive:        true,
			VotesAgainst: 0,
			Role:         "peaceful",
			Ready:        false,
		}
		_, room := getRoomByID(roomID, &Rooms)
		_, player := getPlayerByID(userID, &PlayersInfo)
		room.Users = append(room.Users, playerInRoom)
		player.InRoom = true
		player.RoomID = roomID
		if room.MaxUsers <= len(room.Users) {
			room.State = "game"
		}
		if room.MaxUsers <= len(room.Users) {
			room.State = "game"
			go game(room)
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
