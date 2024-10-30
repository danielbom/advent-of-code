package main

import (
	"fmt"
	"os"
	"strconv"

	"aoc2016/internal/day01"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cmd/main.go DAY")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	switch day {
	case 1:
		day01.Solve()
	}
}
