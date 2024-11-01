package day05

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

func MD5Hash(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

type Gen interface {
	Complete() bool
	Consume(hash string)
	Result() string
}

type Gen1 struct {
	result []byte
}

func NewGen1() Gen1 {
	var g Gen1
	g.result = make([]byte, 0, 8)
	return g
}

func (g *Gen1) Complete() bool {
	return len(g.result) >= 8
}

func (g *Gen1) Consume(hash string) {
	if strings.HasPrefix(hash, "00000") {
		g.result = append(g.result, hash[5])
	}
}

func (g *Gen1) Result() string {
	return string(g.result)
}

type BytePos struct {
	Byte rune
	Pos  int
}

type Gen2 struct {
	pos []BytePos
}

func NewGen2() Gen2 {
	var g Gen2
	g.pos = make([]BytePos, 8)
	return g
}

func (g *Gen2) Complete() bool {
	for _, c := range g.pos {
		if c.Pos == 0 {
			return false
		}
	}
	return true
}

func (g *Gen2) Consume(hash string) {
	if strings.HasPrefix(hash, "00000") {
		i := int(hash[5])
		if int('0') <= i && i <= int('7') {
			i = i - int('0')
			if g.pos[i].Pos == 0 {
				g.pos[i].Pos = i + 1
				g.pos[i].Byte = rune(hash[6])
			}
		}
	}
}

func (g *Gen2) Result() string {
	result := make([]rune, len(g.pos))
	for _, p := range g.pos {
		result[p.Pos-1] = p.Byte
	}
	return string(result)
}

func GeneratePasswordNaive(g Gen, prefix string) string {
	count := 0
	for !g.Complete() {
		input := fmt.Sprintf("%s%d", prefix, count)
		count++
		hash := MD5Hash(input)
		g.Consume(hash)
	}
	return g.Result()
}

func GeneratePasswordFancy(g Gen, prefix string) string {
	var wg sync.WaitGroup

	workers := runtime.NumCPU()
	if workers > 1 {
		workers--
	}
	if workers > 1 {
		workers--
	}

	more := make(chan bool, workers+2)
	in := make(chan int, workers+2)
	out := make(chan byte, 8)
	end := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 1
		for {
			moreCount := <-more
			if moreCount {
				in <- count
				count++
			} else {
				break
			}
		}
		for i := 0; i < workers; i++ {
			in <- 0
		}
	}()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			times := 64 // making the computation more expensive (worthy)
			more <- true
			for {
				count := <-in
				if count == 0 {
					break
				}
				for i := 0; i < times; i++ {
					input := fmt.Sprintf("%s%d", prefix, (count*times)+i)
					hash := MD5Hash(input)
					g.Consume(hash)
				}
				more <- !g.Complete()
			}
		}()
	}

	wg.Wait()
	close(end)
	close(in)
	close(out)
	close(more)

	return g.Result()
}

const NAIVE = false

func part1(prefix string) string {
	g := NewGen1()
	if NAIVE {
		return GeneratePasswordNaive(&g, prefix)
	} else {
		return GeneratePasswordFancy(&g, prefix)
	}
}

func part2(prefix string) string {
	g := NewGen2()
	if NAIVE {
		return GeneratePasswordNaive(&g, prefix)
	} else {
		return GeneratePasswordFancy(&g, prefix)
	}
}

func Solve() {
	input, err := parseFile("./inputs/day-05.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 05")
	utils.TimeIt("Part 1:", "%s", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%s", func() any { return part2(input) })
}
