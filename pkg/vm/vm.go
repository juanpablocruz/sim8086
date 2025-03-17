package vm

import (
	"bytes"
	"fmt"
	"maps"
	"strings"

	"github.com/juanpablocruz/sim8086/pkg/instruction"
	"github.com/juanpablocruz/sim8086/pkg/options"
)

type ComputerRegister struct {
	Wide  bool
	High  bool
	Name  string
	Value int
}

type Computer8086 struct {
	Registers map[string]int
}

func New() *Computer8086 {
	c := Computer8086{}
	c.Registers = map[string]int{
		"AX": 0,
		"BX": 0,
		"CX": 0,
		"DX": 0,
		"SP": 0,
		"BP": 0,
		"SI": 0,
		"DI": 0,
	}
	return &c
}

func (c *Computer8086) PrintRegisters(out *bytes.Buffer, pad ...int) {
	padding := 0
	if len(pad) == 1 {
		padding = pad[0]
	}
	for r, val := range c.Registers {
		fmt.Fprintf(out, "%s%s: 0x%04x (%d)\n", strings.Repeat(" ", padding), r, val, val)
	}
}

func (c *Computer8086) Exec(in instruction.Instruction, flags uint32) error {
	if in.Op == instruction.Op_None {
		return fmt.Errorf("no operation found")
	}

	prevReg := map[string]int{}
	maps.Copy(prevReg, c.Registers)
	exec := c.ExecInstruction(in)
	if exec.Unimplemented {
		return fmt.Errorf("unimplemented instruction (%s)", instruction.GetMnemonic(in.Op))
	}

	var out bytes.Buffer

	out.WriteString(in.String())

	out.WriteString(" ; ")
	if flags&options.SimFlag_NoRegisterDiffs == options.SimFlag_NoRegisterDiffs {
		out = PrintRegisterDifference(&prevReg, &c.Registers, out)
	}
	out.WriteString("\n")

	fmt.Printf("%s", out.String())

	return nil
}

func PrintRegisterDifference(prev *map[string]int, curr *map[string]int, out bytes.Buffer) bytes.Buffer {
	for i, reg := range *prev {
		newVal := (*curr)[i]

		if reg != newVal {
			out.WriteString(fmt.Sprintf("%s:", strings.ToLower(i)))
			out.WriteString(fmt.Sprintf("0x%x->0x%x", reg, newVal))

			out.WriteString(" ")
		}
	}
	return out
}
