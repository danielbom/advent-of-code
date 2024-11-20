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
		resultSequencial := GeneratePasswordSequencial(&g, test.input)
		if resultSequencial != test.expected {
			t.Errorf("GeneratePasswordSequencial(1, %v) = %s, want %s", test.input, resultSequencial, test.expected)
		}
		resultParallel := GeneratePasswordParallel(&g, test.input)
		if resultParallel != resultSequencial {
			t.Errorf("GeneratePasswordParallel(1, %v) = %s, want %s", test.input, resultParallel, test.expected)
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
		resultSequencial := GeneratePasswordSequencial(&g, test.input)
		if resultSequencial != test.expected {
			t.Errorf("GeneratePasswordSequencial(2, %v) = %s, want %s", test.input, resultSequencial, test.expected)
		}
		resultParallel := GeneratePasswordParallel(&g, test.input)
		if resultParallel != resultSequencial {
			t.Errorf("GeneratePasswordParallel(2, %v) = %s, want %s", test.input, resultParallel, test.expected)
		}
	}
}
