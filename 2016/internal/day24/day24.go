package day24

import (
	"fmt"
	"strings"

	"aoc2016/internal/utils"
)

type Item struct {
	X, Y     int
	Wall     bool
	Goal     int
	Distance int
}

type Map struct {
	grid [][]Item
	rows int
	cols int
}

func parseContent(content string) (Map, error) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	var m Map
	m.rows = len(lines)
	m.grid = make([][]Item, 0, m.rows)
	for y, line := range lines {
		cols := len(line)
		row := make([]Item, cols)
		for x := 0; x < cols; x++ {
			item := Item{X: x, Y: y, Wall: true, Goal: -1, Distance: -1}
			if x < len(line) && line[x] != '#' {
				item.Wall = false
				goal := int(line[x]) - int('0')
				if 0 <= goal && goal <= 9 {
					item.Goal = goal
				}
			}
			row[x] = item
		}
		m.grid = append(m.grid, row)
		m.cols = cols
	}
	return m, nil
}

func parseFile(filename string) (Map, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return Map{}, err
	}
	return parseContent(content)
}

func checkNextItem(item, next *Item, stack []*Item) []*Item {
	if !next.Wall {
		if next.Distance == -1 || next.Distance > item.Distance+1 {
			next.Distance = item.Distance + 1
			return append(stack, next)
		}
	}
	return stack
}

func calculateDistances(m Map, location int) []int {
	stack := make([]*Item, 0, 4)
	for y := range m.rows {
		for x := range m.cols {
			item := &m.grid[y][x]
			item.Distance = -1
			if item.Goal == location {
				item.Distance = 0
				stack = append(stack, item)
			}
		}
	}
	if len(stack) == 0 {
		return nil
	}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if item.Y > 0 {
			stack = checkNextItem(item, &m.grid[item.Y-1][item.X], stack)
		}
		if item.Y < m.rows {
			stack = checkNextItem(item, &m.grid[item.Y+1][item.X], stack)
		}
		if item.X > 0 {
			stack = checkNextItem(item, &m.grid[item.Y][item.X-1], stack)
		}
		if item.X < m.cols {
			stack = checkNextItem(item, &m.grid[item.Y][item.X+1], stack)
		}
	}
	distances := make([]int, 10)
	for y := range m.rows {
		for x := range m.cols {
			if m.grid[y][x].Goal >= 0 {
				distances[m.grid[y][x].Goal] = m.grid[y][x].Distance
			}
		}
	}
	return distances
}

func createDistancesGraph(m Map) [][]int {
	allDistances := make([][]int, 0, 10)
	mi := 0
	for i := 0; i < 10; i++ {
		mi = i
		distances := calculateDistances(m, i)
		if distances == nil {
			break
		}
		allDistances = append(allDistances, distances)
	}
	for i := range mi {
		allDistances[i] = allDistances[i][:mi]
	}
	return allDistances
}

func findBestPathRec(g [][]int, current, visited, distance, count int, bestDistance *int, willReturn bool) {
	if count == len(g) {
		nextDistance := distance
		if willReturn {
			nextDistance += g[current][0]
		}
		if *bestDistance == -1 || *bestDistance > nextDistance {
			*bestDistance = nextDistance
		}
	} else {
		for next := range g {
			nextBit := 1 << next
			if visited&nextBit == 0 {
				nextDistance := distance + g[current][next]
				if *bestDistance == -1 || *bestDistance > nextDistance {
					findBestPathRec(g, next, visited|nextBit, nextDistance, count+1, bestDistance, willReturn)
				}
			}
		}
	}
}

func findBestPath(g [][]int, begin int, willReturn bool) int {
	result := -1
	findBestPathRec(g, begin, 1<<begin, 0, 1, &result, willReturn)
	return result
}

func part1(m Map) int {
	g := createDistancesGraph(m)
	return findBestPath(g, 0, false)
}

func part2(m Map) int {
	g := createDistancesGraph(m)
	return findBestPath(g, 0, true)
}

func Solve() {
	input, err := parseFile("./inputs/day-24.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 24")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
