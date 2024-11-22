package day20

import (
	"cmp"
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

type Range struct {
	Begin int
	End   int
}

func NewRange(begin, end int) Range {
	return Range{Begin: begin, End: end}
}

func ParseRange(text string) (Range, error) {
	first, second, found := strings.Cut(text, "-")
	if !found {
		return NewRange(-1, -1), fmt.Errorf("missing range separator")
	}
	begin, err := strconv.Atoi(first)
	if err != nil {
		return NewRange(-1, -1), err
	}
	end, err := strconv.Atoi(second)
	if err != nil {
		return NewRange(-1, -1), err
	}
	return NewRange(begin, end), nil
}

func parseContent(content string) ([]Range, error) {
	content = strings.TrimSpace(content)
	lines := strings.Split(content, "\n")
	result := make([]Range, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		r, err := ParseRange(line)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

func parseFile(filename string) ([]Range, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func mergeInterleavedRanges(ranges []Range) []Range {
	slices.SortFunc(ranges, func(r1, r2 Range) int {
		return cmp.Compare(r1.Begin, r2.Begin)
	})
	result := make([]Range, 0, len(ranges))
	result = append(result, ranges[0])
	for _, r := range ranges {
		if r.Begin <= result[len(result)-1].End+1 {
			result[len(result)-1].End = utils.Max(result[len(result)-1].End, r.End)
		} else {
			result = append(result, r)
		}
	}
	return result
}

func minimumValueAllowed(ranges []Range) int {
	if len(ranges) == 0 {
		return -1
	}
	merged := mergeInterleavedRanges(ranges)
	return merged[0].End + 1
}

func part1(ranges []Range) int {
	return minimumValueAllowed(ranges)
}

func allowedValuesCount(ranges []Range) int {
	if len(ranges) == 0 {
		return -1
	}
	merged := mergeInterleavedRanges(ranges)
	count := 0
	for i := 1; i < len(merged); i++ {
		count += merged[i].Begin - merged[i-1].End - 1
	}
	maximumValue := 4294967295
	if maximumValue > merged[len(merged)-1].End {
		count += maximumValue - merged[len(merged)-1].End - 1
	}
	return count
}

func part2(ranges []Range) int {
	return allowedValuesCount(ranges)
}

func Solve() {
	input, err := parseFile("./inputs/day-20.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 20")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
