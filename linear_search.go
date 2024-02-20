package DSA_Go

func LinearSearch(array []int, target int) bool {
	for _, value := range array {
		if value == target {
			return true
		}
	}
	return false
}
