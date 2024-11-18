package day15

import (
	"testing"
)

func TestPart1(t *testing.T) {
	content1 := `Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.`
	tests := []struct {
		content  string
		expected int
	}{
		{content: content1, expected: 5},
	}

	for _, test := range tests {
		ds, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part1(ds)
		if result != test.expected {
			t.Errorf("part1(%s) = %v; want %v", test.content, result, test.expected)
		}
	}
}
