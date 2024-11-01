package day05

import (
	"testing"
)

func TestGeneratePassword1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "abc", expected: "18f47a30"},
	}

	g := NewGen1()
	for _, test := range tests {
		resultNaive := GeneratePasswordNaive(&g, test.input)
		if resultNaive != test.expected {
			t.Errorf("GeneratePasswordNaive(1, %v) = %s, want %s", test.input, resultNaive, test.expected)
		}
		resultFancy := GeneratePasswordFancy(&g, test.input)
		if resultFancy != resultNaive {
			t.Errorf("GeneratePasswordFancy(1, %v) = %s, want %s", test.input, resultFancy, test.expected)
		}
	}
}

func TestGeneratePassword2(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "abc", expected: "05ace8e3"},
	}

	g := NewGen2()
	for _, test := range tests {
		resultNaive := GeneratePasswordNaive(&g, test.input)
		if resultNaive != test.expected {
			t.Errorf("GeneratePasswordNaive(2, %v) = %s, want %s", test.input, resultNaive, test.expected)
		}
		resultFancy := GeneratePasswordFancy(&g, test.input)
		if resultFancy != resultNaive {
			t.Errorf("GeneratePasswordFancy(2, %v) = %s, want %s", test.input, resultFancy, test.expected)
		}
	}
}
