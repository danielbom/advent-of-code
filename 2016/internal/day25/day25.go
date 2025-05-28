package day25

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

const REVERSED_LOGIC = true
const STANDARD_CODE = true

type Instruction interface {
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

func getReg(valReg ValReg) (byte, bool) {
	switch v := valReg.(type) {
	case Reg:
		return byte(v), true
	}
	return '?', false
}
func getVal(valReg ValReg) (int, bool) {
	switch v := valReg.(type) {
	case Val:
		return int(v), true
	}
	return -1, false
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
type Out struct {
	a ValReg
}

// OBS: extra instruction types to optimize the input instructions
// Look at the input/day-25.txt
type Nop_ struct{}
type Add_ struct {
	a ValReg
	b ValReg
}
type Mul_ struct {
	a ValReg
	b ValReg
}
type Jmp_ struct {
	a ValReg
}
type Div_ struct {
	a ValReg
	b ValReg
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
func (inst Out) String() string {
	return "out " + inst.a.String()
}
func (inst Nop_) String() string {
	return "nop"
}
func (inst Add_) String() string {
	return "add " + inst.a.String() + " " + inst.b.String()
}
func (inst Mul_) String() string {
	return "mul " + inst.a.String() + " " + inst.b.String()
}
func (inst Jmp_) String() string {
	return "jmp " + inst.a.String()
}
func (inst Div_) String() string {
	return "div " + inst.a.String()
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

func parseBiop(line, prefix string) (hasPrefix bool, a ValReg, b ValReg, err error) {
	if rest, found := strings.CutPrefix(line, prefix); found {
		first, second, found := strings.Cut(rest, " ")
		if !found {
			err = fmt.Errorf("invalid instruction '%s'", line)
			return true, a, b, err
		}
		a, err = parseValReg(first)
		if err != nil {
			return true, a, b, err
		}
		b, err = parseValReg(second)
		if err != nil {
			return true, a, b, err
		}
		return true, a, b, nil
	}
	return false, a, b, nil
}

func parseUnop(line, prefix string) (hasPrefix bool, a ValReg, err error) {
	if rest, found := strings.CutPrefix(line, prefix); found {
		a, err = parseValReg(rest)
		if err != nil {
			return true, a, err
		}
		return true, a, nil
	}
	return false, a, nil
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
		// comments
		if strings.HasPrefix(line, "--") {
			continue
		}
		// nop
		if _, found := strings.CutPrefix(line, "nop"); found {
			nop := Nop_{}
			result = append(result, nop)
			continue
		}
		// add a b
		if hasPrefix, a, b, err := parseBiop(line, "add "); hasPrefix {
			if err != nil {
				return nil, err
			}
			add := Add_{a: a, b: b}
			result = append(result, add)
			continue
		}
		// mul a b
		if hasPrefix, a, b, err := parseBiop(line, "mul "); hasPrefix {
			if err != nil {
				return nil, err
			}
			mul := Mul_{a: a, b: b}
			result = append(result, mul)
			continue
		}
		// div a b
		if hasPrefix, a, b, err := parseBiop(line, "div "); hasPrefix {
			if err != nil {
				return nil, err
			}
			div := Div_{a: a, b: b}
			result = append(result, div)
			continue
		}
		// jmp a
		if hasPrefix, a, err := parseUnop(line, "jmp "); hasPrefix {
			if err != nil {
				return nil, err
			}
			jmp := Jmp_{a: a}
			result = append(result, jmp)
			continue
		}
		// cpy a b
		if hasPrefix, a, b, err := parseBiop(line, "cpy "); hasPrefix {
			if err != nil {
				return nil, err
			}
			cpy := Cpy{a: a, b: b}
			result = append(result, cpy)
			continue
		}
		// inc a
		if hasPrefix, a, err := parseUnop(line, "inc "); hasPrefix {
			if err != nil {
				return nil, err
			}
			inc := Inc{a: a}
			result = append(result, inc)
			continue
		}
		// dec a
		if hasPrefix, a, err := parseUnop(line, "dec "); hasPrefix {
			if err != nil {
				return nil, err
			}
			dec := Dec{a: a}
			result = append(result, dec)
			continue
		}
		// jnz a b
		if hasPrefix, a, b, err := parseBiop(line, "jnz "); hasPrefix {
			if err != nil {
				return nil, err
			}
			jnz := Jnz{a: a, b: b}
			result = append(result, jnz)
			continue
		}
		// out a
		if hasPrefix, a, err := parseUnop(line, "out "); hasPrefix {
			if err != nil {
				return nil, err
			}
			out := Out{a: a}
			result = append(result, out)
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
	pc        int
	registers []int
}

func NewProgram(registerCount int) Program {
	registers := make([]int, registerCount)
	return Program{registers: registers, pc: 0}
}

func (p *Program) Restart() {
	p.pc = 0
	for i, _ := range p.registers {
		p.registers[i] = 0
	}
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
		case Out:
			if v, found := getReg(t.a); found {
				maxReg = utils.Max(maxReg, int(v))
			}
		}
	}
	return maxReg - int('a') + 1
}

func (p Program) inBound(is []Instruction) bool {
	return 0 <= p.pc && p.pc < len(is)
}

type Output interface {
	Write(v int)
	Valid() bool
}

func (p *Program) Next(is []Instruction, out Output) bool {
	if p.inBound(is) {
		inst := is[p.pc]
		switch t := inst.(type) {
		case Nop_:
			p.pc++
		case Add_:
			regA := p.getRegister(t.a)
			b, foundB := p.getValue(t.b)
			if regA != nil && foundB {
				*regA += b
			}
			p.pc++
		case Mul_:
			regA := p.getRegister(t.a)
			regB := p.getRegister(t.b)
			if regA != nil && regB != nil {
				*regA *= *regB
			}
			p.pc++
		case Div_:
			regA := p.getRegister(t.a)
			regB := p.getRegister(t.b)
			if regA != nil && regB != nil {
				a := *regA
				b := *regB
				*regA = a / b
				*regB = a % b
			}
			p.pc++
		case Jmp_:
			jump := 1
			val, found := p.getValue(t.a)
			if found {
				jump = val
			}
			p.pc += jump
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
		case Out:
			val, found := p.getValue(t.a)
			if found {
				out.Write(val)
			}
			p.pc++
		}
	}
	return p.inBound(is)
}

func (p *Program) Run(is []Instruction, out Output) {
	for p.Next(is, out) && out.Valid() {
		continue
	}
}

type Pattern01s struct {
	current int
	ok      bool
}

func NewPattern01s() Pattern01s {
	return Pattern01s{current: 0, ok: true}
}

func (p *Pattern01s) Write(v int) {
	if p.ok {
		if v == p.current {
			p.current = (p.current + 1) % 2
		} else {
			p.ok = false
		}
	}
}

func (p *Pattern01s) Valid() bool {
	return p.ok
}

func part1OptimizedCode(is []Instruction) int {
	// if you change your input with an optimized version of the instructions
	// you could run this algorithm to find the pattern
	// - ./inputs/day-25-optimized.txt
	// - ./inputs/day-25-optimized-commented.txt
	registersCount := GetRegistersCount(is)
	p := NewProgram(registersCount)
	a := p.GetRegister('a')
	if a == nil {
		return -1
	}
	for i := 0; ; i++ {
		pattern := NewPattern01s()
		p.Restart()
		*a = i
		p.Run(is, &pattern)
		if pattern.ok {
			return i
		}
	}
	return -1 // unreachable
}

func part1StandardCode(is []Instruction) int {
	// if you not change your input, this algorithm could find the answer
	// OBS: assuming that 'd = 'b * 'c
	// - ./inputs/day-25.txt
	d := 0
	// run until initialize 'b and 'c
	{
		pattern := NewPattern01s()
		registersCount := GetRegistersCount(is)
		p := NewProgram(registersCount)
		initB := false
		initC := false
		for {
			inst := is[p.pc]
			switch v := inst.(type) {
			case Cpy:
				switch reg := v.b.(type) {
				case Reg:
					if reg == 'b' {
						initB = true
					}
					if reg == 'c' {
						initC = true
					}
				}
			}
			if !p.Next(is, &pattern) {
				break
			}
			if initB && initC {
				break
			}
		}
		regB := p.GetRegister('b')
		regC := p.GetRegister('c')
		regD := p.GetRegister('d')
		if regB == nil || regC == nil || regD == nil {
			return -1
		}
		d = *regB * *regC
	}
	if REVERSED_LOGIC {
		// reversed logic to find the minimum value of 'a that computes 010101...
		i := 1
		for i <= d {
			if i%2 == 1 {
				i = (i * 2) + 0
			} else {
				i = (i * 2) + 1
			}
		}
		return i - d
	} else {
		// equivalent implementation of the assembunny (breaking the infinity loop with an ending guards)
		for i := 0; ; i++ {
			correct := true
			one := false
			b := 0
			a := d + i
			for a > 0 && correct {
				b = a % 2
				a = a / 2
				// guard
				if one {
					correct = b%2 == 1
				} else {
					correct = b%2 == 0
				}
				one = !one
			}
			if correct {
				return i
			}
		}
		panic("unreachable")
		return -1
	}
}

func part1(is []Instruction) int {
	if STANDARD_CODE {
		return part1StandardCode(is)
	} else {
		return part1OptimizedCode(is)
	}
}

func part2(is []Instruction) int {
	return 0
}

func parseInput() ([]Instruction, error) {
	if STANDARD_CODE {
		return parseFile("./inputs/day-25.txt")
	} else {
		return parseFile("./inputs/day-25-optimized.txt")
	}
}

func Solve() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 25")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
