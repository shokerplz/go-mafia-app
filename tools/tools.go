package tools

// InIntArray Checks if variable is in array (first variable is what to find, second is array)
func InIntArray(find int, a []int) bool {
	for _, b := range a {
		if find == b {
			return true
		}
	}
	return false
}

// GetItemFromIntArray Gets item from integer array
func GetItemFromIntArray(find int, a *[]int) (bool, *int) {
	i := 0
	for _, item := range *a {
		if find == item {
			return true, &((*a)[i])
		}
		i++
	}
	return false, nil
}

// GetItemIndex Gets item index from integer array
func GetItemIndex(find int, a *[]int) (index int) {
	i := 0
	for _, item := range *a {
		if find == item {
			return i
		}
		i++
	}
	return -1
}

// RemoveItemFromArray Removes item from array with reslicing
func RemoveItemFromArray(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
