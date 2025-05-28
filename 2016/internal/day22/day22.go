package day22

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

type FileSystem struct {
	Used int
	Size int
}

func (fs FileSystem) Available() int {
	return fs.Size - fs.Used
}

type FileSystemItem struct {
	X  int
	Y  int
	FS FileSystem
}

func extractFileSystemPosition(name string) (int, int, error) {
	position := strings.TrimPrefix(name, "/dev/grid/node-")
	x, y, found := strings.Cut(position, "-")
	if found && len(x) > 1 && len(y) > 1 && x[0] == 'x' && y[0] == 'y' {
		xVal, err := strconv.Atoi(x[1:])
		if err != nil {
			return -1, -1, err
		}
		yVal, err := strconv.Atoi(y[1:])
		if err != nil {
			return -1, -1, err
		}
		return xVal, yVal, nil
	}
	return -1, -1, fmt.Errorf("invalid file system name")
}

func cutSpaces(text string) (string, string, bool) {
	first, second, found := strings.Cut(text, " ")
	first = strings.TrimSpace(first)
	second = strings.TrimSpace(second)
	return first, second, found
}

// Ex: Filesystem              Size  Used  Avail  Use%
// --- /dev/grid/node-x0-y0     85T   66T    19T   77%
func parseFileSystemItem(text string) (item FileSystemItem, err error) {
	var rest, name, size, used string
	var found bool
	text = strings.TrimSpace(text)
	name, rest, found = cutSpaces(text)
	if !found {
		return item, fmt.Errorf("invalid file system format")
	}
	size, rest, found = cutSpaces(rest)
	size = strings.TrimSuffix(size, "T")
	if !found {
		return item, fmt.Errorf("invalid file system format")
	}
	used, _, found = cutSpaces(rest)
	used = strings.TrimSuffix(used, "T")
	if !found {
		return item, fmt.Errorf("invalid file system format")
	}
	var fs FileSystem
	fs.Size, err = strconv.Atoi(size)
	if err != nil {
		return item, err
	}
	fs.Used, err = strconv.Atoi(used)
	if err != nil {
		return item, err
	}
	x, y, err := extractFileSystemPosition(name)
	if err != nil {
		return item, err
	}
	item.X = x
	item.Y = y
	item.FS = fs
	return item, nil
}

func parseContent(content string) ([]FileSystemItem, error) {
	content = strings.TrimSpace(content)
	lines := strings.Split(content, "\n")
	lines = lines[2:]
	result := make([]FileSystemItem, 0, len(lines))
	for _, line := range lines {
		fs, err := parseFileSystemItem(line)
		if err != nil {
			return nil, err
		}
		result = append(result, fs)
	}
	return result, nil
}

