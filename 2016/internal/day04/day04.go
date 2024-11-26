package day04

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

type EncryptedName struct {
	Name     string
	SectorID int
	Checksum string
}

func NewEncryptedName(name string, sectorID int, checksum string) EncryptedName {
	var e EncryptedName
	e.Name = name
	e.SectorID = sectorID
	e.Checksum = checksum
	return e
}

type ByteFreq struct {
	Byte rune
	Freq int
}

func (e *EncryptedName) IsReal() bool {
	frequency := make([]ByteFreq, 0, 256) // ASCII size
	for i := 0; i < cap(frequency); i++ {
		f := ByteFreq{Byte: rune(i), Freq: 0}
		frequency = append(frequency, f)
	}
	for _, c := range e.Name {
		if c != '-' {
			k := int(c)
			frequency[k].Freq++
		}
	}
	slices.SortFunc(frequency, func(a, b ByteFreq) int {
		cmpFreq := cmp.Compare(b.Freq, a.Freq) // reversed
		if cmpFreq != 0 {
			return cmpFreq
		}
		return cmp.Compare(a.Byte, b.Byte)
	})
	for i, c := range e.Checksum {
		if frequency[i].Byte != c {
			return false
		}
	}
	return true
}

// Example: "aaaaa-bbb-z-y-x-123[abxyz]"
func parseEncryptedName(text string) (EncryptedName, error) {
	var e EncryptedName
	ix_dash := strings.LastIndexByte(text, '-')
	if ix_dash == -1 {
		return e, fmt.Errorf("dash not found")
	}
	ix_ob := strings.LastIndexByte(text, '[')
	if ix_ob == -1 {
		return e, fmt.Errorf("open brace not found")
	}
	if !strings.HasSuffix(text, "]") {
		return e, fmt.Errorf("close brace not found")
	}
	sectorIDStr := text[ix_dash+1 : ix_ob]
	sectorID, err := strconv.Atoi(sectorIDStr)
	if err != nil {
		return e, fmt.Errorf("invalid sector id: %v", err)
	}
	e.Name = text[0:ix_dash]
	e.SectorID = sectorID
	e.Checksum = text[ix_ob+1 : len(text)-1]
	return e, nil
}

func parseContent(content string) ([]EncryptedName, error) {
	lines := strings.Split(content, "\n")
	result := make([]EncryptedName, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		name, err := parseEncryptedName(line)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, line)
		}
		result = append(result, name)
	}
	return result, nil
}

func parseFile(filename string) ([]EncryptedName, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func part1(names []EncryptedName) int {
	sum := 0
	for _, n := range names {
		if n.IsReal() {
			sum += n.SectorID
		}
	}
	return sum
}

func ShiftCipher(input string, shift int) string {
	shift = shift % (int('z') - int('a') + 1)
	bytes := make([]rune, 0, len(input))
	for _, c := range input {
		if c == '-' {
			bytes = append(bytes, ' ')
		} else if int('a') <= int(c) && int(c) <= int('z') {
			k := (int(c) + shift - int('a')) % (int('z') - int('a') + 1)
			bytes = append(bytes, rune(k+int('a')))
		} else if int('A') <= int(c) && int(c) <= int('Z') {
			k := (int(c) + shift - int('A')) % (int('Z') - int('A') + 1)
			bytes = append(bytes, rune(k+int('A')))
		} else {
			bytes = append(bytes, c)
		}
	}
	return string(bytes)
}

func part2(names []EncryptedName) int {
	for _, n := range names {
		if n.IsReal() {
			s := ShiftCipher(n.Name, n.SectorID)
			if s == "northpole object storage" {
				return n.SectorID
			}
		}
	}
	return 0
}

func Solve() {
	input, err := parseFile("./inputs/day-04.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 04")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
