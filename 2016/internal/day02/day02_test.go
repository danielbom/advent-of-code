package day02

import (
	"testing"
)

func TestPart1(t *testing.T) {
	content1 := `ULL
RRDDD
LURDL
UUUUD`
	tests := []struct {
		content  string
		expected string
	}{
		{content: content1, expected: "1985"},
	}

	for _, test := range tests {
		iss, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part1(iss)
		if result != test.expected {
			t.Errorf("part1(%s) = %v; want %v", test.content, result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	content1 := `ULL
RRDDD
LURDL
UUUUD`
	tests := []struct {
		content  string
		expected string
	}{
		{content: content1, expected: "5DB3"},
	}

	for _, test := range tests {
		iss, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part2(iss)
		if result != test.expected {
			t.Errorf("part1(%s) = %v; want %v", test.content, result, test.expected)
		}
	}
}
