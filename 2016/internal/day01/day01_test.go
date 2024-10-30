package day01

import (
	"testing"
)

func TestParseInstruction(t *testing.T) {
	tests := []struct {
		input    string
		expected Instruction
		hasError bool
	}{
		{input: "R5", expected: Instruction{Turn: RIGHT, Count: 5}, hasError: false},
		{input: "L3", expected: Instruction{Turn: LEFT, Count: 3}, hasError: false},
		{input: "X10", expected: Instruction{}, hasError: true},
	}

	for _, test := range tests {
		result, err := parseInstruction(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("parseInstruction(%s) error = %v; want error = %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("parseInstruction(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		content  string
		expected int
	}{
		{content: "R2, L3", expected: 5},
		{content: "R2, R2, R2", expected: 2},
		{content: "R5, L5, R5, R3", expected: 12},
	}

	for _, test := range tests {
		instructions, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part1(instructions)
		if result != test.expected {
			t.Errorf("part1(%s) = %v; want %v", test.content, result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		content  string
		expected int
	}{
		{content: "R8, R4, R4, R8", expected: 4},
	}

	for _, test := range tests {
		instructions, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part2(instructions)
		if result != test.expected {
			t.Errorf("part2(%s) = %v; want %v", test.content, result, test.expected)
		}
	}
}
