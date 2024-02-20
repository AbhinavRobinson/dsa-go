package DSA_Go

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	testArray := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	var tests = []struct {
		name  string
		input int
		want  bool
	}{
		{"69 should be true", 69, true},
		{"1336 should be false", 1336, false},
		{"69420 should be true", 69420, true},
		{"69421 should be false", 69421, false},
		{"1 should be true", 1, true},
		{"0 should be false", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := LinearSearch(testArray, tt.input)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