func parseFile(filename string) ([]FileSystemItem, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func part1(is []FileSystemItem) int {
	count := 0
	for i, item1 := range is {
		if item1.FS.Used > 0 {
			for j, item2 := range is {
				if i != j && item1.FS.Used <= item2.FS.Available() {
					count++
				}
			}
		}
	}
	return count
}

type Point struct {
	Y int
	X int
}

type Step struct {
	Next Point
	Dist int
}

func NewStep(next Point, dist int) Step {
	return Step{Next: next, Dist: dist}
}

type PriorityQueue []Step

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dist <= pq[j].Dist
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x any) {
	item := x.(Step)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func NewPoint(y, x int) Point {
	return Point{Y: y, X: x}
}

type Grid struct {
	grid [][]FileSystem
	rows int
	cols int
}

func MakeGrid(is []FileSystemItem) Grid {
	var rows, cols int
	var grid [][]FileSystem
	{
		maxx := -1
		maxy := -1
		for _, item := range is {
			maxx = utils.Max(maxx, item.X)
			maxy = utils.Max(maxy, item.Y)
		}
		grid = make([][]FileSystem, maxy+1)
		for i, _ := range grid {
			grid[i] = make([]FileSystem, maxx+1)
		}
		rows = maxy + 1
		cols = maxx + 1
		for _, item := range is {
			grid[item.Y][item.X] = item.FS
		}
	}
	return Grid{grid: grid, rows: rows, cols: cols}
}

func MinDistance(g Grid, target Point) [][]int {
	grid, rows, cols := g.grid, g.rows, g.cols
	dist := make([][]int, rows)
	for i, _ := range dist {
		dist[i] = make([]int, cols)
		for j, _ := range dist[i] {
			dist[i][j] = 1 << 32
		}
	}
	dist[target.Y][target.X] = 0
	pq := make(PriorityQueue, 1)
	pq = append(pq, NewStep(target, 0))
	for len(pq) > 0 {
		curr := heap.Pop(&pq).(Step)
		y := curr.Next.Y
		x := curr.Next.X
		if y > 0 && grid[y-1][x].Used <= grid[y][x].Size && dist[y][x]+1 < dist[y-1][x] {
			next := NewStep(NewPoint(y-1, x), dist[y][x]+1)
			dist[y-1][x] = dist[y][x] + 1
			heap.Push(&pq, next)
		}
		if y < rows-1 && grid[y+1][x].Used <= grid[y][x].Size && dist[y][x]+1 < dist[y+1][x] {
			next := NewStep(NewPoint(y+1, x), dist[y][x]+1)
			dist[y+1][x] = dist[y][x] + 1
			heap.Push(&pq, next)
		}
		if x > 0 && grid[y][x-1].Used <= grid[y][x].Size && dist[y][x]+1 < dist[y][x-1] {
			next := NewStep(NewPoint(y, x-1), dist[y][x]+1)
			dist[y][x-1] = dist[y][x] + 1
			heap.Push(&pq, next)
		}
		if x < cols-1 && grid[y][x+1].Used <= grid[y][x].Size && dist[y][x]+1 < dist[y][x+1] {
			next := NewStep(NewPoint(y, x+1), dist[y][x]+1)
			dist[y][x+1] = dist[y][x] + 1
			heap.Push(&pq, next)
		}
	}
	return dist
}

const DEBUG bool = false

func part2(is []FileSystemItem) int {
	g := MakeGrid(is)
	grid, rows, cols := g.grid, g.rows, g.cols
	goal := NewPoint(0, g.cols-1)
	dist := MinDistance(g, goal)
	dist1 := 1 << 32
	best := NewPoint(-1, -1)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if y > 0 && grid[y-1][x].Used <= grid[y][x].Available() {
				if dist[y][x] < dist1 {
					best = NewPoint(y-1, x)
					dist1 = dist[y][x]
				}
			}
			if y < rows-1 && grid[y+1][x].Used <= grid[y][x].Available() {
				if dist[y][x] < dist1 {
					best = NewPoint(y-1, x)
					dist1 = dist[y][x]
				}
			}
			if x > 0 && grid[y][x-1].Used <= grid[y][x].Available() {
				if dist[y][x] < dist1 {
					best = NewPoint(y-1, x)
					dist1 = dist[y][x]
				}
			}
			if x < cols-1 && grid[y][x+1].Used <= grid[y][x].Available() {
				if dist[y][x] < dist1 {
					best = NewPoint(y-1, x)
					dist1 = dist[y][x]
				}
			}
		}
	}
	dist2 := dist[0][0]
	if DEBUG {
		path := make(map[Point]bool)
		curr := best
		for dist[curr.Y][curr.X] != dist1 {
			y, x := curr.Y, curr.X
			path[curr] = true
			if y > 0 && dist[y][x]-1 == dist[y-1][x] {
				curr = NewPoint(y-1, x)
				continue
			}
			if y < rows-1 && dist[y][x]-1 == dist[y+1][x] {
				curr = NewPoint(y+1, x)
				continue
			}
			if x > 0 && dist[y][x]-1 == dist[y][x-1] {
				curr = NewPoint(y, x-1)
				continue
			}
			if x < cols-1 && dist[y][x]-1 == dist[y][x+1] {
				curr = NewPoint(y, x+1)
				continue
			}
			break
		}
		fmt.Println(dist1, dist2)
		mode := 1
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if x > 0 {
					fmt.Printf(" ")
				}
				p := NewPoint(y, x)
				if mode == 1 {
					if !path[p] {
						fmt.Printf("%3s", " ")
					} else {
						fmt.Printf("%3d", dist1-dist[y][x])
					}
				} else {
					if dist[y][x] == 1<<32 {
						fmt.Printf("%3s", "- ")
					} else {
						fmt.Printf("%3d", dist1-dist[y][x])
					}
				}
			}
			fmt.Println()
		}
	}
	return (dist1 + 1) + (dist2-1)*5
}

func Solve() {
	input, err := parseFile("./inputs/day-22.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 22")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
