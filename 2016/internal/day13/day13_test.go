package day13

import (
	"testing"
)

func TestCountStepsTo(t *testing.T) {
	tests := []struct {
		seed     int
		x        int
		y        int
		expected int
	}{
		{seed: 10, x: 7, y: 4, expected: 11},
	}

	for _, test := range tests {
		result := countSteps(test.seed, test.x, test.y)
		if result != test.expected {
			t.Errorf("countStepsTo(%v, %v, %v) = %v; want = %v", test.seed, test.x, test.y, result, test.expected)
		}
	}
}
