package day09

import (
	"testing"
)

func TestDecompressLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "ADVENT", expected: 6},
		{input: "A(1x5)BC", expected: 7},
		{input: "(3x3)XYZ", expected: 9},
		{input: "A(2x2)BCD(2x2)EFG", expected: 11},
		{input: "(6x1)(1x3)A", expected: 6},
		{input: "X(8x2)(3x3)ABCY", expected: 18},
	}

	d := NewD()
	for _, test := range tests {
		result := d.DecompressLength(test.input)
		if result != test.expected {
			t.Errorf("part1(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestDecompress(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "ADVENT", expected: "ADVENT"},
		{input: "A(1x5)BC", expected: "ABBBBBC"},
		{input: "(3x3)XYZ", expected: "XYZXYZXYZ"},
		{input: "A(2x2)BCD(2x2)EFG", expected: "ABCBCDEFEFG"},
		{input: "(6x1)(1x3)A", expected: "(1x3)A"},
		{input: "X(8x2)(3x3)ABCY", expected: "X(3x3)ABC(3x3)ABCY"},
	}

	d := NewD()
	for _, test := range tests {
		result := d.Decompress(test.input)
		if result != test.expected {
			t.Errorf("part1(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestDeepDecompressLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "ADVENT", expected: 6},
		{input: "A(1x5)BC", expected: 7},
		{input: "(3x3)XYZ", expected: 9},
		{input: "A(2x2)BCD(2x2)EFG", expected: 11},
		{input: "(6x1)(1x3)A", expected: 3},
		{input: "X(8x2)(3x3)ABCY", expected: 20},
		{input: "X(9x2)(3x3)ABCY", expected: 21},
		{input: "(27x12)(20x12)(13x14)(7x10)(1x12)A", expected: 241920},
		{input: "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", expected: 445},
	}

	d := NewD()
	for _, test := range tests {
		result := d.DeepDecompressLength(test.input)
		if result != test.expected {
			t.Errorf("part1(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
