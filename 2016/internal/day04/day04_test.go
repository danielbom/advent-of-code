package day04

import (
	"testing"
)

func TestParseEncryptedName(t *testing.T) {
	tests := []struct {
		input    string
		expected EncryptedName
	}{
		{input: "aaaaa-bbb-z-y-x-123[abxyz]", expected: NewEncryptedName("aaaaa-bbb-z-y-x", 123, "abxyz")},
		{input: "a-b-c-d-e-f-g-h-987[abcde]", expected: NewEncryptedName("a-b-c-d-e-f-g-h", 987, "abcde")},
	}

	for _, test := range tests {
		result, err := parseEncryptedName(test.input)
		if err != nil {
			t.Errorf("parseEncryptedName(%v) = error '%v', want value %v", test.input, err, test.expected)
		}
		if result != test.expected {
			t.Errorf("parseEncryptedName(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestEncryptedNameIsReal(t *testing.T) {
	tests := []struct {
		content  string
		expected bool
	}{
		{content: "aaaaa-bbb-z-y-x-123[abxyz]", expected: true},
		{content: "a-b-c-d-e-f-g-h-987[abcde]", expected: true},
		{content: "not-a-real-room-404[oarel]", expected: true},
		{content: "totally-real-room-200[decoy]", expected: false},
	}

	for _, test := range tests {
		e, err := parseEncryptedName(test.content)
		if err != nil {
			t.Errorf("parseEncryptedName(%s) failed prematurely", test.content)
		}
		result := e.IsReal()
		if result != test.expected {
			t.Errorf("EncryptedName.IsReal(%v) = %v, want %v", test.content, result, test.expected)
		}
	}
}

func TestShiftCipher(t *testing.T) {
	tests := []struct {
		input    string
		shift    int
		expected string
	}{
		{input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", shift: int('x') - int('a'), expected: "XYZABCDEFGHIJKLMNOPQRSTUVW"},
		{input: "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", shift: int('x') - int('a'), expected: "QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD"},
		{input: "qzmt-zixmtkozy-ivhz", shift: 343, expected: "very encrypted name"},
	}

	for _, test := range tests {
		result := ShiftCipher(test.input, test.shift)
		if result != test.expected {
			t.Errorf("ShiftCipher(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
