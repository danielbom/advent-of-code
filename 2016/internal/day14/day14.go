package day14

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"

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
	content, err := readAllFile(filename)
	if err != nil {
		return "", err
	}
	content = strings.TrimSpace(content)
	return content, err
}

func MD5Hash1(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func MD5Hash(text string, stretch int) string {
	hash := MD5Hash1(text)
	for i := 0; i < stretch; i++ {
		hash = MD5Hash1(hash)
	}
	return hash
}

type HashPattern struct {
	Size3 int
	Size5 int
}

func CollectHashPattern(hash string) HashPattern {
	size3 := -1
	size5 := -1
	count := 1
	prev := '?'
	for _, ch := range hash {
		if ch == prev {
			count++
		} else {
			count = 1
		}
		if count == 3 && size3 == -1 {
			size3 = int(ch)
		}
		if count == 5 && size5 == -1 {
			size5 = int(ch)
		}
		if size3 > 0 && size5 > 0 {
			break
		}
		prev = ch
	}
	var p HashPattern
	p.Size3 = size3
	p.Size5 = size5
	return p
}

func NextHashPattern(salt string, count int, stretch int) HashPattern {
	input := fmt.Sprintf("%s%d", salt, count)
	hash := MD5Hash(input, stretch)
	return CollectHashPattern(hash)
}

type HashFinder struct {
	salt    string
	hs      []HashPattern
	count   int
	index   int
	stretch int
}

func NewHashFinder(salt string, stretch int) HashFinder {
	var f HashFinder
	f.salt = salt
	f.hs = make([]HashPattern, 0, 1001)
	f.count = 0
	f.index = 0
	f.stretch = stretch
	return f
}

func (f *HashFinder) Next() {
	h := NextHashPattern(f.salt, f.count, f.stretch)
	f.hs = append(f.hs, h)
	f.count++
}

func HashMatch(h HashPattern, hs []HashPattern) bool {
	if h.Size3 == -1 {
		return false
	}
	for _, o := range hs {
		if o.Size5 == h.Size3 {
			return true
		}
	}
	return false
}

func (f *HashFinder) ComputeParallel() {
	type Pattern struct {
		pattern HashPattern
		index   int
	}
	var wg sync.WaitGroup
	workers := runtime.NumCPU()
	n := workers*1000 + 1000
	hs := make([]HashPattern, n)
	patterns := make(chan Pattern, workers*2)
	index := make(chan int, workers)
	go func() {
		for p := range patterns {
			hs[p.index] = p.pattern
		}
	}()
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range index {
				p := NextHashPattern(f.salt, f.count+i, f.stretch)
				patterns <- Pattern{pattern: p, index: i}
			}
		}()
	}
	for i := 0; i < n; i++ {
		index <- i
	}
	close(index)
	wg.Wait()
	close(patterns)
	if len(f.hs) == 0 {
		f.hs = hs
	} else {
		f.hs = append(f.hs[len(f.hs)-1000:], hs...)
	}
	f.count += n
}

func (f *HashFinder) RunParallel() {
	matches := 0
	for matches < 64 {
		f.ComputeParallel()
		for i := 0; i < len(f.hs)-1000; i++ {
			f.index = f.count - len(f.hs) + i
			curr := f.hs[i]
			next := f.hs[i+1 : i+1000+1]
			if HashMatch(curr, next) {
				matches++
				if matches == 64 {
					break
				}
			}
		}
	}
}

func (f *HashFinder) Run() {
	for i := 0; i < 1000; i++ {
		f.Next()
	}
	matches := 0
	for matches < 64 {
		f.Next()
		curr := f.hs[f.index]
		next := f.hs[f.index+1 : f.index+1000+1]
		if HashMatch(curr, next) {
			matches++
		}
		f.index++
	}
	f.index--
}

const PARALLEL bool = true

func part1(salt string) int {
	f := NewHashFinder(salt, 0)
	if PARALLEL {
		f.RunParallel()
	} else {
		f.Run()
	}
	return f.index
}

func part2(salt string) int {
	f := NewHashFinder(salt, 2016)
	if PARALLEL {
		f.RunParallel()
	} else {
		f.Run()
	}
	return f.index
}

func Solve() {
	input, err := parseFile("./inputs/day-14.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 14")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
