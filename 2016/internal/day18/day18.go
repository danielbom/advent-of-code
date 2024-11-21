package day18

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"aoc2016/internal/utils"
)

type Row []byte

const SAFE byte = '.'
const TRAP byte = '^'

func NewRow(text string) (Row, error) {
	row := make([]byte, len(text)+2)
	row[0] = SAFE
	row[len(row)-1] = SAFE
	for i, ch := range []byte(text) {
		if ch == SAFE || ch == TRAP {
			row[i+1] = ch
			continue
		}
		return Row(nil), fmt.Errorf("invalid row format")
	}
	return Row(row), nil
}

var TRAP_PATTERN1 = []byte{TRAP, TRAP, SAFE}
var TRAP_PATTERN2 = []byte{SAFE, TRAP, TRAP}
var TRAP_PATTERN3 = []byte{TRAP, SAFE, SAFE}
var TRAP_PATTERN4 = []byte{SAFE, SAFE, TRAP}

func (row Row) isTrap(i int) bool {
	return (row[i-1] == TRAP && row[i+1] == SAFE) ||
		(row[i-1] == SAFE && row[i+1] == TRAP)
	tiles := row[i-1 : i+2]
	return slices.Compare(tiles, TRAP_PATTERN1) == 0 ||
		slices.Compare(tiles, TRAP_PATTERN2) == 0 ||
		slices.Compare(tiles, TRAP_PATTERN3) == 0 ||
		slices.Compare(tiles, TRAP_PATTERN4) == 0
}

func (row Row) NextRow() Row {
	next := make([]byte, len(row))
	next[0] = SAFE
	next[len(next)-1] = SAFE
	for i := 1; i < len(row)-1; i++ {
		if row.isTrap(i) {
			next[i] = TRAP
		} else {
			next[i] = SAFE
		}
	}
	return Row(next)
}

func (row Row) SafesCount() int {
	count := 0
	for _, tile := range row[1 : len(row)-1] {
		if tile == SAFE {
			count++
		}
	}
	return count
}

func (row Row) String() string {
	return string(row[1 : len(row)-1])
}

func readAllFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseFile(filename string) (Row, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return Row(nil), err
	}
	content = strings.TrimSpace(content)
	return NewRow(content)
}

func countSafeTiles(initial Row, rows int) int {
	current := initial
	count := 0
	for i := 0; i < rows; i++ {
		//fmt.Println(current.String())
		count += current.SafesCount()
		current = current.NextRow()
	}
	return count
}

func part1(initial Row) int {
	return countSafeTiles(initial, 40)
}

func part2(initial Row) int {
	return countSafeTiles(initial, 400000)
}

func Solve() {
	input, err := parseFile("./inputs/day-18.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 18")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
