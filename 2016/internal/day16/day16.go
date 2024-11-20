package day16

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

func parseFile(filename string) (Data, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return Data(""), err
	}
	return NewData(content)
}

type Data string

func NewData(content string) (Data, error) {
	content = strings.TrimSpace(content)
	for _, ch := range content {
		if ch != '0' && ch != '1' {
			return Data(""), fmt.Errorf("invalid data")
		}
	}
	return Data(content), nil
}

func (data Data) Increase() Data {
	// modified dragon curve: https://en.wikipedia.org/wiki/Dragon_curve
	var sb strings.Builder
	sb.Grow(len(data)*2 + 1)
	sb.WriteString(string(data))
	sb.WriteByte('0')
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == '1' {
			sb.WriteByte('0')
		} else {
			sb.WriteByte('1')
		}
	}
	return Data(sb.String())
}

func (data Data) Take(size int) Data {
	return Data(data[:size])
}

func (data Data) checksum1() Data {
	var sb strings.Builder
	sb.Grow(len(data) / 2)
	for i := 0; i+1 < len(data); i += 2 {
		if data[i] == data[i+1] {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return Data(sb.String())
}

func (data Data) Checksum() Data {
	result := data.checksum1()
	for len(result)%2 == 0 {
		result = result.checksum1()
	}
	return result
}

func (data Data) IncreaseUntilSize(size int) Data {
	current := data
	for len(current) < size {
		current = current.Increase()
	}
	return current
}

func (data Data) ChecksumWithSize(size int) Data {
	return data.IncreaseUntilSize(size).Take(size).Checksum()
}

func part1(data Data) string {
	return string(data.ChecksumWithSize(272))
}

func part2(data Data) string {
	return string(data.ChecksumWithSize(35651584))
}

func Solve() {
	input, err := parseFile("./inputs/day-16.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 16")
	utils.TimeIt("Part 1:", "%s", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%s", func() any { return part2(input) })
}
