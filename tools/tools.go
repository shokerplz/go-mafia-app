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
