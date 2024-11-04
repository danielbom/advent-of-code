package day08

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

func readAllFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type FillRect struct {
	Row int
	Col int
}
type RotateRow struct {
	Row   int
	Shift int
}
type RotateCol struct {
	Col   int
	Shift int
}
type Instruction struct {
	Union interface{}
}

func NewInstructionFillRect(row, col int) Instruction {
	return Instruction{Union: FillRect{Row: row, Col: col}}
}
func NewInstructionRotateRow(row, shift int) Instruction {
	return Instruction{Union: RotateRow{Row: row, Shift: shift}}
}
func NewInstructionRotateCol(col, shift int) Instruction {
	return Instruction{Union: RotateCol{Col: col, Shift: shift}}
}

func parseInstruction(text string) (Instruction, error) {
	var inst Instruction
	if rest, found := strings.CutPrefix(text, "rect "); found {
		aStr, bStr, _ := strings.Cut(rest, "x")
		a, err := strconv.Atoi(aStr)
		if err != nil {
			return inst, err
		}
		b, err := strconv.Atoi(bStr)
		if err != nil {
			return inst, err
		}
		return NewInstructionFillRect(b, a), nil
	}
	if rest, found := strings.CutPrefix(text, "rotate row y="); found {
		aStr, bStr, _ := strings.Cut(rest, " by ")
		a, err := strconv.Atoi(aStr)
		if err != nil {
			return inst, err
		}
		b, err := strconv.Atoi(bStr)
		if err != nil {
			return inst, err
		}
		return NewInstructionRotateRow(a, b), nil
	}
	if rest, found := strings.CutPrefix(text, "rotate column x="); found {
		aStr, bStr, _ := strings.Cut(rest, " by ")
		a, err := strconv.Atoi(aStr)
		if err != nil {
			return inst, err
		}
		b, err := strconv.Atoi(bStr)
		if err != nil {
			return inst, err
		}
		return NewInstructionRotateCol(a, b), nil
	}
	return inst, fmt.Errorf("invalid instruction")
}

func parseContent(content string) ([]Instruction, error) {
	lines := strings.Split(content, "\n")
	result := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		result = append(result, instruction)
	}
	return result, nil
}

func parseFile(filename string) ([]Instruction, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

var letters = `
  A
..*..
.*.*.
*...*
*****
*...*
*...*
  B
***..
*..*.
***..
*..*.
*..*.
***..
  C
.**..
*..*.
*....
*....
*..*.
.**..
  D
**...
*.*..
*..*.
*..*.
*.*..
**...
  E
****.
*....
***..
*....
*....
****.
  F
****.
*....
***..
*....
*....
*....
  G
.**..
*..*.
*....
*.**.
*..*.
.**..
  H
*..*.
*..*.
****.
*..*.
*..*.
*..*.
  I
.***.
..*..
..*..
..*..
..*..
.***.
  K
*..*.
*.*..
**...
*.*..
*.*..
*..*.
  J
...*.
...*.
...*.
*..*.
*..*.
.**..
  L
*....
*....
*....
*....
*....
****.
  M
*...*
**.**
*.*.*
*...*
*...*
*...*
  N
*...*
**..*
*.*.*
*.*.*
*.,**
*...*
  O
.**..
*..*.
*..*.
*..*.
*..*.
.**..
  P
***..
*..*.
***..
*....
*....
*....
  Q
.**..
*..*.
*..*.
*..*.
*..*.
.**.*
  R
**...
*.*..
*.*..
**...
*.*..
*..*.
  S
.***.
*....
*....
.**..
...*.
***..
  T
*****
..*..
..*..
..*..
..*..
..*..
  U
*..*.
*..*.
*..*.
*..*.
*..*.
.**..
  V
*...*
*...*
*...*
*...*
.*.*.
..*..
  W
*...*
*...*
*...*
*.*.*
**.**
*...*
  X
*...*
.*.*.
..*..
..*..
.*.*.
*...*
  Y
*...*
*...*
.*.*.
..*..
..*..
..*..
  Z
*****
...*.
..*..
.*...
*....
*****
`

type Grid struct {
	grid []bool
	rows int
	cols int
}

func NewGrid(rows, cols int) Grid {
	grid := make([]bool, rows*cols)
	return Grid{rows: rows, cols: cols, grid: grid}
}

func CompileLetters() []Grid {
	result := make([]Grid, 0, 26)
	letter := NewGrid(6, 5)
	letterRow := 0
	for _, line := range strings.Split(letters, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if letterRow == 0 {
			letterRow++
			continue
		}
		for x, b := range line {
			if b == '*' {
				letter.LitUp(letterRow-1, x)
			}
		}
		letterRow++
		if letterRow-1 == letter.rows {
			result = append(result, letter)
			letter = NewGrid(6, 5)
			letterRow = 0
		}
	}
	if len(result) != cap(result) {
		panic("CompileLetters() failed")
	}
	return result
}

func (g *Grid) String() string {
	var b strings.Builder
	i := 0
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			if g.grid[i] {
				b.WriteByte('*')
			} else {
				b.WriteByte(' ')
			}
			i++
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (g *Grid) Display() {
	i := 0
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			if g.grid[i] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			i++
		}
		fmt.Println()
	}
}

