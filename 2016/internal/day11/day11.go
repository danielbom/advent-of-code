package day11

import (
	"container/heap"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

// Based on https://github.com/Kezzryn/Advent-of-Code/blob/main/2016/Day%2011/RTFElevator.cs

func parseContent(content string) ([]int, error) {
	names := make(map[string]int)
	floors := map[string]int{"first": 1, "second": 2, "third": 3, "fourth": 4}
	reFloor := regexp.MustCompile(`The ([a-z]+) floor contains`)
	reSep := regexp.MustCompile(`( and a| a|, and a|, a)`)
	result := make([]int, 0)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if !reFloor.MatchString(line) {
			return nil, fmt.Errorf("invalid floor format")
		}
		matches := reFloor.FindStringSubmatch(line)
		floor, ok := floors[matches[1]]
		if !ok {
			return nil, fmt.Errorf("invalid floor ordinal format")
		}
		line = line[len(matches[0]):]
		parts := reSep.Split(line, -1)
		for i, part := range parts {
			part := strings.TrimSpace(part)
			part = strings.TrimSuffix(part, ".")
			parts[i] = part
			if name, found := strings.CutSuffix(part, " generator"); found {
				if _, exists := names[name]; !exists {
					names[name] = len(names) * 2
					result = append(result, 0, 0)
				}
			}
			if name, found := strings.CutSuffix(part, "-compatible microchip"); found {
				if _, exists := names[name]; !exists {
					names[name] = len(names) * 2
					result = append(result, 0, 0)
				}
			}
		}
		for _, part := range parts {
			if name, found := strings.CutSuffix(part, " generator"); found {
				ix := names[name]
				result[ix+0] = floor
				continue
			}
			if name, found := strings.CutSuffix(part, "-compatible microchip"); found {
				ix := names[name]
				result[ix+1] = floor
				continue
			}
		}
	}
	return result, nil
}

func parseFile(filename string) ([]int, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

type State1 struct {
	components []uint
	floor      int
	steps      int
	count      int
}

func NewState1(cs []uint, floor, steps, count int) State1 {
	var s State1
	s.components = cs
	s.floor = floor
	s.steps = steps
	s.count = count
	return s
}

func state1FromFloors(floors []int) State1 {
	cs := make([]uint, 4)
	for i, f := range floors {
		cs[f-1] = cs[f-1] | (1 << i)
	}
	return NewState1(cs, 0, 0, len(floors)/2)
}

func (s State1) Copy() State1 {
	cs := make([]uint, len(s.components))
	copy(cs, s.components)
	return NewState1(cs, s.floor, s.steps, s.count)
}

func (s State1) Hash() string {
	var sb strings.Builder
	sb.Grow(64)
	sb.WriteByte(byte(int('1') + s.floor))
	for f := 0; f < 4; f++ {
		ps := 0
		cs := 0
		gs := 0
		for i := 0; i < s.count*2; i += 2 {
			hg := s.components[f]&uint(1<<(i+0)) > 0
			hc := s.components[f]&uint(1<<(i+1)) > 0
			if hg && !hc {
				gs++
			}
			if !hg && hc {
				cs++
			}
			if hg && hc {
				ps++
			}
		}
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(f))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(ps))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(cs))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(gs))
	}
	return sb.String()
}

func (s State1) String() string {
	var sb strings.Builder
	sb.WriteByte(byte(int('1') + s.floor))
	for f := 0; f < 4; f++ {
		sb.WriteByte('|')
		for i := 0; i < s.count*2; i++ {
			k := uint(1 << i)
			if s.components[f]&k > 0 {
				if i%2 == 0 {
					sb.WriteByte(byte(int('a') + (i / 2)))
				} else {
					sb.WriteByte(byte(int('A') + (i / 2)))
				}
			}
		}
	}
	return sb.String()
}

func (s State1) Done() bool {
	for f := 0; f < 3; f++ {
		if s.components[f] > 0 {
			return false
		}
	}
	return true
}

func (s State1) Valid() bool {
	for f := 0; f < 4; f++ {
		hasGs := false
		hasUs := false
		for i := 0; i < s.count; i++ {
			if s.components[f]&(1<<(i*2+0)) > 0 {
				hasGs = true
			} else if s.components[f]&(1<<(i*2+1)) > 0 {
				hasUs = true
			}
		}
		if hasGs && hasUs {
			return false
		}
	}
	return true
}

func (s *State1) MoveGenerator(ix, i, j int) {
	s.components[i] &= ^(1 << (ix*2 + 0))
	s.components[j] |= (1 << (ix*2 + 0))
}

func (s *State1) MoveMicrochip(ix, i, j int) {
	s.components[i] &= ^(1 << (ix*2 + 1))
	s.components[j] |= (1 << (ix*2 + 1))
}

