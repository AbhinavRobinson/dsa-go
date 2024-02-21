package DSA_Go

func BinarySearch(array []int, target int) bool {
	low, high := 0, len(array)
	for high > low {
		mid := low + (high-low)/2
		if array[mid] == target {
			return true
		}
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
