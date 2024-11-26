package day10

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"aoc2016/internal/utils"
)

type GiveTo int

const (
	OUTPUT GiveTo = iota
	BOT
)

type Give struct {
	Bot    int
	LowTo  GiveTo
	Low    int
	HighTo GiveTo
	High   int
}
type Take struct {
	Bot   int
	Value int
}
type Instruction struct {
	Union interface{}
}
type Instructions struct {
	Gives []Give
	Takes []Take
}

func NewInstructionTake(value, bot int) Instruction {
	var i Take
	i.Value = value
	i.Bot = bot
	var inst Instruction
	inst.Union = i
	return inst
}

func NewInstructionGive(bot, low, high int, lowTo, highTo GiveTo) Instruction {
	var i Give
	i.Bot = bot
	i.Low = low
	i.High = high
	i.LowTo = lowTo
	i.HighTo = highTo
	var inst Instruction
	inst.Union = i
	return inst
}

type ParseInstruction struct {
	valueRe *regexp.Regexp
	givenRe *regexp.Regexp
}

func NewParseInstruction() ParseInstruction {
	var p ParseInstruction
	p.valueRe = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	p.givenRe = regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)
	return p
}

func (p *ParseInstruction) GiveToFromString(text string) GiveTo {
	if text == "bot" {
		return BOT
	} else {
		return OUTPUT
	}
}

func (p *ParseInstruction) Parse(text string) (inst Instruction, err error) {
	if p.valueRe.MatchString(text) {
		matches := p.valueRe.FindStringSubmatch(text)
		value, _ := strconv.Atoi(matches[1])
		bot, _ := strconv.Atoi(matches[2])
		return NewInstructionTake(value, bot), nil
	} else if p.givenRe.MatchString(text) {
		matches := p.givenRe.FindStringSubmatch(text)
		bot, _ := strconv.Atoi(matches[1])
		lowTo := p.GiveToFromString(matches[2])
		low, _ := strconv.Atoi(matches[3])
		highTo := p.GiveToFromString(matches[4])
		high, _ := strconv.Atoi(matches[5])
		return NewInstructionGive(bot, low, high, lowTo, highTo), nil
	} else {
		return inst, fmt.Errorf("invalid instruction pattern")
	}
}

func (p *ParseInstruction) ParseLines(content string) (Instructions, error) {
	lines := strings.Split(content, "\n")
	var none Instructions
	var instructions Instructions
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		inst, err := p.Parse(line)
		if err != nil {
			return none, err
		}
		switch v := inst.Union.(type) {
		case Give:
			instructions.Gives = append(instructions.Gives, v)
		case Take:
			instructions.Takes = append(instructions.Takes, v)
		}
	}
	return instructions, nil
}

func parseFile(filename string) (Instructions, error) {
	content, err := utils.ReadAllFile(filename)
	if err != nil {
		var none Instructions
		return none, err
	}
	p := NewParseInstruction()
	return p.ParseLines(content)
}

type Bot struct {
	values []int
}

func (b *Bot) Add(microchip int) {
	ix, _ := slices.BinarySearch(b.values, microchip)
	b.values = slices.Insert(b.values, ix, microchip)
}

func (b *Bot) PopLowHigh() (int, int) {
	low := b.values[0]
	high := b.values[len(b.values)-1]
	b.values = b.values[1 : len(b.values)-1]
	return low, high
}

func minMax(value, min, max int) (int, int) {
	if value < min {
		min = value
	}
	if value > max {
		max = value
	}
	return min, max
}

type Problem struct {
	Gives   []Give
	Takes   []Take
	Outputs []int
	minBot  int
	bots    []Bot
	queue   []int
	currBot int
	ix      int
}

