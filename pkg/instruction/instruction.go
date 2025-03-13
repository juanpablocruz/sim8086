package instruction

import (
	"bytes"
	"fmt"
	"strings"
)

type OperandType int

const (
	Operand_None OperandType = iota
	Operand_Register
	Operand_Memory
	Operand_Immediate
)

type OpCode uint8

const (
	NONE OpCode = 0
	MOV  OpCode = (1 << 1) | (1<<7)>>2
)

// 10001001 11011001

type Mode byte

const (
	// Memory Mode, no displacement follows (except R/M=110, then 16-bit disp)
	Memory Mode = 0b00
	// Memory Mode, 8-bit displacement
	Displ8 Mode = 0b01
	// Memory Mode, 16-bit displacement
	Disply16 Mode = 0b10
	// Register Mode, no displacement
	Reg Mode = 0b11
)

func (m Mode) String() string {
	switch m {
	case Memory:
		return fmt.Sprintf("Memory (%08b)", Memory)
	case Displ8:
		return fmt.Sprintf("Disply8 (%08b)", Displ8)
	case Disply16:
		return fmt.Sprintf("Disply16 (%08b)", Disply16)
	case Reg:
		return fmt.Sprintf("Reg (%08b)", Reg)
	default:
		return "Memory"
	}
}

type EffectiveAddressExpression struct {
	Terms           [2]Register
	ExplicitSegment int
	Displacement    int
	Flags           int
}

func (eae EffectiveAddressExpression) String() string {
	var out bytes.Buffer

	out.WriteString(strings.ToLower(eae.Terms[0].Name))
	if eae.Terms[1].Code > 0 {
		out.WriteString(fmt.Sprintf(" + %s", strings.ToLower(eae.Terms[1].Name)))
	}
	if eae.Displacement != 0 {
		switch eae.Terms[0].Code {
		case 5, 6, 7, 3:
		default:
			out.WriteString(fmt.Sprintf(" + %d", eae.Displacement))
		}
	}
	return fmt.Sprintf("[%s]", out.String())
}

type Immediate struct {
	Value int
	Flags int
}

type Register struct {
	Name string
	Code byte
}

func (r InstructionOperand) String() string {
	if r.Type == Operand_Register {
		return strings.ToLower(r.Name)
	}
	if r.Type == Operand_Immediate {
		return fmt.Sprintf("%d", r.Value)
	}
	if r.Type == Operand_Memory {
		return r.EffectiveAddressExpression.String()
	}
	return fmt.Sprintf("ERROR(REG): %v %v %v", r.Type, r.Name, r.Code)
}

func (i Instruction) String() string {
	var out bytes.Buffer

	out.WriteString(i.Op.String() + " ")
	if i.RM.Type != Operand_Register {
		out.WriteString(i.Reg.String() + ",")
		out.WriteString(i.RM.String())
	} else if i.Direction {
		out.WriteString(i.Reg.String() + ",")
		out.WriteString(i.RM.String())
	} else {
		out.WriteString(i.RM.String() + ",")
		out.WriteString(i.Reg.String())
	}
	return out.String()
}

type InstructionOperand struct {
	Type OperandType
	Immediate
	EffectiveAddressExpression
	Register
}

type Instruction struct {
	Direction bool
	Wide      bool
	Mode      Mode
	Reg       InstructionOperand
	RM        InstructionOperand

	Op OperationType
}
