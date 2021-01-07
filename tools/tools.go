package tools

// InIntArray Checks if variable find is in array a
func InIntArray(find int, a []int) bool {
	for _, b := range a {
		if find == b {
			return true
		}
	}
	return false
}
