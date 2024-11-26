package day03

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

type Triangle struct {
	A int
	B int
	C int
}

func NewTriangle(a, b, c int) Triangle {
	var t Triangle
	t.A = a
	t.B = b
	t.C = c
	return t
}

func TriangleIsValid(a, b, c int) bool {
	if b > c {
		return TriangleIsValid(a, c, b)
	}
	if a > c {
		return TriangleIsValid(c, b, a)
	}
	if a > b {
		return TriangleIsValid(b, a, c)
	}
	return a+b > c
}

func (t *Triangle) IsValid() bool {
	return TriangleIsValid(t.A, t.B, t.C)
}

func parseContent(content string) ([]Triangle, error) {
	lines := strings.Split(content, "\n")
	result := make([]Triangle, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		fields := strings.Fields(line)
		if len(fields) != 3 {
			return nil, fmt.Errorf("invalid line in content: '%s'", line)
		}
		a, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		c, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, err
		}
		result = append(result, NewTriangle(a, b, c))
	}
	return result, nil
}

func parseFile(filename string) ([]Triangle, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func part1(triangles []Triangle) int {
	valid := 0
	for _, t := range triangles {
		if t.IsValid() {
			valid++
		}
	}
	return valid
}

func part2(triangles []Triangle) int {
	valid := 0
	for i := 0; i < len(triangles)-2; i += 3 {
		ta := NewTriangle(triangles[i+0].A, triangles[i+1].A, triangles[i+2].A)
		if ta.IsValid() {
			valid++
		}
		tb := NewTriangle(triangles[i+0].B, triangles[i+1].B, triangles[i+2].B)
		if tb.IsValid() {
			valid++
		}
		tc := NewTriangle(triangles[i+0].C, triangles[i+1].C, triangles[i+2].C)
		if tc.IsValid() {
			valid++
		}
	}
	return valid
}

func Solve() {
	input, err := parseFile("./inputs/day-03.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 03")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