func (g *Grid) CountLitUp() int {
	count := 0
	for _, item := range g.grid {
		if item {
			count++
		}
	}
	return count
}

func (g *Grid) Get(y, x int) bool {
	return g.grid[y*g.cols+x]
}

func (g *Grid) LitUp(y, x int) {
	g.grid[y*g.cols+x] = true
}

func (g *Grid) FillRect(row, col int) {
	row = utils.Min(row, g.rows)
	col = utils.Min(col, g.cols)
	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
			g.LitUp(y, x)
		}
	}
}

func (g *Grid) RotateRow(row, shift int) {
	y := utils.Min(row, g.rows-1)
	shift = shift % g.cols
	for i := 0; i < shift; i++ {
		last := g.grid[(y+1)*g.cols-1]
		for x := g.cols - 2; x >= 0; x-- {
			g.grid[y*g.cols+x+1] = g.grid[y*g.cols+x]
		}
		g.grid[y*g.cols+0] = last
	}
}

func (g *Grid) RotateCol(col, shift int) {
	x := utils.Min(col, g.cols-1)
	shift = shift % g.rows
	for i := 0; i < shift; i++ {
		last := g.grid[(g.rows-1)*g.cols+x]
		for y := g.rows - 2; y >= 0; y-- {
			g.grid[(y+1)*g.cols+x] = g.grid[y*g.cols+x]
		}
		g.grid[0*g.cols+x] = last
	}
}

func (g *Grid) Apply(instructions []Instruction) {
	for _, i := range instructions {
		switch v := i.Union.(type) {
		case FillRect:
			g.FillRect(v.Row, v.Col)
		case RotateRow:
			g.RotateRow(v.Row, v.Shift)
		case RotateCol:
			g.RotateCol(v.Col, v.Shift)
		}
	}
}

func (g *Grid) Text() string {
	letters := CompileLetters()
	stepy := letters[0].rows
	stepx := letters[0].cols
	var b strings.Builder
	for gy := 0; gy <= g.rows-stepy; gy += stepy {
		for gx := 0; gx <= g.cols-stepx; gx += stepx {
			foundIx := -1
			for ix, letter := range letters {
				equals := true
				for ly := 0; ly < letter.rows; ly++ {
					for lx := 0; lx < letter.cols; lx++ {
						if equals = letter.Get(ly, lx) == g.Get(gy+ly, gx+lx); !equals {
							break
						}
					}
					if !equals {
						break
					}
				}
				if equals {
					foundIx = ix
					break
				}
			}
			if foundIx == -1 {
				b.WriteByte(' ')
			} else {
				b.WriteByte(byte(int('A') + foundIx))
			}
		}
	}
	return b.String()
}

const (
	ROWS = 6
	COLS = 50
)

func solvePart1(instructions []Instruction, rows, cols int) int {
	grid := NewGrid(rows, cols)
	grid.Apply(instructions)
	return grid.CountLitUp()
}

func part1(instructions []Instruction) int {
	return solvePart1(instructions, ROWS, COLS)
}

func part2(instructions []Instruction) string {
	grid := NewGrid(ROWS, COLS)
	grid.Apply(instructions)
	//grid.Display()
	return grid.Text()
}

func Solve() {
	input, err := parseFile("./inputs/day-08.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 08")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%s", func() any { return part2(input) })
}