func NewProblem(gives []Give, takes []Take) Problem {
	var p Problem
	p.Gives = gives
	p.Takes = takes
	p.currBot = -1
	{
		slices.SortFunc(p.Gives, func(a, b Give) int {
			return cmp.Compare(a.Bot, b.Bot)
		})
	}
	{
		maxOutput := 1
		for _, v := range p.Gives {
			if v.LowTo == OUTPUT {
				maxOutput = max(maxOutput, v.Low)
			}
			if v.HighTo == OUTPUT {
				maxOutput = max(maxOutput, v.High)
			}
		}
		p.Outputs = make([]int, maxOutput+1)
	}
	{
		minBot := 1 << 32
		maxBot := -(1 << 31)
		for _, v := range p.Takes {
			minBot, maxBot = minMax(v.Bot, minBot, maxBot)
		}
		for _, v := range p.Gives {
			if v.LowTo == BOT {
				minBot, maxBot = minMax(v.Low, minBot, maxBot)
			}
			if v.HighTo == BOT {
				minBot, maxBot = minMax(v.High, minBot, maxBot)
			}
		}
		rangeBot := maxBot - minBot + 1
		p.bots = make([]Bot, rangeBot)
		p.minBot = minBot
		for _, v := range p.Takes {
			if p.AddToBot(v.Bot, v.Value) {
				p.queuePush(v.Bot)
			}
		}
	}
	return p
}

func (p *Problem) GetBot(id int) *Bot {
	return &p.bots[id-p.minBot]
}

func (p *Problem) AddToBot(id, microchip int) bool {
	bot := p.GetBot(id)
	bot.Add(microchip)
	return len(bot.values) == 2
}

func (p *Problem) queuePush(bot int) {
	p.queue = append(p.queue, bot)
}

func (p *Problem) queuePop() int {
	bot := p.queue[0]
	p.queue = p.queue[1:]
	return bot
}

func (p *Problem) findStartGives(bot int) int {
	ix, _ := slices.BinarySearchFunc(p.Gives, bot, func(g Give, bot int) int {
		return cmp.Compare(g.Bot, bot)
	})
	for ix >= 1 && p.Gives[ix-1].Bot == bot {
		ix--
	}
	return ix
}

func (p *Problem) Complete() bool {
	return len(p.queue) == 0 && p.currBot == -1
}

func (p *Problem) Next() (bool, int, int, int) {
	if p.Complete() {
		return false, 0, 0, 0
	}

	if p.currBot == -1 {
		p.currBot = p.queuePop()
		ix := p.findStartGives(p.currBot)
		p.ix = ix
	}

	if p.ix >= 0 {
		if p.ix < len(p.Gives) && p.Gives[p.ix].Bot == p.currBot {
			bot := p.GetBot(p.currBot)
			if len(bot.values) < 2 {
				panic("unreachable len(bot.values) < 2")
			}
			low, high := bot.PopLowHigh()
			v := p.Gives[p.ix]
			{
				if v.LowTo == BOT {
					if p.AddToBot(v.Low, low) {
						p.queuePush(v.Low)
					}
				} else {
					p.Outputs[v.Low] = low
				}
			}
			{
				if v.HighTo == BOT {
					if p.AddToBot(v.High, high) {
						p.queuePush(v.High)
					}
				} else {
					p.Outputs[v.High] = high
				}
			}
			p.ix++
			return true, p.currBot, low, high
		}
	}
	p.currBot = -1
	return p.Next()
}

func findBot(is Instructions, needleLow, needleHigh int) int {
	p := NewProblem(is.Gives, is.Takes)
	for {
		if found, bot, low, high := p.Next(); found {
			if low == needleLow && high == needleHigh {
				return bot
			}
			continue
		}
		break
	}
	return 0
}

func part1(is Instructions) int {
	return findBot(is, 17, 61)
}

func multiplyFirstsOutputs(is Instructions) int {
	p := NewProblem(is.Gives, is.Takes)
	for {
		if found, _, _, _ := p.Next(); found {
			continue
		}
		break
	}
	return p.Outputs[0] * p.Outputs[1] * p.Outputs[2]
}

func part2(is Instructions) int {
	return multiplyFirstsOutputs(is)
}

func Solve() {
	input, err := parseFile("./inputs/day-10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 10")
	utils.TimeIt("Part 1:", "%d", func() any { return part1(input) })
	utils.TimeIt("Part 2:", "%d", func() any { return part2(input) })
}