func (s *State1) Move(k, i, j int) {
	if k%2 == 0 {
		s.MoveGenerator(k/2, i, j)
	} else {
		s.MoveMicrochip(k/2, i, j)
	}
}

type State2 struct {
	floors []int
	floor  int
	steps  int
}

func NewState2(floors []int, floor, steps int) State2 {
	var s State2
	s.floors = floors
	s.floor = floor
	s.steps = steps
	return s
}

func state2FromFloors(floors []int) State2 {
	return NewState2(floors, 1, 0)
}

func (s State2) Copy() State2 {
	fs := make([]int, len(s.floors))
	copy(fs, s.floors)
	return NewState2(fs, s.floor, s.steps)
}

func (s State2) Hash() string {
	var sb strings.Builder
	sb.Grow(64)
	sb.WriteByte(byte(int('0') + s.floor))
	fs := s.floors
	for f := 1; f <= 4; f++ {
		ps := 0
		cs := 0
		gs := 0
		for i := 0; i < len(fs); i += 2 {
			if fs[i] == f && fs[i+1] == f {
				ps++
			}
			if fs[i] == f && fs[i+1] != f {
				gs++
			}
			if fs[i] != f && fs[i+1] == f {
				cs++
			}
		}
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(f))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(ps))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(cs))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(gs))
	}
	return sb.String()
}

func (s State2) String() string {
	var sb strings.Builder
	for _, f := range s.floors {
		sb.WriteByte(byte(int('0') + f))
	}
	return sb.String()
}

func (s State2) Done() bool {
	for _, f := range s.floors {
		if f != 4 {
			return false
		}
	}
	return true
}

func (s State2) Valid() bool {
	for c := 1; c <= 4; c++ {
		hasGs := false
		hasUs := false
		for i := 0; i < len(s.floors); i += 2 {
			if s.floors[i] == c {
				hasGs = true
			}
			if s.floors[i] != c && s.floors[i+1] == c {
				hasUs = true
			}
		}
		if hasGs && hasUs {
			return false
		}
	}
	return true
}

const (
	INDEX_ELEVATOR = 1
	INDEX_STEPS    = 0
	FLOOR_BOTTOM   = 1
	FLOOR_TOP      = 4
	START_ITEMS    = 2
)

type State3 []int

func state3FromFloors(floors []int) State3 {
	s := make([]int, len(floors)+2)
	copy(s[2:], floors)
	s[INDEX_ELEVATOR] = 1
	s[INDEX_STEPS] = 0
	return State3(s)
}

func (s State3) Done() bool {
	for _, f := range s[START_ITEMS:] {
		if f != FLOOR_TOP {
			return false
		}
	}
	return true
}

func (s State3) Valid() bool {
	for floor := FLOOR_BOTTOM; floor <= FLOOR_TOP; floor++ {
		hasGs := false
		hasUs := false
		for item := START_ITEMS; item < len(s); item += 2 {
			if s[item] == floor {
				hasGs = true
			}
			if s[item+1] == floor && s[item] != floor {
				hasUs = true
			}
			if hasGs && hasUs {
				return false
			}
		}
	}
	return true
}

func (s State3) Hash() string {
	var sb strings.Builder
	sb.Grow(64)
	fs := []int(s)
	sb.WriteString(strconv.Itoa(fs[INDEX_ELEVATOR]))
	for f := FLOOR_BOTTOM; f <= FLOOR_TOP; f++ {
		ps := 0
		cs := 0
		gs := 0
		for i := START_ITEMS; i < len(fs); i += 2 {
			if fs[i] == f && fs[i+1] == f {
				ps++
			}
			if fs[i] == f && fs[i+1] != f {
				gs++
			}
			if fs[i] != f && fs[i+1] == f {
				cs++
			}
		}
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(f))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(ps))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(cs))
		sb.WriteByte('|')
		sb.WriteString(strconv.Itoa(gs))
	}
	return sb.String()
}

type Goal interface {
	Steps() int
	Done() bool
	Valid() bool
	Hash() string
	WithSteps(steps int) Goal
	EnqueueNextSteps(c *CounterX)
}

func (s State1) WithSteps(steps int) Goal {
	s.steps = steps
	return s
}
func (s State1) Steps() int {
	return s.steps
}
func (s State1) EnqueueNextSteps(c *CounterX) {
	for d := 1; d >= -1; d -= 2 {
		fc := s.floor
		fn := s.floor + d
		if 0 <= fn && fn <= 3 {
			for i := 0; i < s.count*2; i++ {
				if s.components[fc]&(1<<i) > 0 {
					s1 := s.Copy()
					s1.steps++
					s1.floor = fn
					s1.Move(i, fc, fn)
					c.Enqueue(s1)
					for j := i + 1; j < s.count*2; j++ {
						if s.components[fc]&(1<<j) > 0 && ((i%2 == 0 && i+1 == j || j%2 == 0) || (i%2 == 1 && j%2 == 1)) {
							s2 := s1.Copy()
							s2.Move(j, fc, fn)
							c.Enqueue(s2)
						}
					}
				}
			}
		}
	}
}

