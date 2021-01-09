package api

func getPlayerByID(id int, a *[]playerInfoJSON) (bool, *playerInfoJSON) {
	i := 0
	for _, item := range *a {
		if item.ID == id {
			return true, &((*a)[i])
		}
		i++
	}
	return false, nil
}

func getRoomByID(id int, a *[]roomJSON) (bool, *roomJSON) {
	i := 0
	for _, item := range *a {
		if item.ID == id {
			return true, &((*a)[i])
		}
		i++
	}
	return false, nil
}

func getPlayerInRoom(userID int, roomID int, a *[]roomJSON) (bool, *playerInRoomJSON) {
	i := 0
	ok, room := getRoomByID(roomID, a)
	if !ok {
		return false, nil
	}
	for _, user := range room.Users {
		if user.ID == userID {
			return true, &(room.Users[i])
		}
		i++
	}
	return false, nil
}
