package day07

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

type IPv7 string

// IP supports TLS (transport-layer snooping)
func (ip *IPv7) SupportsTSL() bool {
	abbaFound := false
	bracedFound := false
	braced := false
	for bytes := []byte(*ip); len(bytes) >= 4; bytes = bytes[1:] {
		if bytes[0] == '[' {
			braced = true
			continue
		} else if bytes[0] == ']' {
			braced = false
			continue
		}
		if bytes[0] != bytes[1] && bytes[0] == bytes[3] && bytes[1] == bytes[2] {
			if braced {
				bracedFound = true
			} else {
				abbaFound = true
			}
		}
	}
	return !bracedFound && abbaFound
}

// IP supports SLL (super-secret listening)
func (ip *IPv7) SupportsSSL() bool {
	braced := false
	babFound := make([]string, 0, 4)
	abaFound := make([]string, 0, 4)
	for bytes := []byte(*ip); len(bytes) >= 3; bytes = bytes[1:] {
		if bytes[0] == '[' {
			braced = true
			continue
		} else if bytes[0] == ']' {
			braced = false
			continue
		}
		if bytes[1] == '[' || bytes[1] == ']' {
			continue
		}
		if bytes[0] != bytes[1] && bytes[0] == bytes[2] {
			if braced {
				babFound = append(babFound, string(bytes[:3]))
			} else {
				abaFound = append(abaFound, string(bytes[:3]))
			}
		}
	}
	for _, aba := range abaFound {
		for _, bab := range babFound {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}
	return false
}

func parseContent(content string) ([]IPv7, error) {
	lines := strings.Split(content, "\n")
	result := make([]IPv7, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			result = append(result, IPv7(line))
		}
	}
	return result, nil
}

func parseFile(filename string) ([]IPv7, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func part1(ips []IPv7) int {
	count := 0
	for _, ip := range ips {
		if ip.SupportsTSL() {
			count++
		}
	}
	return count
}

func part2(ips []IPv7) int {
	count := 0
	for _, ip := range ips {
		if ip.SupportsSSL() {
			count++
		}
	}
	return count
}

func Solve() {
	input, err := parseFile("./inputs/day-07.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 07")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
