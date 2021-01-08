package api

// Players contain all players ids
var Players []int

// PlayersInfo contain info about all players
var PlayersInfo []playerInfoJSON

// Rooms contain all info about in-game rooms
var Rooms []roomJSON

type userIDJSON struct {
	UserID int `json:"USER_ID"`
}

type playerInfoJSON struct {
	ID     int  `json:"id"`
	InRoom bool `json:"in_room"`
	RoomID int  `json:"room_id"`
}

type playerInRoomJSON struct {
	ID           int    `json:"id"`
	VotesAgainst int    `json:"votes_against"`
	Alive        bool   `json:"alive"`
	Role         string `json:"role"`
	Ready        bool   `json:"ready"`
}

type roomJSON struct {
	ID          int                `json:"id"`
	Users       []playerInRoomJSON `json:"users"`
	MaxUsers    int                `json:"max_users"`
	State       string             `json:"state"`
	Daytime     string             `json:"daytime"`
	Alive       []int              `json:"alive"`
	Killed      []int              `json:"killed"`
	Jailed      []int              `json:"jailed"`
	Mafia       []int              `json:"mafia"`
	Peaceful    []int              `json:"peaceful"`
	Cicle       string             `json:"cicle"`
	Voted       int                `json:"voted"`
	VotedToKill int                `json:"voted_to_kill"`
}
