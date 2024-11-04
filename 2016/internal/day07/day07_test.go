package day07

import (
	"testing"
)

func TestSupportTSL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "abba[mnop]qrst", expected: true},
		{input: "abcd[bddb]xyyx", expected: false},
		{input: "aaaa[qwer]tyui", expected: false},
		{input: "ioxxoj[asdfgh]zxcvbn", expected: true},
		{input: "qrst[mnop]abba", expected: true},
	}

	for _, test := range tests {
		ip := IPv7(test.input)
		result := ip.SupportsTSL()
		if result != test.expected {
			t.Errorf("IPv7.SupportsTSL(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestSupportSSL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "aba[bab]xyz", expected: true},
		{input: "xyx[xyx]xyx", expected: false},
		{input: "aaa[kek]eke", expected: true},
		{input: "zazbz[bzb]cdb", expected: true},
	}

	for _, test := range tests {
		ip := IPv7(test.input)
		result := ip.SupportsSSL()
		if result != test.expected {
			t.Errorf("IPv7.SupportsSSL(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
