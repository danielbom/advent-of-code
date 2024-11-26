package day21

import (
	"reflect"
	"testing"
)

func TestInstructions(t *testing.T) {
	tests := []struct {
		input    string
		inst     Instruction
		expected string
	}{
		{input: "abc", inst: SwapPosition{X: 0, Y: 2}, expected: "cba"},
		{input: "abc", inst: SwapLetter{A: 'a', B: 'c'}, expected: "cba"},
		{input: "abcde", inst: SwapPosition{X: 0, Y: 4}, expected: "ebcda"},
		{input: "ebcda", inst: SwapLetter{A: 'd', B: 'b'}, expected: "edcba"},
		{input: "abc", inst: RotateLeft{Steps: 0}, expected: "abc"},
		{input: "abc", inst: RotateLeft{Steps: 1}, expected: "bca"},
		{input: "abcde", inst: RotateLeft{Steps: 1}, expected: "bcdea"},
		{input: "abc", inst: RotateRight{Steps: 0}, expected: "abc"},
		{input: "abc", inst: RotateRight{Steps: 1}, expected: "cab"},
		{input: "abcde", inst: RotateRight{Steps: 1}, expected: "eabcd"},
		{input: "abdec", inst: RotateOnLetter{A: 'b'}, expected: "ecabd"},
		{input: "ecabd", inst: RotateOnLetter{A: 'd'}, expected: "decab"},
		{input: "ebcda", inst: ReversePosition{X: 0, Y: 4}, expected: "adcbe"},
		{input: "abcde", inst: ReversePosition{X: 1, Y: 3}, expected: "adcbe"},
		{input: "bcdea", inst: MovePosition{X: 1, Y: 4}, expected: "bdeac"},
		{input: "bdeac", inst: MovePosition{X: 3, Y: 0}, expected: "abdec"},
	}

	for i, test := range tests {
		result := ApplyInstruction(test.inst, test.input)
		if result != test.expected {
			ty := reflect.TypeOf(test.inst)
			t.Errorf("%d: %v%v.Apply(%v) = %v; want = %v", i, ty.Name(), test.inst, test.input, result, test.expected)
			continue
		}
	}
}

func TestApplyInstructions(t *testing.T) {
	instructions := `swap position 4 with position 0
  swap letter d with letter b
  reverse positions 0 through 4
  rotate left 1 step
  move position 1 to position 4
  move position 3 to position 0
  rotate based on position of letter b
  rotate based on position of letter d`
	input := "abcde"
	expected := "decab"
	is, err := parseContent(instructions)
	if err != nil {
		t.Errorf("parseContent() failed prematurly")
		return
	}
	result := applyInstructions(is, input)
	if result != expected {
		t.Errorf("applyInstructions(%v) = %v; want %v", input, result, expected)
	}
}
