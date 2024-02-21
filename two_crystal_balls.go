package DSA_Go

import "math"

func TwoCrystalBalls(array []bool) int {
	if array[0] {
		return -1
	}
	jmp := int(math.Sqrt(float64(len(array))))
	ptr := 0
	for i := 0; i < len(array); i += jmp {
		if array[i] {
			break
		}
		ptr = i
	}
	for i := ptr; i < len(array); i++ {
		ptr = i
		if array[i] {
			break
		}
	}
	return ptr
}
