package day19

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

func parseFile(filename string) (int, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return 0, err
	}
	content = strings.TrimSpace(content)
	return strconv.Atoi(content)
}

// Double Linked List Indexed
type Elf struct {
	Id    int
	Prev  int
	Next  int
	Gifts int
}

type Party []Elf

func NewParty(elvesCount int) Party {
	elves := make([]Elf, elvesCount)
	for i, _ := range elves {
		elves[i].Id = i + 1
		elves[i].Gifts = 1
		if i == 0 {
			elves[i].Prev = elvesCount - 1
		} else {
			elves[i].Prev = i - 1
		}
		if i == elvesCount-1 {
			elves[i].Next = 0
		} else {
			elves[i].Next = i + 1
		}
	}
	return Party(elves)
}

func (party Party) Winner1() int {
	current := 0
	for i := 1; i < len(party); i++ {
		{
			stolen := party[current].Next
			party[current].Gifts += party[stolen].Gifts
			party[stolen].Gifts = 0
			party[party[stolen].Prev].Next = party[stolen].Next
			party[party[stolen].Next].Prev = party[stolen].Prev
		}
		current = party[current].Next
	}
	return party[current].Id
}

func (party Party) Winner2() int {
	current := 0
	count := len(party)
	center := count / 2
	for count > 1 {
		{
			stolen := center
			party[current].Gifts += party[stolen].Gifts
			party[stolen].Gifts = 0
			party[party[stolen].Prev].Next = party[stolen].Next
			party[party[stolen].Next].Prev = party[stolen].Prev
			center = party[center].Next
			if count%2 == 1 {
				center = party[center].Next
			}
		}
		current = party[current].Next
		count--
	}
	return party[current].Id
}

func part1(elvesCount int) int {
	party := NewParty(elvesCount)
	return party.Winner1()
}

func part2(elvesCount int) int {
	party := NewParty(elvesCount)
	return party.Winner2()
}

func Solve() {
	input, err := parseFile("./inputs/day-19.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 19")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
