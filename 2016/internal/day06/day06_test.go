package day06

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input1 := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	tests := []struct {
		content  string
		expected string
	}{
		{content: input1, expected: "easter"},
	}

	for _, test := range tests {
		input, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part1(input)
		if result != test.expected {
			t.Errorf("part1(%v) = %s, want %s", test.content, result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	input1 := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	tests := []struct {
		content  string
		expected string
	}{
		{content: input1, expected: "advent"},
	}

	for _, test := range tests {
		input, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
		}
		result := part2(input)
		if result != test.expected {
			t.Errorf("part2(%v) = %s, want %s", test.content, result, test.expected)
		}
	}
}
