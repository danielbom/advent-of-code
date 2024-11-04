package day09

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"aoc2016/internal/utils"
)

func readAllFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseFile(filename string) (string, error) {
	return readAllFile(filename)
}

// Decompressor
type D struct {
	re *regexp.Regexp
}

func NewD() D {
	var d D
	d.re = regexp.MustCompile(`^\(\d+x\d+\)`)
	return d
}

type DMatch struct {
	End    int
	Length int
	Repeat int
	Found  bool
}

func (m *DMatch) Skip(input string) string {
	return input[m.End+m.Length+1:]
}

func (m *DMatch) Collect(input string) string {
	return input[m.End+1 : m.End+1+m.Length]
}

func (d *D) Match(input string) (m DMatch) {
	if m.Found = d.re.MatchString(input); m.Found {
		m.End = strings.IndexByte(input, ')')
		cmd := input[1:m.End]
		start, end, _ := strings.Cut(cmd, "x")
		m.Length, _ = strconv.Atoi(start)
		m.Repeat, _ = strconv.Atoi(end)
	}
	return m
}

func (d *D) Decompress(input string) string {
	var b strings.Builder
	for len(input) > 0 {
		if unicode.IsSpace(rune(input[0])) {
			input = input[1:]
			continue
		}
		if m := d.Match(input); m.Found {
			for i := 0; i < m.Repeat; i++ {
				b.WriteString(m.Collect(input))
			}
			input = m.Skip(input)
		} else {
			b.WriteByte(input[0])
			input = input[1:] // tail
		}
	}
	return b.String()
}

func (d *D) DecompressLength(input string) int {
	lengthSum := 0
	for len(input) > 0 {
		if unicode.IsSpace(rune(input[0])) {
			input = input[1:]
			continue
		}
		if m := d.Match(input); m.Found {
			lengthSum += m.Repeat * m.Length
			input = m.Skip(input)
		} else {
			lengthSum++
			input = input[1:] // tail
		}
	}
	return lengthSum
}

func (d *D) DeepDecompressLength(input string) int {
	lengthSum := 0
	for len(input) > 0 {
		if unicode.IsSpace(rune(input[0])) {
			input = input[1:]
			continue
		}
		if m := d.Match(input); m.Found {
			deepLength := d.DeepDecompressLength(m.Collect(input))
			lengthSum += m.Repeat * deepLength
			input = m.Skip(input)
		} else {
			lengthSum++
			input = input[1:] // tail
		}
	}
	return lengthSum
}

func part1(input string) int {
	d := NewD()
	return d.DecompressLength(input)
}

func part2(input string) int {
	d := NewD()
	return d.DeepDecompressLength(input)
}

func Solve() {
	input, err := parseFile("./inputs/day-09.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 09")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
