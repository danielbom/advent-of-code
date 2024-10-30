package day01

import (
	"fmt"
	"os"
	"slices"
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

type Turn int

const (
	RIGHT Turn = iota
	LEFT
)

type Instruction struct {
	Turn  Turn
	Count int
}

func parseInstruction(text string) (Instruction, error) {
	var i Instruction
	if countText, found := strings.CutPrefix(text, "R"); found {
		count, err := strconv.Atoi(countText)
		if err != nil {
			return i, err
		}
		i.Turn = RIGHT
		i.Count = count
		return i, nil
	}
	if countText, found := strings.CutPrefix(text, "L"); found {
		count, err := strconv.Atoi(countText)
		if err != nil {
			return i, err
		}
		i.Turn = LEFT
		i.Count = count
		return i, nil
	}
	return i, fmt.Errorf("invalid direction")
}

func parseContent(content string) ([]Instruction, error) {
	words := strings.Split(content, " ")

	result := make([]Instruction, 0, len(words))
	for _, word := range words {
		word, _ = strings.CutSuffix(word, ",")
		word = strings.TrimSpace(word)

		instruction, err := parseInstruction(word)
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

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (d *Direction) Turn(t Turn) {
	if t == LEFT {
		switch *d {
		case NORTH:
			*d = WEST
		case EAST:
			*d = NORTH
		case SOUTH:
			*d = EAST
		case WEST:
			*d = SOUTH
		}
	} else {
		switch *d {
		case NORTH:
			*d = EAST
		case EAST:
			*d = SOUTH
		case SOUTH:
			*d = WEST
		case WEST:
			*d = NORTH
		}
	}
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	var point Point
	point.X = x
	point.Y = y
	return point
}

func (p *Point) Move(dir Direction, count int) {
	switch dir {
	case NORTH:
		p.Y += count
	case EAST:
		p.X -= count
	case SOUTH:
		p.Y -= count
	case WEST:
		p.X += count
	}
}

func part1(instructions []Instruction) int {
	dir := NORTH
	p := NewPoint(0, 0)

	for _, i := range instructions {
		dir.Turn(i.Turn)
		p.Move(dir, i.Count)
	}

	return utils.Abs(p.X) + utils.Abs(p.Y)
}

func part2(instructions []Instruction) int {
	dir := NORTH
	curr := NewPoint(0, 0)
	seen := []Point{}

	for _, i := range instructions {
		begin := curr

		dir.Turn(i.Turn)
		curr.Move(dir, i.Count)

		found := false
		inc := 1
		if begin.X > curr.X || begin.Y > curr.Y {
			inc = -1
		}

		if begin.X != curr.X {
			begin.X += inc
			for begin.X != curr.X {
				if slices.Contains(seen, begin) {
					curr = begin
					found = true
					break
				}
				seen = append(seen, begin)
				begin.X += inc
			}
		}

		if begin.Y != curr.Y {
			begin.Y += inc
			for begin.Y != curr.Y {
				if slices.Contains(seen, begin) {
					curr = begin
					found = true
					break
				}
				seen = append(seen, begin)
				begin.Y += inc
			}
		}

		if found {
			break
		}
	}

	return utils.Abs(curr.X) + utils.Abs(curr.Y)
}

func Solve() {
	instructions, err := parseFile("./inputs/day-01.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 01")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(instructions) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(instructions) })
}
