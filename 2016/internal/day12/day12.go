package day12

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"aoc2016/internal/utils"
)

type Copy struct {
	Src string
	Val int
	Dst string
}

type Increment struct {
	Reg string
}

type Decrement struct {
	Reg string
}

type JumpNotZero struct {
	Reg  string
	Val  int
	Jump int
}

type Instruction struct {
	As interface{}
}

func checkRegister(name string) (string, error) {
	if len(name) != 1 {
		return "", fmt.Errorf("register '%s' should have 1 character", name)
	}
	if !unicode.IsLower(rune(name[0])) {
		return "", fmt.Errorf("register '%s' should be in lower case", name)
	}
	return name, nil
}

func parseInstruction(text string) (Instruction, error) {
	var result Instruction
	parts := strings.Split(text, " ")
	switch parts[0] {
	case "cpy":
		var cpy Copy
		val, err := strconv.Atoi(parts[1])
		if err == nil {
			cpy.Val = val
		} else {
			reg, err := checkRegister(parts[1])
			if err != nil {
				return result, err
			}
			cpy.Src = reg
		}
		{
			reg, err := checkRegister(parts[2])
			if err != nil {
				return result, err
			}
			cpy.Dst = reg
		}
		result.As = cpy
		return result, nil
	case "inc":
		reg, err := checkRegister(parts[1])
		if err != nil {
			return result, err
		}
		var inc Increment
		inc.Reg = reg
		result.As = inc
		return result, nil
	case "dec":
		reg, err := checkRegister(parts[1])
		if err != nil {
			return result, err
		}
		var dec Decrement
		dec.Reg = reg
		result.As = dec
		return result, nil
	case "jnz":
		reg := ""
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			reg, err = checkRegister(parts[1])
			if err != nil {
				return result, err
			}
		}
		jump, err := strconv.Atoi(parts[2])
		if err != nil {
			return result, err
		}
		var jnz JumpNotZero
		jnz.Reg = reg
		jnz.Val = val
		jnz.Jump = jump
		result.As = jnz
		return result, nil
	}
	return result, fmt.Errorf("invalid instruction")
}

func parseContent(content string) ([]Instruction, error) {
	lines := strings.Split(content, "\n")
	result := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		inst, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		result = append(result, inst)
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

type Program struct {
	Registers []int
	PC        int
}

func CharIndex(ch byte) int {
	return int(ch) - int('a')
}

func (p *Program) SetupRegisters(is []Instruction) {
	maxReg := 0
	for _, inst := range is {
		switch i := inst.As.(type) {
		case Copy:
			if len(i.Src) != 0 {
				reg := CharIndex(i.Src[0])
				maxReg = utils.Max(maxReg, reg)
			}
			{
				reg := CharIndex(i.Dst[0])
				maxReg = utils.Max(maxReg, reg)
			}
		case Increment:
			reg := CharIndex(i.Reg[0])
			maxReg = utils.Max(maxReg, reg)
		case Decrement:
			reg := CharIndex(i.Reg[0])
			maxReg = utils.Max(maxReg, reg)
		case JumpNotZero:
			if len(i.Reg) > 0 {
				reg := CharIndex(i.Reg[0])
				maxReg = utils.Max(maxReg, reg)
			}
		}
	}
	p.Registers = make([]int, maxReg+1)
}

func (p *Program) GetRegister(name string) *int {
	if len(name) == 1 {
		ix := CharIndex(name[0])
		if ix < len(p.Registers) {
			return &p.Registers[ix]
		}
	}
	return nil
}

func (p *Program) Execute(is []Instruction) {
	for 0 <= p.PC && p.PC < len(is) {
		inst := is[p.PC]
		switch i := inst.As.(type) {
		case Copy:
			val := i.Val
			if len(i.Src) != 0 {
				val = *p.GetRegister(i.Src)
			}
			dst := p.GetRegister(i.Dst)
			*dst = val
			p.PC++
		case Increment:
			reg := p.GetRegister(i.Reg)
			(*reg)++
			p.PC++
		case Decrement:
			reg := p.GetRegister(i.Reg)
			(*reg)--
			p.PC++
		case JumpNotZero:
			val := i.Val
			if len(i.Reg) > 0 {
				val = *p.GetRegister(i.Reg)
			}
			if val != 0 {
				p.PC += i.Jump
			} else {
				p.PC++
			}
		}
	}
}

func part1(is []Instruction) int {
	var p Program
	p.SetupRegisters(is)
	p.Execute(is)
	return *p.GetRegister("a")
}

func part2(is []Instruction) int {
	var p Program
	p.SetupRegisters(is)
	*p.GetRegister("c") = 1
	p.Execute(is)
	return *p.GetRegister("a")
}

func Solve() {
	input, err := parseFile("./inputs/day-12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 12")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
