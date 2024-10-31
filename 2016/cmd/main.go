package main

import (
	"fmt"
	"os"
	"strconv"

	"aoc2016/internal/day01"
	"aoc2016/internal/day02"
	"aoc2016/internal/day03"
	"aoc2016/internal/day04"
	"aoc2016/internal/day05"
	"aoc2016/internal/day06"
	"aoc2016/internal/day07"
	"aoc2016/internal/day08"
	"aoc2016/internal/day09"
	"aoc2016/internal/day10"
	"aoc2016/internal/day11"
	"aoc2016/internal/day12"
	"aoc2016/internal/day13"
	"aoc2016/internal/day14"
	"aoc2016/internal/day15"
	"aoc2016/internal/day16"
	"aoc2016/internal/day17"
	"aoc2016/internal/day18"
	"aoc2016/internal/day19"
	"aoc2016/internal/day20"
	"aoc2016/internal/day21"
	"aoc2016/internal/day22"
	"aoc2016/internal/day23"
	"aoc2016/internal/day24"
	"aoc2016/internal/day25"
)

func runDay(day int) {
	switch day {
	case 1:
		day01.Solve()
	case 2:
		day02.Solve()
	case 3:
		day03.Solve()
	case 4:
		day04.Solve()
	case 5:
		day05.Solve()
	case 6:
		day06.Solve()
	case 7:
		day07.Solve()
	case 8:
		day08.Solve()
	case 9:
		day09.Solve()
	case 10:
		day10.Solve()
	case 11:
		day11.Solve()
	case 12:
		day12.Solve()
	case 13:
		day13.Solve()
	case 14:
		day14.Solve()
	case 15:
		day15.Solve()
	case 16:
		day16.Solve()
	case 17:
		day17.Solve()
	case 18:
		day18.Solve()
	case 19:
		day19.Solve()
	case 20:
		day20.Solve()
	case 21:
		day21.Solve()
	case 22:
		day22.Solve()
	case 23:
		day23.Solve()
	case 24:
		day24.Solve()
	case 25:
		day25.Solve()
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cmd/main.go DAY")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	if day == 0 {
		for day <= 25 {
			day++
			runDay(day)
		}
	} else if day <= 25 {
		runDay(day)
	} else {
		panic(fmt.Errorf("invalid day: %d, expects a value between 0 to 25", day))
	}
}
