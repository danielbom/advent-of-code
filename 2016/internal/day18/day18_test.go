package day18

import (
	"testing"
)

func TestNextRow(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "..^^.", expected: ".^^^^"},
		{input: ".^^^^", expected: "^^..^"},
	}

	for _, test := range tests {
		row, err := NewRow(test.input)
		if err != nil {
			t.Errorf("NewRow(%v) failed prematurely", test.input)
			continue
		}
		result := row.NextRow().String()
		if test.expected != result {
			t.Errorf("Row.NextRow(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCountSafeTiles(t *testing.T) {
	tests := []struct {
		input    string
		rows     int
		expected int
	}{
		{input: ".^^.^.^^^^", rows: 10, expected: 38},
	}

	for _, test := range tests {
		row, err := NewRow(test.input)
		if err != nil {
			t.Errorf("NewRow(%v) failed prematurely", test.input)
			continue
		}
		result := countSafeTiles(row, test.rows)
		if test.expected != result {
			t.Errorf("countSafeTiles(%v, %v) = %v; want %v", test.input, test.rows, result, test.expected)
		}
	}
}
