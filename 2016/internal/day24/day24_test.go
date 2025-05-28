package day24

import (
	"testing"
)

func TestPart1(t *testing.T) {
	content := `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`
	tests := []struct {
		content  string
		expected int
	}{
		{content: content, expected: 14},
	}

	for _, test := range tests {
		m, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent() failed prematurly")
			continue
		}
		result := part1(m)
		if result != test.expected {
			t.Errorf("part1() = %v; want %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	content := `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`
	tests := []struct {
		content  string
		expected int
	}{
		{content: content, expected: 20},
	}

	for _, test := range tests {
		m, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent() failed prematurly")
			continue
		}
		result := part2(m)
		if result != test.expected {
			t.Errorf("part2() = %v; want %v", result, test.expected)
		}
	}
}
