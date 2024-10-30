package day03

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input1 := `101 102 103
201 202 203
301 302 303
401 402 403
501 502 503
601 602 603`

	tests := []struct {
		content  string
		expected int
	}{
		{content: input1, expected: 6},
	}

	for _, test := range tests {
		triangles, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part1(triangles)
		if result != test.expected {
			t.Errorf("part1(%v) = %d, want %d", test.content, result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	input1 := `101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603`

	tests := []struct {
		content  string
		expected int
	}{
		{content: input1, expected: 6},
	}

	for _, test := range tests {
		triangles, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part2(triangles)
		if result != test.expected {
			t.Errorf("part2(%v) = %d, want %d", test.content, result, test.expected)
		}
	}
}
