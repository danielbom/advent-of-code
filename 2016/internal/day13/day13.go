package day13

import (
	"container/heap"
	"fmt"
	"math/bits"
	"os"
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

func parseFile(filename string) (int, error) {
	content, err := readAllFile(filename)
	if err != nil {
		return 0, err
	}
	content = strings.TrimSpace(content)
	return strconv.Atoi(content)
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p1 Point) Distance(p2 Point) int {
	return utils.Abs(p1.X-p2.X) + utils.Abs(p1.Y-p2.Y)
}

type MapGenerator struct {
	Seed int
}

func NewMapGenerator(seed int) MapGenerator {
	var g MapGenerator
	g.Seed = seed
	return g
}

func (g *MapGenerator) IsWall(p Point) bool {
	x, y := p.X, p.Y
	v := x*x + 3*x + 2*x*y + y + y*y
	v += g.Seed
	count := bits.OnesCount(uint(v))
	return count%2 == 1
}

// https://pkg.go.dev/container/heap#example-package-PriorityQueue
type Item struct {
	Point    Point
	count    int
	distance int
	index    int
}

func NewItem(count, distance int, point Point) Item {
	var i Item
	i.count = count
	i.Point = point
	i.distance = distance
	i.index = 0
	return i
}

type PriorityQueue1 []*Item

func (pq PriorityQueue1) Len() int {
	return len(pq)
}
func (pq PriorityQueue1) Less(i, j int) bool {
	return pq[i].distance <= pq[j].distance
}
func (pq PriorityQueue1) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue1) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue1) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type CountSteps struct {
	expected Point
	mg       MapGenerator
	pq       PriorityQueue1
	seen     map[Point]bool
	minCount int
}

func NewCountSteps(seed, x, y int) CountSteps {
	var cs CountSteps
	cs.expected = NewPoint(x, y)
	cs.mg = NewMapGenerator(seed)
	cs.seen = make(map[Point]bool)
	cs.minCount = 1 << 32
	return cs
}

func (cs *CountSteps) PushPoint(count int, point Point) {
	if point.Y < 0 || point.X < 0 {
		return
	}
	if cs.seen[point] {
		return
	}
	if cs.mg.IsWall(point) {
		return
	}
	cs.seen[point] = true
	item := NewItem(count, point.Distance(cs.expected), point)
	heap.Push(&cs.pq, &item)
	heap.Fix(&cs.pq, item.index)
}

func (cs *CountSteps) Run() {
	for len(cs.pq) > 0 {
		item := heap.Pop(&cs.pq).(*Item)
		if item.Point == cs.expected {
			cs.minCount = item.count
			return
		}
		if item.count+1 > cs.minCount {
			return
		}
		for k := -1; k <= 1; k += 2 {
			cs.PushPoint(item.count+1, NewPoint(item.Point.X+k, item.Point.Y))
			cs.PushPoint(item.count+1, NewPoint(item.Point.X, item.Point.Y+k))
		}
	}
}

func countSteps(seed, expectedX, expectedY int) int {
	cs := NewCountSteps(seed, expectedX, expectedY)
	cs.PushPoint(0, NewPoint(1, 1))
	cs.Run()
	return cs.minCount
}

type PriorityQueue2 [](*Item)

func (pq PriorityQueue2) Len() int {
	return len(pq)
}
func (pq PriorityQueue2) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}
func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue2) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue2) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type CountLocations struct {
	expected  int
	seen      map[Point]int
	pq        PriorityQueue2
	mg        MapGenerator
	locations map[Point]bool
}

func NewCountLocations(seed, count int) CountLocations {
	var cl CountLocations
	cl.expected = count
	cl.mg = NewMapGenerator(seed)
	cl.seen = make(map[Point]int)
	cl.locations = make(map[Point]bool)
	return cl
}

func (cl *CountLocations) PushPoint(count int, point Point) {
	if point.Y < 0 || point.X < 0 {
		return
	}
	if cl.mg.IsWall(point) {
		return
	}
	if cl.expected < count {
		return
	}
	if prev, found := cl.seen[point]; found && prev <= count {
		return
	}
	cl.seen[point] = count
	cl.locations[point] = true
	item := NewItem(count, 0, point)
	heap.Push(&cl.pq, &item)
	heap.Fix(&cl.pq, item.index)
}

func (cl *CountLocations) Run() {
	for len(cl.pq) > 0 {
		item := heap.Pop(&cl.pq).(*Item)
		for k := -1; k <= 1; k += 2 {
			cl.PushPoint(item.count+1, NewPoint(item.Point.X+k, item.Point.Y))
			cl.PushPoint(item.count+1, NewPoint(item.Point.X, item.Point.Y+k))
		}
	}
}

func (cl *CountLocations) Draw() {
	count := cl.expected * 2 / 3
	for y := 0; y <= count; y++ {
		for x := 0; x <= count; x++ {
			p := NewPoint(x, y)
			if cl.mg.IsWall(p) {
				fmt.Print(" ## ")
			} else {
				d, found := cl.seen[p]
				if !found {
					fmt.Print("    ")
				} else {
					fmt.Printf(" %2d ", d)
				}
			}
		}
		fmt.Println()
	}
}

func countLocations(seed, count int) int {
	cl := NewCountLocations(seed, count)
	cl.PushPoint(0, NewPoint(1, 1))
	cl.Run()
	//cl.Draw()
	return len(cl.locations)
}

func part1(seed int) int {
	return countSteps(seed, 31, 39)
}

func part2(seed int) int {
	return countLocations(seed, 50)
}

func Solve() {
	input, err := parseFile("./inputs/day-13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 13")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
