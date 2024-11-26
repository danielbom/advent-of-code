package day06

import (
	"fmt"
	"strings"

	"aoc2016/internal/utils"
)

func parseContent(content string) ([]string, error) {
	result := strings.Split(content, "\n")
	// remove empty strings
	for i, s := range result {
		result[i] = strings.TrimSpace(s)
	}
	for len(result) > 0 && len(result[len(result)-1]) == 0 {
		result = result[:len(result)-1]
	}
	return result, nil
}

func parseFile(filename string) ([]string, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func computeFrequency(inputs []string, ix int) []int {
	frequency := make([]int, 256) // ASCII size
	for _, input := range inputs {
		k := int(input[ix])
		frequency[k]++
	}
	return frequency
}

func mostCommonByteAt(inputs []string, ix int) byte {
	frequency := computeFrequency(inputs, ix)
	maxIx := 0
	for ix, f := range frequency {
		maxF := frequency[maxIx]
		if f > maxF {
			maxIx = ix
		}
	}
	return byte(maxIx)
}

func part1(input []string) string {
	result := make([]byte, 0, len(input[0]))
	for i := 0; i < cap(result); i++ {
		ch := mostCommonByteAt(input, i)
		result = append(result, ch)
	}
	return string(result)
}

func leastCommonByteAt(inputs []string, ix int) byte {
	frequency := computeFrequency(inputs, ix)
	minIx := 0
	for ix, f := range frequency {
		minF := frequency[minIx]
		if minF == 0 {
			minIx = ix
			continue
		}
		if minF > f && f > 0 {
			minIx = ix
		}
	}
	return byte(minIx)
}

func part2(input []string) string {
	result := make([]byte, 0, len(input[0]))
	for i := 0; i < cap(result); i++ {
		ch := leastCommonByteAt(input, i)
		result = append(result, ch)
	}
	return string(result)
}

func Solve() {
	input, err := parseFile("./inputs/day-06.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 06")
	utils.TimeIt("Part 1:", "%s", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%s", func() any { return part2(input) })
}
