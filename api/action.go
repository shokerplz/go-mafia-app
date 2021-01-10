package api

import (
	"fmt"
	"mafia-app/tools"
	"net/http"
	"net/url"
	"strconv"
)

func actionValidator(w http.ResponseWriter, data url.Values) (rCode bool) {
	rCode = true
	var room *roomJSON
	var playerInRoom *playerInRoomJSON
	if _, ok := data["user_id"]; !ok {
		http.Error(w, "User id not found", http.StatusBadRequest)
		rCode = false
		return
	} else {
		userID, err := strconv.Atoi(data["user_id"][0])
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
		} else if ok, room = getRoomByID(player.RoomID, &Rooms); !ok {
			http.Error(w, "No such room", http.StatusBadRequest)
			rCode = false
		} else if ok, playerInRoom = getPlayerInRoom(userID, player.RoomID, &Rooms); !ok {
			http.Error(w, "You are not in this room", http.StatusBadRequest)
			rCode = false
		} else if playerInRoom.Ready {
			http.Error(w, "You already voted", http.StatusBadRequest)
			rCode = false
		}
	}
	if _, ok := data["target_id"]; !ok {
		http.Error(w, "Target id not found", http.StatusBadRequest)
		rCode = false
		return
	} else {
		targetID, err := strconv.Atoi(data["target_id"][0])
		if err != nil {
			http.Error(w, "Target id can only be integer in range 0-999", http.StatusBadRequest)
			rCode = false
		}
		ok, player := getPlayerByID(targetID, &PlayersInfo)
		if !ok {
			http.Error(w, "Target not found", http.StatusBadRequest)
			rCode = false
		} else if !player.InRoom {
			http.Error(w, "Target is not in room", http.StatusBadRequest)
			rCode = false
		} else if ok, _ := getRoomByID(player.RoomID, &Rooms); !ok {
			http.Error(w, "No such room", http.StatusBadRequest)
			rCode = false
		} else if ok, targetInRoom := getPlayerInRoom(targetID, player.RoomID, &Rooms); !ok {
			http.Error(w, "Target is not in this room", http.StatusBadRequest)
			rCode = false
		} else if targetInRoom.Alive == false {
			http.Error(w, "Target is not alive", http.StatusBadRequest)
			rCode = false
		} else if !tools.InIntArray(targetID, room.Alive) {
			http.Error(w, "Target is not alive", http.StatusBadRequest)
			rCode = false
		}
	}
	if action, ok := data["action"]; !ok {
		http.Error(w, "Action not found", http.StatusBadRequest)
		rCode = false
		return
	} else {
		if action[0] != "vote" && action[0] != "kill" {
			http.Error(w, "Wrong action", http.StatusBadRequest)
			rCode = false
		} else if action[0] == "vote" && room.State != "vote" {
			http.Error(w, "You can not vote now", http.StatusBadRequest)
			rCode = false
		} else if action[0] == "kill" && room.Daytime != "night" {
			http.Error(w, "You can not vote now", http.StatusBadRequest)
			rCode = false
		} else if (action[0] == "kill" && playerInRoom.Role != "mafia") ||
			(action[0] == "kill" && !tools.InIntArray(playerInRoom.ID, room.Mafia)) {
			http.Error(w, "You are not mafia", http.StatusBadRequest)
			rCode = false
		}
	}
	return
}

// Action allows user to vote and kill
func Action(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "POST":
		data := request.URL.Query()
		if !actionValidator(w, data) {
			return
		}
		actionToDo := data["action"][0]
		userID, _ := strconv.Atoi(data["user_id"][0])
		targetID, _ := strconv.Atoi(data["target_id"][0])
		_, player := getPlayerByID(userID, &PlayersInfo)
		_, playerInRoom := getPlayerInRoom(userID, player.RoomID, &Rooms)
		_, target := getPlayerInRoom(targetID, player.RoomID, &Rooms)
		_, room := getRoomByID(player.RoomID, &Rooms)
		target.VotesAgainst++
		if actionToDo == "vote" {
			room.Voted++
		}
		if actionToDo == "kill" {
			room.VotedToKill++
		}
		playerInRoom.Ready = true
		w.Write([]byte(`{"success":"true"}`))
		return
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Only POST method is allowed")
	}
}
