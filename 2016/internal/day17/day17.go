package day17

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"aoc2016/internal/utils"
)

func MD5Hash(passcode string, path string) string {
	hasher := md5.New()
	io.WriteString(hasher, passcode)
	io.WriteString(hasher, path)
	return hex.EncodeToString(hasher.Sum(nil))
}

func readAllFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseFile(filename string) (string, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return "", err
	}
	content = strings.TrimSpace(content)
	return content, nil
}

func isOpen(ch byte) bool {
	return int('b') <= int(ch) && int(ch) <= int('f')
}

type PathFinder interface {
	Decide(best, path string) string
	Break(best, path string) bool
}

type ShortestPathFinder struct{}

func (pf ShortestPathFinder) Decide(best, path string) string {
	if len(best) > 0 && len(best) < len(path) {
		return best
	} else {
		return path
	}
}

func (pf ShortestPathFinder) Break(best, path string) bool {
	return len(best) > 0 && len(path) >= len(best)
}

type LongestPathFinder struct{}

func (pf LongestPathFinder) Decide(best, path string) string {
	if len(best) > len(path) {
		return best
	} else {
		return path
	}
}

func (pf LongestPathFinder) Break(best, path string) bool {
	return false
}

func findPathRec(pf PathFinder, x, y int, passcode, path, best string) []string {
	if !(0 <= x && x <= 3) || !(0 <= y && y <= 3) {
		return nil
	}
	if x == 3 && y == 3 {
		return []string{pf.Decide(best, path)}
	}
	if pf.Break(best, path) {
		return nil
	}
	hash := MD5Hash(passcode, path)
	if isOpen(hash[1]) /* down */ {
		result := findPathRec(pf, x, y+1, passcode, path+"D", best)
		if len(result) > 0 {
			best = result[0]
		}
	}
	if isOpen(hash[3]) /* right */ {
		result := findPathRec(pf, x+1, y, passcode, path+"R", best)
		if len(result) > 0 {
			best = result[0]
		}
	}
	if isOpen(hash[2]) /* left */ {
		result := findPathRec(pf, x-1, y, passcode, path+"L", best)
		if len(result) > 0 {
			best = result[0]
		}
	}
	if isOpen(hash[0]) /* up */ {
		result := findPathRec(pf, x, y-1, passcode, path+"U", best)
		if len(result) > 0 {
			best = result[0]
		}
	}
	if len(best) != 0 {
		return []string{best}
	}
	return nil
}

func findPath(pf PathFinder, passcode string) string {
	result := findPathRec(pf, 0, 0, passcode, "", "")
	if len(result) == 1 {
		return result[0]
	}
	return ""
}

func part1(passcode string) string {
	var shortest ShortestPathFinder
	return findPath(shortest, passcode)
}

func part2(passcode string) int {
	var longest LongestPathFinder
	path := findPath(longest, passcode)
	return len(path)
}

func Solve() {
	input, err := parseFile("./inputs/day-17.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 17")
	utils.TimeIt("Part 1:", "%s", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
