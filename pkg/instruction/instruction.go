package instruction

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/juanpablocruz/aware/pkg/token"
)

type OpCode uint8

const (
	NONE OpCode = 0
	MOV  OpCode = (1 << 1) | (1<<7)>>2
)

// 10001001 11011001

type Mode byte

const (
	Memory   Mode = 0b00
	Displ8   Mode = 0b01
	Disply16 Mode = 0b10
	Reg      Mode = 0b11
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

type Register struct {
	Name string
	Code byte
}

func (r Register) String() string {
	return strings.ToLower(r.Name)
}

func (i Instruction) String() string {
	var out bytes.Buffer

	out.WriteString(i.Op.String() + " ")
	if i.Direction {

		out.WriteString(i.Reg.String() + ",")
		out.WriteString(i.RM.String())
	} else {

		out.WriteString(i.RM.String() + ",")
		out.WriteString(i.Reg.String())
	}
	return out.String()
}

type Instruction struct {
	OpCode    token.TokenType
	Direction bool
	Wide      bool
	Mode      Mode
	Reg       Register
	RM        Register

	Op OperationType
}
