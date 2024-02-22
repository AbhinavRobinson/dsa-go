package DSA_Go

import (
	"testing"
)

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	testArray := []int{9, 3, 7, 4, 69, 420, 42}
	sortedArray := []int{3, 4, 7, 9, 42, 69, 420}
	var tests = []struct {
		name  string
		input []int
		want  []int
	}{
		{"Array should be sorted", testArray, sortedArray},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := BubbleSort(tt.input)
			if !IntArrayEquals(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
