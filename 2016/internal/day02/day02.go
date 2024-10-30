package day02

import (
	"fmt"
	"os"
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

type Instruction int

const (
	UP Instruction = iota
	DOWN
	LEFT
	RIGHT
)

func parseContent(content string) ([][]Instruction, error) {
	lines := strings.Split(content, "\n")
	iss := make([][]Instruction, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		is := make([]Instruction, 0, len(line))
		for _, ch := range line {
			switch ch {
			case 'U':
				is = append(is, UP)
			case 'L':
				is = append(is, LEFT)
			case 'R':
				is = append(is, RIGHT)
			case 'D':
				is = append(is, DOWN)
			}
		}
		iss = append(iss, is)
	}
	return iss, nil
}

func parseFile(filename string) ([][]Instruction, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
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

func composeKeys(keys []int) int64 {
	power := int64(1)
	result := int64(0)
	for i := len(keys) - 1; i >= 0; i-- {
		result += int64(keys[i]) * power
		power *= int64(10)
	}
	return result
}

func computeCode(keypad []string, y, x int, iss [][]Instruction) string {
	p := NewPoint(x, y) // 5
	keys := make([]byte, 0, len(iss))
	for _, is := range iss {
		for _, i := range is {
			switch i {
			case LEFT:
				if p.X > 0 {
					p.X -= 1
					if keypad[p.Y][p.X] == ' ' {
						p.X += 1
					}
				}
			case RIGHT:
				if p.X < len(keypad[p.Y])-1 {
					p.X += 1
					if keypad[p.Y][p.X] == ' ' {
						p.X -= 1
					}
				}
			case UP:
				if p.Y > 0 && p.X < len(keypad[p.Y-1]) {
					p.Y -= 1
					if keypad[p.Y][p.X] == ' ' {
						p.Y += 1
					}
				}
			case DOWN:
				if p.Y < len(keypad)-1 && p.X < len(keypad[p.Y+1]) {
					p.Y += 1
					if keypad[p.Y][p.X] == ' ' {
						p.Y -= 1
					}
				}
			}
		}
		key := keypad[p.Y][p.X]
		keys = append(keys, key)
	}
	return string(keys)
}

func part1(iss [][]Instruction) string {
	keypad := []string{
		"123",
		"456",
		"789",
	}
	return computeCode(keypad, 1, 1, iss)
}

func part2(iss [][]Instruction) string {
	keypad := []string{
		"  1  ",
		" 234 ",
		"56789",
		" ABC ",
		"  D  ",
	}
	return computeCode(keypad, 2, 0, iss)
}

func Solve() {
	input, err := parseFile("./inputs/day-02.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 02")
	utils.TimeIt("Part 1:", "%s", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%s", func() any { return part2(input) })
}
