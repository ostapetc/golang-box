package slices

func Contains(arr []string, search string) bool {
	for _, val := range arr {
		if val == search {
			return true
		}
	}

	return false
}

func ContainsAll(arr []string, search []string) bool {
	for _, val := range search {
		if !Contains(arr, val) {
			return false
		}
	}

	return true
}

func ContainsAny(arr []string, search []string) bool {
	for _, val := range search {
		if Contains(arr, val) {
			return true
		}
	}

	return false
}

func ContainsInt32(arr []int32, search int32) bool {
	for _, val := range arr {
		if val == search {
			return true
		}
	}

	return false
}

func ContainsAllInt32(arr []int32, search []int32) bool {
	for _, val := range search {
		if !ContainsInt32(arr, val) {
			return false
		}
	}

	return true
}
