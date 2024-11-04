package day08

import (
	"reflect"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	tests := []struct {
		input    string
		expected Instruction
	}{
		{input: "rect 3x2", expected: NewInstructionFillRect(2, 3)},
		{input: "rotate column x=1 by 1", expected: NewInstructionRotateCol(1, 1)},
		{input: "rotate row y=0 by 4", expected: NewInstructionRotateRow(0, 4)},
		{input: "rotate column x=1 by 1", expected: NewInstructionRotateCol(1, 1)},
	}

	for _, test := range tests {
		result, err := parseInstruction(test.input)
		if err != nil {
			t.Errorf("parseInstruction(%v) = error '%v', wants value %v", test.input, err, test.expected)
			continue
		}
		if result != test.expected {
			t.Errorf("parseInstruction(%v) = %v, wants %v", test.input, result, test.expected)
		}
	}
}

func TestFillRect(t *testing.T) {
	grid1 := NewGrid(3, 3)
	grid1.LitUp(0, 0)
	grid1.LitUp(1, 0)
	grid1.LitUp(0, 1)
	grid1.LitUp(1, 1)
	tests := []struct {
		input    Grid
		row      int
		col      int
		expected Grid
	}{
		{input: NewGrid(3, 3), row: 2, col: 2, expected: grid1},
	}

	for _, test := range tests {
		result := NewGrid(test.input.rows, test.input.cols)
		result.FillRect(test.row, test.col)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Grid.FillRect(row=%d, col=%d) = %v, wants %v", test.row, test.col, result.String(), test.expected.String())
		}
	}
}

func copyGrid(g Grid) Grid {
	r := NewGrid(g.rows, g.cols)
	copy(r.grid, g.grid)
	return r
}

func TestRotateRow(t *testing.T) {
	base := NewGrid(3, 3)
	base.LitUp(0, 0)
	base.LitUp(0, 1)
	base.LitUp(1, 0)
	base.LitUp(1, 1)
	grid1 := NewGrid(3, 3)
	grid1.LitUp(0, 0)
	grid1.LitUp(0, 1)
	grid1.LitUp(1, 0)
	grid1.LitUp(1, 2)
	grid2 := NewGrid(3, 3)
	grid2.LitUp(0, 0)
	grid2.LitUp(0, 1)
	grid2.LitUp(1, 1)
	grid2.LitUp(1, 2)
	tests := []struct {
		input    Grid
		row      int
		shift    int
		expected Grid
	}{
		{input: NewGrid(3, 3), row: 1, shift: 2, expected: grid1},
		{input: NewGrid(3, 3), row: 1, shift: 5, expected: grid1},

		{input: NewGrid(3, 3), row: 1, shift: 1, expected: grid2},
		{input: NewGrid(3, 3), row: 1, shift: 4, expected: grid2},
	}

	for _, test := range tests {
		result := copyGrid(base)
		result.RotateRow(test.row, test.shift)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Grid.RotateRow(row=%d, shift=%d) = %v, wants %v", test.row, test.shift, result.String(), test.expected.String())
		}
	}
}

func TestRotateCol(t *testing.T) {
	base := NewGrid(3, 3)
	base.LitUp(0, 0)
	base.LitUp(0, 1)
	base.LitUp(1, 0)
	base.LitUp(1, 1)
	grid1 := NewGrid(3, 3)
	grid1.LitUp(0, 0)
	grid1.LitUp(0, 1)
	grid1.LitUp(1, 0)
	grid1.LitUp(2, 1)
	grid2 := NewGrid(3, 3)
	grid2.LitUp(0, 0)
	grid2.LitUp(1, 1)
	grid2.LitUp(1, 0)
	grid2.LitUp(2, 1)
	tests := []struct {
		input    Grid
		col      int
		shift    int
		expected Grid
	}{
		{input: NewGrid(3, 3), col: 1, shift: 2, expected: grid1},
		{input: NewGrid(3, 3), col: 1, shift: 5, expected: grid1},

		{input: NewGrid(3, 3), col: 1, shift: 1, expected: grid2},
		{input: NewGrid(3, 3), col: 1, shift: 4, expected: grid2},
	}

	for _, test := range tests {
		result := copyGrid(base)
		result.RotateCol(test.col, test.shift)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Grid.RotateCol(col=%d, shift=%d) = %v, wants %v", test.col, test.shift, result.String(), test.expected.String())
		}
	}
}

func TestSolvePart1(t *testing.T) {
	rows, cols := 3, 7
	content1 := `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1`
	tests := []struct {
		content  string
		expected int
	}{
		{content: content1, expected: 6},
	}

	for _, test := range tests {
		is, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%s) failed prematurely", test.content)
			continue
		}
		result := solvePart1(is, rows, cols)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("solvePart1(%v) = %v, wants %v", test.content, result, test.expected)
		}
	}
}
