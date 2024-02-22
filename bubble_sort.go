package DSA_Go

func BubbleSort(array []int) []int {
	max := len(array)
	for max > 0 {
		for index := 0; index < max-1; index++ {
			if array[index] > array[index+1] {
				swap := array[index+1]
				array[index+1] = array[index]
				array[index] = swap
			}
		}
		max--
	}
	return array
}
