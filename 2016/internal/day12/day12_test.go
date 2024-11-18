package day12

import (
	"slices"
	"testing"
)

func TestProgramExecute(t *testing.T) {
	content1 := `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`
	tests := []struct {
		content   string
		registers []int
	}{
		{content: content1, registers: []int{42}},
	}

	for _, test := range tests {
		is, err := parseContent(test.content)
		if err != nil {
			t.Errorf("parseContent(%v) fail prematurely", test.content)
			continue
		}
		var p Program
		p.SetupRegisters(is)
		p.Execute(is)
		if slices.Compare(p.Registers, test.registers) != 0 {
			t.Errorf("Program.Registers = %v; want %v", p.Registers, test.registers)
		}
	}
}
