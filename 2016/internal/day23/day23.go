package day23

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

type Instruction interface {
	Toggle() Instruction
	String() string
}
type ValReg interface {
	String() string
}

type Val int
type Reg byte

func (val Val) String() string {
	return strconv.Itoa(int(val))
}
func (reg Reg) String() string {
	return string([]byte{byte(reg)})
}

type Cpy struct {
	a ValReg
	b ValReg
}
type Inc struct {
	a ValReg
}
type Dec struct {
	a ValReg
}
type Jnz struct {
	a ValReg
	b ValReg
}
type Tgl struct {
	a ValReg
}

func (inst Cpy) String() string {
	return "cpy " + inst.a.String() + " " + inst.b.String()
}
func (inst Inc) String() string {
	return "inc " + inst.a.String()
}
func (inst Dec) String() string {
	return "dec " + inst.a.String()
}
func (inst Jnz) String() string {
	return "jnz " + inst.a.String() + " " + inst.b.String()
}
func (inst Tgl) String() string {
	return "tgl " + inst.a.String()
}

func (inst Cpy) Toggle() Instruction {
	return Jnz{a: inst.a, b: inst.b}
}
func (inst Inc) Toggle() Instruction {
	return Dec{a: inst.a}
}
func (inst Dec) Toggle() Instruction {
	return Inc{a: inst.a}
}
func (inst Jnz) Toggle() Instruction {
	return Cpy{a: inst.a, b: inst.b}
}
func (inst Tgl) Toggle() Instruction {
	return Inc{a: inst.a}
}

func parseRegister(text string) (byte, error) {
	if len(text) == 1 {
		if int('a') <= int(text[0]) && int(text[0]) <= int('z') {
			return text[0], nil
		}
	}
	return byte(0), fmt.Errorf("invalid register")
}

func parseValReg(text string) (ValReg, error) {
	reg, err := parseRegister(text)
	if err == nil {
		return Reg(reg), nil
	}
	val, err := strconv.Atoi(text)
	if err == nil {
		return Val(val), nil
	}
	return Reg('?'), fmt.Errorf("invalid value or register '%s'", text)
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
		// cpy a b
		if rest, found := strings.CutPrefix(line, "cpy "); found {
			first, second, found := strings.Cut(rest, " ")
			if !found {
				return nil, fmt.Errorf("invalid instruction '%s'", line)
			}
			a, err := parseValReg(first)
			if err != nil {
				return nil, err
			}
			b, err := parseValReg(second)
			if err != nil {
				return nil, err
			}
			cpy := Cpy{a: a, b: b}
			result = append(result, cpy)
			continue
		}
		// int a
		if rest, found := strings.CutPrefix(line, "inc "); found {
			a, err := parseValReg(rest)
			if err != nil {
				return nil, err
			}
			inc := Inc{a: a}
			result = append(result, inc)
			continue
		}
		// dec a
		if rest, found := strings.CutPrefix(line, "dec "); found {
			a, err := parseValReg(rest)
			if err != nil {
				return nil, err
			}
			dec := Dec{a: a}
			result = append(result, dec)
			continue
		}
		// jnz a b
		if rest, found := strings.CutPrefix(line, "jnz "); found {
			first, second, found := strings.Cut(rest, " ")
			if !found {
				return nil, fmt.Errorf("invalid instruction '%s'", line)
			}
			a, err := parseValReg(first)
			if err != nil {
				return nil, err
			}
			b, err := parseValReg(second)
			if err != nil {
				return nil, err
			}
			jnz := Jnz{a: a, b: b}
			result = append(result, jnz)
			continue
		}
		// tgl a
		if rest, found := strings.CutPrefix(line, "tgl "); found {
			a, err := parseValReg(rest)
			if err != nil {
				return nil, err
			}
			tgl := Tgl{a: a}
			result = append(result, tgl)
			continue
		}
		return nil, fmt.Errorf("unknown instruction '%s'", line)
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
	registers []int
	pc        int
}

func NewProgram(registerCount int) Program {
	registers := make([]int, registerCount)
	return Program{registers: registers, pc: 0}
}

func (p Program) GetRegister(reg byte) *int {
	ix := int(reg) - int('a')
	if 0 <= ix && ix <= len(p.registers) {
		return &p.registers[ix]
	}
	return nil
}

func (p Program) getRegister(valReg ValReg) *int {
	switch reg := valReg.(type) {
	case Reg:
		return p.GetRegister(byte(reg))
	}
	return nil
}

func (p Program) getValue(valReg ValReg) (int, bool) {
	switch v := valReg.(type) {
	case Reg:
		reg := p.GetRegister(byte(v))
		if reg != nil {
			return *reg, true
		}
	case Val:
		return int(v), true
	}
	return 0, false
}

func (p Program) Run(isBase []Instruction) {
	is := make([]Instruction, len(isBase))
	copy(is, isBase)
	for 0 <= p.pc && p.pc < len(is) {
		inst := is[p.pc]
		switch t := inst.(type) {
		case Inc:
			reg := p.getRegister(t.a)
			if reg != nil {
				*reg++
			}
			p.pc++
		case Dec:
			reg := p.getRegister(t.a)
			if reg != nil {
				*reg--
			}
			p.pc++
		case Jnz:
			jump := 1
			a, foundA := p.getValue(t.a)
			b, foundB := p.getValue(t.b)
			if foundA && foundB && a != 0 {
				jump = b
			}
			p.pc += jump
		case Cpy:
			val, found := p.getValue(t.a)
			reg := p.getRegister(t.b)
			if found && reg != nil {
				*reg = val
			}
			p.pc++
		case Tgl:
			val, found := p.getValue(t.a)
			if found {
				ix := p.pc + val
				if 0 <= ix && ix < len(is) {
					is[ix] = is[ix].Toggle()
				}
			}
			p.pc++
		}
	}
}

func getReg(valReg ValReg) (byte, bool) {
	switch v := valReg.(type) {
	case Reg:
		return byte(v), true
	}
	return '?', false
}

func GetRegistersCount(is []Instruction) int {
	maxReg := 0
	for _, inst := range is {
		switch t := inst.(type) {
		case Inc:
			if v, found := getReg(t.a); found {
				maxReg = utils.Max(maxReg, int(v))
			}
		case Dec:
			if v, found := getReg(t.a); found {
				maxReg = utils.Max(maxReg, int(v))
			}
		case Jnz:
			if v, found := getReg(t.a); found {
				maxReg = utils.Max(maxReg, int(v))
			}
			if v, found := getReg(t.b); found {
				maxReg = utils.Max(maxReg, int(v))
			}
		case Cpy:
			if v, found := getReg(t.a); found {
				maxReg = utils.Max(maxReg, int(v))
			}
			if v, found := getReg(t.b); found {
				maxReg = utils.Max(maxReg, int(v))
			}
		case Tgl:
			if v, found := getReg(t.a); found {
				maxReg = utils.Max(maxReg, int(v))
			}
		}
	}
	return maxReg - int('a') + 1
}

func part1(is []Instruction) int {
	registersCount := GetRegistersCount(is)
	p := NewProgram(registersCount)
	a := p.GetRegister('a')
	if a == nil {
		return -1
	}
	*a = 7
	p.Run(is)
	return *p.GetRegister('a')
}

func part2(is []Instruction) int {
	registersCount := GetRegistersCount(is)
	p := NewProgram(registersCount)
	a := p.GetRegister('a')
	if a == nil {
		return -1
	}
	*a = 12
	p.Run(is)
	return *a
}

func Solve() {
	input, err := parseFile("./inputs/day-23.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 23")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
