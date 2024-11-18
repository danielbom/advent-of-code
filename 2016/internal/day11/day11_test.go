package day11

import (
	"testing"
)

func TestPart1(t *testing.T) {
	content1 := `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`
	tests := []struct {
		content  string
		expected int
	}{
		{content: content1, expected: 11},
	}

	for _, test := range tests {
		floors, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
			return
		}
		result := part1(floors)
		if result != test.expected {
			t.Errorf("part1(%v) = %v; want %v", test.content, result, test.expected)
		}
	}
}
