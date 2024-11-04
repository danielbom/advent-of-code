package day10

import (
	"testing"
)

func TestParseInstruction(t *testing.T) {
	tests := []struct {
		input    string
		expected Instruction
	}{
		{input: "value 5 goes to bot 2", expected: NewInstructionTake(5, 2)},
		{input: "bot 2 gives low to bot 1 and high to bot 0", expected: NewInstructionGive(2, 1, 0, BOT, BOT)},
		{input: "bot 1 gives low to output 1 and high to bot 0", expected: NewInstructionGive(1, 1, 0, OUTPUT, BOT)},
		{input: "bot 0 gives low to output 2 and high to output 0", expected: NewInstructionGive(0, 2, 0, OUTPUT, OUTPUT)},
	}

	p := NewParseInstruction()
	for _, test := range tests {
		result, err := p.Parse(test.input)
		if err != nil {
			t.Errorf("ParseInstruction.Parse(%v) = error '%v', wants value %v", test.input, err, test.expected)
		}
		if result != test.expected {
			t.Errorf("ParseInstruction.Parse(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}

// 0
// 1 3
// 2 2 5 (*)
// --
// 0 5
// 1 2 3 (*)
// 2
// --
// 0 3 5 (*)
// 1
// 2
// out1 2
// --
// 0
// 1
// 2
// out0 5
// out2 3
func TestFindBot(t *testing.T) {
	input := `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`
	tests := []struct {
		input       string
		compareLow  int
		compareHigh int
		expected    int
	}{
		{input: input, compareLow: 2, compareHigh: 5, expected: 2},
		{input: input, compareLow: 2, compareHigh: 3, expected: 1},
		{input: input, compareLow: 3, compareHigh: 5, expected: 0},
	}

	p := NewParseInstruction()
	for _, test := range tests {
		is, err := p.ParseLines(test.input)
		if err != nil {
			t.Errorf("ParseInstruction.ParseLines(%v) = error '%v', wants value %v", test.input, err, test.expected)
			continue
		}
		result := findBot(is, test.compareLow, test.compareHigh)
		if result != test.expected {
			t.Errorf("findBot(%v, %v, %v) = %v, want %v", test.input, test.compareLow, test.compareHigh, result, test.expected)
		}
	}
}

func TestMultiplyFirstsOutputs(t *testing.T) {
	input := `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`
	tests := []struct {
		input    string
		expected int
	}{
		{input: input, expected: 30},
	}

	p := NewParseInstruction()
	for _, test := range tests {
		is, err := p.ParseLines(test.input)
		if err != nil {
			t.Errorf("ParseInstruction.ParseLines(%v) = error '%v', wants value %v", test.input, err, test.expected)
			continue
		}
		result := multiplyFirstsOutputs(is)
		if result != test.expected {
			t.Errorf("multiplyFirstsOutputs(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
