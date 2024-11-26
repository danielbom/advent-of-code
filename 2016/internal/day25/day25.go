package day25

import (
	"fmt"

	"aoc2016/internal/utils"
)

func parseFile(filename string) (any, error) {
	_, err := utils.ReadAllFile(filename)
	if err != nil {
		return 0, err
	}
	return 0, err
}

func part1(input any) int {
	return 0
}

func part2(input any) int {
	return 0
}

func Solve() {
	input, err := parseFile("./inputs/day-25.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 25")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
