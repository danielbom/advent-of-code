package day21

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

type Instruction interface {
	ApplyBuf(src, dst []byte)
	Reversed() Instruction
}

type SwapPosition struct {
	X int
	Y int
}
type SwapLetter struct {
	A byte
	B byte
}
type RotateLeft struct {
	Steps int
}
type RotateRight struct {
	Steps int
}
type RotateOnLetterReverse struct {
	A byte
}
type RotateOnLetter struct {
	A byte
}
type ReversePosition struct {
	X int
	Y int
}
type MovePosition struct {
	X int
	Y int
}

func applySwapPosition(x, y int, src, dst []byte) {
	for i, ch := range src {
		switch i {
		case x:
			dst[i] = src[y]
		case y:
			dst[i] = src[x]
		default:
			dst[i] = ch
		}
	}
}
func (inst SwapPosition) ApplyBuf(src, dst []byte) {
	applySwapPosition(inst.X, inst.Y, src, dst)
}
func (inst SwapPosition) Reversed() Instruction {
	return inst
}

func applySwapLetter(a, b byte, src, dst []byte) {
	for i, ch := range src {
		switch ch {
		case a:
			dst[i] = b
		case b:
			dst[i] = a
		default:
			dst[i] = ch
		}
	}
}
func (inst SwapLetter) ApplyBuf(src, dst []byte) {
	applySwapLetter(inst.A, inst.B, src, dst)
}
func (inst SwapLetter) Reversed() Instruction {
	return inst
}

func applyRotateLeft(steps int, src, dst []byte) {
	if steps < 0 {
		steps = len(src) + steps
	}
	n := len(src)
	steps = steps % n
	if steps == 0 {
		copy(dst, src)
		return
	}
	for i, _ := range src {
		dst[i] = src[(i+steps)%n]
	}
}
func (inst RotateLeft) ApplyBuf(src, dst []byte) {
	applyRotateLeft(inst.Steps, src, dst)
}
func (inst RotateLeft) Reversed() Instruction {
	return RotateRight{Steps: inst.Steps}
}

func applyRotateRight(steps int, src, dst []byte) {
	n := len(src)
	steps = steps % n
	if steps == 0 {
		copy(dst, src)
		return
	}
	for i, _ := range src {
		dst[i] = src[(i-steps+n)%n]
	}
}
func (inst RotateRight) ApplyBuf(src, dst []byte) {
	applyRotateRight(inst.Steps, src, dst)
}
func (inst RotateRight) Reversed() Instruction {
	return RotateLeft{Steps: inst.Steps}
}

const VERSION int = 2

func applyRotateOnLetterReverse(a byte, src, dst []byte) {
	// OBS: The reverse of rotate on letter is not unique.
	//      This function only returns a single solution.
	if VERSION == 1 {
		index := slices.Index(src, a)
		if index == 0 {
			applyRotateLeft(1, src, dst)
		} else if index%2 == 1 {
			applyRotateLeft((index+1)/2, src, dst)
		} else {
			applyRotateLeft(((index+4)/2)-5, src, dst)
		}
		{
			tmp := make([]byte, len(src))
			applyRotateOnLetter(a, dst, tmp)
			if slices.Compare(tmp, src) != 0 {
				panic("unreachable")
			}
		}
	} else {
		buf := make([]byte, len(src))
		for i := 0; i < len(src); i++ {
			applyRotateLeft(i, src, dst)
			applyRotateOnLetter(a, dst, buf)
			if slices.Compare(buf, src) == 0 {
				return
			}
		}
		panic("unreachable")
	}
}
func (inst RotateOnLetterReverse) ApplyBuf(src, dst []byte) {
	applyRotateOnLetterReverse(inst.A, src, dst)
}
func (inst RotateOnLetterReverse) Reversed() Instruction {
	return RotateOnLetter{A: inst.A}
}

func applyRotateOnLetter(a byte, src, dst []byte) {
	index := slices.Index(src, a)
	steps := index + 1
	if index >= 4 {
		steps++
	}
	applyRotateRight(steps, src, dst)
}
func (inst RotateOnLetter) ApplyBuf(src, dst []byte) {
	applyRotateOnLetter(inst.A, src, dst)
}
func (inst RotateOnLetter) Reversed() Instruction {
	return RotateOnLetterReverse{A: inst.A}
}

func applyReversePosition(x, y int, src, dst []byte) {
	for i, ch := range src {
		if x <= i && i <= y {
			dst[i] = src[y-(i-x)]
		} else {
			dst[i] = ch
		}
	}
}
func (inst ReversePosition) ApplyBuf(src, dst []byte) {
	applyReversePosition(inst.X, inst.Y, src, dst)
}
func (inst ReversePosition) Reversed() Instruction {
	return inst
}

