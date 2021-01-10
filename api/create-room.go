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

func createRoomValidator(w http.ResponseWriter, data url.Values) (rCode bool) {
	rCode = true
	if _, ok := data["user_id"]; !ok {
		http.Error(w, "User id not found", http.StatusBadRequest)
		rCode = false
	} else {
		userID, err := strconv.Atoi(data["user_id"][0])
		if err != nil {
			http.Error(w, "User id can only be integer in range 0-999", http.StatusBadRequest)
			rCode = false
		}
		if !tools.InIntArray(userID, Players) {
			http.Error(w, "Player not found", http.StatusBadRequest)
			rCode = false
		} else {
			if ok, player := getPlayerByID(userID, &PlayersInfo); ok {
				if player.InRoom == true {
					http.Error(w, "You are already in room", http.StatusBadRequest)
					rCode = false
				}
			}
		}
	}
	if _, ok := data["users"]; !ok {
		http.Error(w, "Max users not set", http.StatusBadRequest)
		rCode = false
	} else {
		_, err := strconv.Atoi(data["users"][0])
		if err != nil {
			http.Error(w, "Max users can only be integer in range 0-128", http.StatusBadRequest)
			rCode = false
		}
	}
	return
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
			roomID := rand.Intn(99) + 1
			if ok, _ := getRoomByID(roomID, &Rooms); !ok {
				room.ID = roomID
				room.MaxUsers = maxUsers
				room.State = "joining"
				room.Daytime = "day"
				room.Mafia = make([]int, 0)
				room.Peaceful = make([]int, 0)
				room.Alive = make([]int, 0)
				room.Killed = make([]int, 0)
				room.Jailed = make([]int, 0)
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
