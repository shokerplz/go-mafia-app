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