func applyMovePosition(x, y int, src, dst []byte) {
	if x == y {
		copy(dst, src)
		return
	}
	if x < y {
		for i, ch := range src {
			if x <= i && i < y {
				dst[i] = src[i+1]
			} else if i == y {
				dst[i] = src[x]
			} else {
				dst[i] = ch
			}
		}
	} else {
		for i, ch := range src {
			if y < i && i <= x {
				dst[i] = src[i-1]
			} else if i == y {
				dst[i] = src[x]
			} else {
				dst[i] = ch
			}
		}
	}
}
func (inst MovePosition) ApplyBuf(src, dst []byte) {
	applyMovePosition(inst.X, inst.Y, src, dst)
}
func (inst MovePosition) Reversed() Instruction {
	return MovePosition{X: inst.Y, Y: inst.X}
}

func ApplyInstruction(inst Instruction, input string) string {
	dst := make([]byte, len(input))
	inst.ApplyBuf([]byte(input), dst)
	return string(dst)
}

func Prefixed(text, prefix string) (string, bool) {
	if strings.HasPrefix(text, prefix) {
		return text[len(prefix):], true
	}
	return "", false
}

func parseContent(content string) ([]Instruction, error) {
	content = strings.TrimSpace(content)
	lines := strings.Split(content, "\n")
	result := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if rest, found := Prefixed(line, "swap position "); found {
			first, second, found := strings.Cut(rest, " with position ")
			if !found {
				return nil, fmt.Errorf("invalid SwapPosition instruction")
			}
			x, err := strconv.Atoi(first)
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(second)
			if err != nil {
				return nil, err
			}
			inst := SwapPosition{X: x, Y: y}
			result = append(result, inst)
			continue
		}
		if rest, found := Prefixed(line, "swap letter "); found {
			first, second, found := strings.Cut(rest, " with letter ")
			if !found {
				return nil, fmt.Errorf("invalid SwapPosition instruction")
			}
			a := first[0]
			b := second[0]
			inst := SwapLetter{A: a, B: b}
			result = append(result, inst)
			continue
		}
		if rest, found := Prefixed(line, "rotate left "); found {
			value := strings.TrimSuffix(rest, "s")
			value = strings.TrimSuffix(value, " step")
			steps, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			inst := RotateLeft{Steps: steps}
			result = append(result, inst)
			continue
		}
		if rest, found := Prefixed(line, "rotate right "); found {
			value := strings.TrimSuffix(rest, "s")
			value = strings.TrimSuffix(value, " step")
			steps, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			inst := RotateRight{Steps: steps}
			result = append(result, inst)
			continue
		}
		if rest, found := Prefixed(line, "rotate based on position of letter "); found {
			a := rest[0]
			inst := RotateOnLetter{A: a}
			result = append(result, inst)
			continue
		}
		if rest, found := Prefixed(line, "reverse positions "); found {
			first, second, found := strings.Cut(rest, " through ")
			if !found {
				return nil, fmt.Errorf("invalid SwapPosition instruction")
			}
			x, err := strconv.Atoi(first)
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(second)
			if err != nil {
				return nil, err
			}
			inst := ReversePosition{X: x, Y: y}
			result = append(result, inst)
			continue
		}
		if rest, found := Prefixed(line, "move position "); found {
			first, second, found := strings.Cut(rest, " to position ")
			if !found {
				return nil, fmt.Errorf("invalid SwapPosition instruction")
			}
			x, err := strconv.Atoi(first)
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(second)
			if err != nil {
				return nil, err
			}
			inst := MovePosition{X: x, Y: y}
			result = append(result, inst)
			continue
		}
		return nil, fmt.Errorf("unknown instruction: %s", line)
	}
	return result, nil
}

func parseFile(filename string) ([]Instruction, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContent(content)
}

func applyInstructions(is []Instruction, initial string) string {
	buf1 := make([]byte, len(initial))
	buf2 := make([]byte, len(initial))
	copy(buf1, initial)
	for _, inst := range is {
		inst.ApplyBuf(buf1, buf2)
		tmp := buf1
		buf1 = buf2
		buf2 = tmp
	}
	return string(buf1)
}

func unapplyInstructions(is []Instruction, initial string) string {
	n := len(is)
	buf1 := make([]byte, len(initial))
	buf2 := make([]byte, len(initial))
	copy(buf1, initial)
	for i, _ := range is {
		inst := is[n-1-i].Reversed()
		inst.ApplyBuf(buf1, buf2)
		tmp := buf1
		buf1 = buf2
		buf2 = tmp
	}
	return string(buf1)
}

func part1(is []Instruction) string {
	return applyInstructions(is, "abcdefgh")
}

func part2(is []Instruction) string {
	return unapplyInstructions(is, "fbgdceah")
	//return applyInstructions(is, "bdgheacf")
	//return applyInstructions(is, "bdgheacf")
}

func Solve() {
	input, err := parseFile("./inputs/day-21.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 21")
	utils.TimeIt("Part 1:", "%s", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%s", func() any { return part2(input) })
}
