package day15

import (
	"fmt"
	"os"
	"regexp"
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

type Disc struct {
	ID        int
	Time      int
	Position  int
	Positions int
}

type DiscParser struct {
	re *regexp.Regexp
}

func NewDiscParser() DiscParser {
	var p DiscParser
	p.re = regexp.MustCompile(`Disc #(\d+) has (\d+) positions; at time=(\d+), it is at position (\d+)\.`)
	return p
}

func (p *DiscParser) Parse(text string) (Disc, error) {
	var result Disc
	if !p.re.MatchString(text) {
		return result, fmt.Errorf("invalid disc format")
	}
	matches := p.re.FindStringSubmatch(text)
	result.ID, _ = strconv.Atoi(matches[1])
	result.Positions, _ = strconv.Atoi(matches[2])
	result.Time, _ = strconv.Atoi(matches[3])
	result.Position, _ = strconv.Atoi(matches[4])
	return result, nil
}

func (p *DiscParser) ParseLines(content string) ([]Disc, error) {
	lines := strings.Split(content, "\n")
	result := make([]Disc, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		disc, err := p.Parse(line)
		if err != nil {
			return nil, err
		}
		result = append(result, disc)
	}
	return result, nil
}

func parseContent(content string) ([]Disc, error) {
	p := NewDiscParser()
	return p.ParseLines(content)
}

func parseFile(filename string) ([]Disc, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func AdvanceTime(ds, rs []Disc, t int) {
	for i, d := range ds {
		rs[i].Time = d.Time + t + 1
		rs[i].Position = (d.Position + i + t + 1) % d.Positions
	}
}

func IsAligned(ds []Disc) bool {
	for _, d := range ds {
		if d.Position != 0 {
			return false
		}
	}
	return true
}

func findAlignedTime(ds []Disc) int {
	rs := make([]Disc, len(ds))
	copy(rs, ds)
	t := 0
	for ; !IsAligned(rs); t++ {
		AdvanceTime(ds, rs, t)
	}
	return t - 1
}

func part1(ds []Disc) int {
	return findAlignedTime(ds)
}

func part2(ds []Disc) int {
	extra := Disc{
		ID:        len(ds) + 1,
		Time:      0,
		Position:  0,
		Positions: 11,
	}
	ds = append(ds, extra)
	return findAlignedTime(ds)
}

func Solve() {
	input, err := parseFile("./inputs/day-15.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 15")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