func (s State2) WithSteps(steps int) Goal {
	s.steps = steps
	return s
}
func (s State2) Steps() int {
	return s.steps
}
func (s State2) EnqueueNextSteps(c *CounterX) {
	for d := -1; d <= 1; d += 2 {
		fc := s.floor
		fn := s.floor + d
		if 1 <= fn && fn <= 4 {
			for i, f1 := range s.floors {
				if f1 == fc {
					s1 := s.Copy()
					s1.steps++
					s1.floor = fn
					s1.floors[i] = fn
					c.Enqueue(s1)
					for j, f2 := range s.floors {
						if i < j && f2 == fc {
							s2 := s1.Copy()
							s2.floors[j] = fn
							c.Enqueue(s2)
						}
					}
				}
			}
		}
	}
}

func (s State3) WithSteps(steps int) Goal {
	r := make([]int, len(s))
	copy(r, s)
	r[INDEX_STEPS] = steps
	return State3(r)
}
func (s State3) Steps() int {
	return s[INDEX_STEPS]
}
func (s State3) EnqueueNextSteps(c *CounterX) {
	for nextFloor := -1; nextFloor <= 1; nextFloor += 2 {
		if s[INDEX_ELEVATOR]+nextFloor < FLOOR_BOTTOM || s[INDEX_ELEVATOR]+nextFloor > FLOOR_TOP {
			continue
		}
		for item1 := START_ITEMS; item1 < len(s); item1++ {
			if s[item1] == s[INDEX_ELEVATOR] {
				nextStep1 := make([]int, len(s))
				copy(nextStep1, s)
				nextStep1[INDEX_STEPS]++
				nextStep1[INDEX_ELEVATOR] += nextFloor
				nextStep1[item1] += nextFloor
				c.Enqueue(State3(nextStep1))
				for item2 := START_ITEMS; item2 < len(s); item2++ {
					if nextStep1[item2] == s[INDEX_ELEVATOR] {
						nextStep2 := make([]int, len(s))
						copy(nextStep2, nextStep1)
						nextStep2[item2] += nextFloor
						c.Enqueue(State3(nextStep2))
					}
				}
			}
		}
	}
}

type PriorityQueueX []Goal

func (pq PriorityQueueX) Len() int {
	return len(pq)
}
func (pq PriorityQueueX) Less(i, j int) bool {
	return pq[i].Steps() < pq[j].Steps()
}
func (pq PriorityQueueX) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueueX) Push(x any) {
	item := x.(Goal)
	*pq = append(*pq, item)
}
func (pq *PriorityQueueX) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type CounterX struct {
	pq   PriorityQueueX
	seen map[string]int
	best Goal
}

func NewCounterX(initialState Goal) CounterX {
	var c CounterX
	c.seen = make(map[string]int)
	c.pq = make([]Goal, 0, 1024)
	c.pq = append(c.pq, initialState)
	c.best = initialState.WithSteps(1 << 32)
	return c
}

func (c *CounterX) Enqueue(s Goal) {
	if s.Valid() {
		key := s.Hash()
		steps, found := c.seen[key]
		if found {
			if s.Steps() < steps {
				c.seen[key] = s.Steps()
				heap.Push(&c.pq, s)
			}
		} else {
			c.seen[key] = s.Steps()
			heap.Push(&c.pq, s)
		}
	}
}
func (c *CounterX) Run() {
	for len(c.pq) > 0 {
		s := heap.Pop(&c.pq).(Goal)
		if s.Steps() < c.best.Steps() {
			if s.Done() {
				c.best = s
				continue
			}
			s.EnqueueNextSteps(c)
		}
	}
}

const VERSION int = 1

func goalFromFloors(floors []int) Goal {
	switch VERSION {
	case 1:
		return state1FromFloors(floors)
	case 2:
		return state2FromFloors(floors)
	case 3:
		return state3FromFloors(floors)
	}
	var none Goal
	return none
}

func countStepsX(floors []int) int {
	c := NewCounterX(goalFromFloors(floors))
	c.Run()
	return c.best.Steps()
}

func part1(floors []int) int {
	return countStepsX(floors)
}

func part2(floors []int) int {
	// Extras to the 1st floor:
	//    An elerium generator.
	//    An elerium-compatible microchip.
	//    A dilithium generator.
	//    A dilithium-compatible microchip.
	floors2 := make([]int, len(floors)+4)
	copy(floors2, floors)
	for i := len(floors); i < len(floors2); i++ {
		floors2[i] = 1
	}
	return countStepsX(floors2)
}

func Solve() {
	input, err := parseFile("./inputs/day-11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 11")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
