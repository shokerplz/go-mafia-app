package api

import (
	"mafia-app/tools"
	"math"
	"math/rand"
	"time"
)

var bFactor float64 = 3

type actionGame struct {
	vote int
	kill int
}

var actions = actionGame{vote: 0, kill: 1}

func game(room *roomJSON) {
	var blackPlayers []playerInRoomJSON
	var redPlayers []playerInRoomJSON
	for !getPlayersReady(room) {
		time.Sleep(250 * time.Millisecond)
	}
	time.Sleep(2500 * time.Millisecond)
	bAmount := math.Round(float64(room.MaxUsers) / bFactor)
	bIndexes := rand.Perm(len(room.Users))[:int(bAmount)]
	for _, idx := range bIndexes {
		room.Users[idx].Role = "mafia"
	}
	i := 0
	for _, player := range room.Users {
		switch player.Role {
		case "peaceful":
			room.Peaceful = append(room.Peaceful, player.ID)
			redPlayers = append(redPlayers, player)
		case "mafia":
			room.Mafia = append(room.Mafia, player.ID)
			blackPlayers = append(blackPlayers, player)
		}
		room.Users[i].Alive = true
		room.Users[i].Ready = false
		room.Alive = append(room.Alive, player.ID)
		i++
	}
	room.Daytime = "night"
	room.Cicle = 1
	for _, idx := range bIndexes {
		if room.Users[idx].Ready == false {
			time.Sleep(250 * time.Millisecond)
		}
	}
	time.Sleep(2500 * time.Millisecond)
	for {
		ended, whoWon := checkIfGameEnded(room)
		if !ended {
			for i := 0; i < len(room.Users); i++ {
				room.Users[i].Ready = false
			}
			room.Daytime = "day"
			room.Cicle++
			action(room, actions.vote)
			time.Sleep(3000 * time.Millisecond)
			room.Daytime = "night"
			time.Sleep(1000 * time.Millisecond)
			action(room, actions.kill)
		} else {
			room.State = "ended won: " + whoWon
			break
		}
	}
	time.Sleep(5000 * time.Millisecond)
	room = nil
	return
}

func checkIfGameEnded(room *roomJSON) (ended bool, whoWon string) {
	if len(room.Mafia) <= 0 {
		ended, whoWon = true, "peaceful"
	} else if len(room.Peaceful) <= 0 {
		ended, whoWon = true, "mafia"
	} else if len(room.Mafia) == 1 && len(room.Peaceful) == 1 {
		ended, whoWon = true, "mafia"
	} else {
		ended, whoWon = false, ""
	}
	return
}

func action(room *roomJSON, action int) {
	if len(room.Mafia) == 0 {
		return
	}
	room.State = "vote"
	var count *int
	var maxCount int
	if action == actions.vote {
		count = &room.Voted
		maxCount = len(room.Alive)
	} else if action == actions.kill {
		count = &room.VotedToKill
		maxCount = len(room.Mafia)
	}
	for *count < maxCount {
		time.Sleep(250 * time.Millisecond)
	}
	votesMax := 0
	var votesMaxUserIndex int
	var votesMaxUserIndexRole int
	for i := 0; i < len(room.Users); i++ {
		if room.Users[i].VotesAgainst >= votesMax {
			votesMax = room.Users[i].VotesAgainst
			votesMaxUserIndex = i
		}
		room.Users[i].VotesAgainst = 0
	}
	if action == actions.vote {
		room.Jailed = append(room.Jailed, room.Users[votesMaxUserIndex].ID)
	} else if action == actions.kill {
		room.Killed = append(room.Killed, room.Users[votesMaxUserIndex].ID)
	}
	room.Alive = tools.RemoveItemFromArray(room.Alive, votesMaxUserIndex)
	switch room.Users[votesMaxUserIndex].Role {
	case "mafia":
		for i := 0; i < len(room.Mafia); i++ {
			if room.Mafia[i] == room.Users[votesMaxUserIndex].ID {
				votesMaxUserIndexRole = i
				break
			}
		}
		room.Mafia = tools.RemoveItemFromArray(room.Mafia, votesMaxUserIndexRole)
	case "peaceful":
		for i := 0; i < len(room.Peaceful); i++ {
			if room.Peaceful[i] == room.Users[votesMaxUserIndex].ID {
				votesMaxUserIndexRole = i
				break
			}
		}
		room.Peaceful = tools.RemoveItemFromArray(room.Peaceful, votesMaxUserIndexRole)
	}
	room.Voted = 0
	for i := 0; i < len(room.Users); i++ {
		room.Users[i].Ready = false
	}
}
