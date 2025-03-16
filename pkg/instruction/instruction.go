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
	Displ16 Mode = 0b10
	// Register Mode, no displacement
	Reg Mode = 0b11
)

func (m Mode) String() string {
	switch m {
	case Memory:
		return fmt.Sprintf("Memory (%08b)", Memory)
	case Displ8:
		return fmt.Sprintf("Disply8 (%08b)", Displ8)
	case Displ16:
		return fmt.Sprintf("Disply16 (%08b)", Displ16)
	case Reg:
		return fmt.Sprintf("Reg (%08b)", Reg)
	default:
		return "Memory"
	}
}

type EffectiveAddressExpression struct {
	Terms             [2]Register
	DisplacementValue int
	Displacement      int
	Flags             int
}

func (eae EffectiveAddressExpression) String() string {
	var out bytes.Buffer

	if eae.Terms[0].Name == "" && eae.Terms[1].Name == "" && eae.DisplacementValue != 0 {
		return fmt.Sprintf("[%d]", eae.DisplacementValue)
	}

	out.WriteString(strings.ToLower(eae.Terms[0].Name))
	if eae.Terms[1].Code > 0 {
		out.WriteString(fmt.Sprintf(" + %s", strings.ToLower(eae.Terms[1].Name)))
	}
	if eae.DisplacementValue != 0 {
		if eae.DisplacementValue < 0 {
			out.WriteString(fmt.Sprintf(" - %d", eae.DisplacementValue*-1))
		} else {
			out.WriteString(fmt.Sprintf(" + %d", eae.DisplacementValue))
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
	if i.IsArithmetic() && i.Reg.Type == Operand_Memory && i.RM.Type == Operand_Immediate {
		if i.Wide {
			out.WriteString("word ")
		} else {
			out.WriteString("byte ")
		}
	}
	out.WriteString(i.Reg.String() + ", ")
	if !i.IsArithmetic() && i.Reg.Type == Operand_Memory && i.RM.Type == Operand_Immediate {
		if i.Wide {
			out.WriteString("word ")
		} else {
			out.WriteString("byte ")
		}
	}
	out.WriteString(i.RM.String())

	return out.String()
}

func (i Instruction) IsArithmetic() bool {
	switch i.Op {
	case Op_add, Op_sub:
		return true
	default:
		return false
	}
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

type InstructionBitsUsage byte

const (
	Bits_End InstructionBitsUsage = iota
	Bits_Literal
	Bits_D
	Bits_S
	Bits_W
	Bits_V
	Bits_Z
	Bits_MOD
	Bits_REG
	Bits_RM
	Bits_SR
	Bits_Disp
	Bits_Data

	Bits_DispAlwaysW
	Bits_WMakesDataW
	Bits_RMRegAlwaysW
	Bits_RelJMPDisp
	Bits_Far

	Bits_Count
)

type InstructionBits struct {
	Usage    InstructionBitsUsage
	BitCount byte
	Shift    byte
	Value    byte
}

type InstructionEncoding struct {
	Op   OperationType
	Bits []InstructionBits
}
