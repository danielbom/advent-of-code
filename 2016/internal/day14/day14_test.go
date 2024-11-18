package day14

import (
	"testing"
)

func TestCollectHashPattern(t *testing.T) {
	tests := []struct {
		input string
		size3 int
		size5 int
	}{
		{input: "aaaaa", size3: int('a'), size5: int('a')},
		{input: "aaabbbbb", size3: int('a'), size5: int('b')},
		{input: "aaabbbbbccc", size3: int('a'), size5: int('b')},
		{input: "aaabbbbbcccddddd", size3: int('a'), size5: int('b')},
		{input: "xyaaaaa", size3: int('a'), size5: int('a')},
		{input: "xyaaabbbbb", size3: int('a'), size5: int('b')},
		{input: "xyaaabbbbbccc", size3: int('a'), size5: int('b')},
		{input: "xyaaabbbbbcccddddd", size3: int('a'), size5: int('b')},
	}

	for _, test := range tests {
		result := CollectHashPattern(test.input)
		if result.Size3 != test.size3 || result.Size5 != test.size5 {
			t.Errorf("CollectHashPatter(%v) = (%v, %v); want = (%v, %v)", test.input, result.Size3, result.Size5, test.size3, test.size5)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		salt     string
		expected int
	}{
		{salt: "abc", expected: 22728},
	}

	for _, test := range tests {
		result := part1(test.salt)
		if result != test.expected {
			t.Errorf("part1(%v) = %v; want = %v", test.salt, result, test.expected)
		}
	}
}

func TestMD5Hash(t *testing.T) {
	tests := []struct {
		input    string
		stretch  int
		expected string
	}{
		{input: "abc0", stretch: 0, expected: "577571be4de9dcce85a041ba0410f29f"},
		{input: "abc0", stretch: 1, expected: "eec80a0c92dc8a0777c619d9bb51e910"},
		{input: "abc0", stretch: 2, expected: "16062ce768787384c81fe17a7a60c7e3"},
		{input: "abc0", stretch: 2016, expected: "a107ff634856bb300138cac6568c0f24"},
	}

	for _, test := range tests {
		result := MD5Hash(test.input, test.stretch)
		if result != test.expected {
			t.Errorf("MD5Hash(%v, %v) = %v; want = %v", test.input, test.stretch, result, test.expected)
		}
	}
}
