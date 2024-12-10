package day23

import (
	"testing"
)

func TestPart1(t *testing.T) {
	content := `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`
	tests := []struct {
		content  string
		expected int
	}{
		{content: content, expected: 3},
	}

	for _, test := range tests {
		is, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent() failed prematurly")
			continue
		}
		result := part1(is)
		if result != test.expected {
			t.Errorf("part1() = %v; want %v", result, test.expected)
		}
	}
}
