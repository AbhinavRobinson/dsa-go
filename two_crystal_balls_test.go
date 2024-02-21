package DSA_Go

import (
	"math/rand"
	"testing"
)

// Helper to create N array with same value
func RepeatedSlice[T any](elem T, size int) []T {
	s := make([]T, size)
	for i := range s {
		s[i] = elem
	}
	return s
}

func TestTwoCrystalBalls(t *testing.T) {
	// 1 - Array with target within range
	idx := rand.Intn(10000)
	a1 := RepeatedSlice(true, 10000)
	for i := range a1 {
		if i == idx {
			break
		}
		a1[i] = false
	}

	// 2 - Array with no target
	a2 := RepeatedSlice(true, 821)

	var tests = []struct {
		name  string
		input []bool
		want  int
	}{
		{"index should exist", a1, idx},
		{"index should not exist", a2, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := TwoCrystalBalls(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
