package vm

import (
	"fmt"

	"github.com/juanpablocruz/sim8086/pkg/instruction"
)

type ExecResult struct {
	ShiftCount         int
	RepCount           int
	BranchTaken        bool
	AddressIsUnaligned bool
	Unimplemented      bool
}

func (c *Computer8086) ExecInstruction(in instruction.Instruction) ExecResult {
	res := ExecResult{}

	wWidth := 1
	if in.Wide {
		wWidth = 2
	}

	switch in.Op {
	case instruction.Op_mov:
		c.WriteN(in.Reg, in.RM, wWidth)
	default:
		res.Unimplemented = true
	}

	return res
}

func (c *Computer8086) WriteN(reg instruction.InstructionOperand, rm instruction.InstructionOperand, size int) error {
	regIdx := c.ResolveRegister(reg.Register)

	if regIdx.Name == "" {
		return fmt.Errorf("not found register: %s", reg.Name)
	}

	val := c.AccessRegister(rm)

	if regIdx.Wide {
		c.Registers[regIdx.Name] = val.Value
	} else if regIdx.High {
		c.Registers[regIdx.Name] = val.Value << 8
	} else {
		c.Registers[regIdx.Name] = int(int8(val.Value))
	}
	return nil
}

func (c *Computer8086) AccessRegister(reg instruction.InstructionOperand) ComputerRegister {
	switch reg.Type {
	case instruction.Operand_Immediate:
		return ComputerRegister{Value: reg.Value}
	case instruction.Operand_Register:
		r := c.ResolveRegister(reg.Register)
		val := c.Registers[r.Name]
		return ComputerRegister{Value: val}
	case instruction.Operand_Memory:
	}
	return ComputerRegister{}
}

func (c *Computer8086) ResolveRegister(reg instruction.Register) ComputerRegister {
	switch reg.Name {
	case "AX":
		return ComputerRegister{Name: "AX", Wide: true}
	case "AL":
		return ComputerRegister{Name: "AX"}
	case "AH":
		return ComputerRegister{Name: "AX", High: true}
	case "BX":
		return ComputerRegister{Name: "BX", Wide: true}
	case "BL":
		return ComputerRegister{Name: "BX"}
	case "BH":
		return ComputerRegister{Name: "BX", High: true}
	case "CX":
		return ComputerRegister{Name: "CX", Wide: true}
	case "CL":
		return ComputerRegister{Name: "CX"}
	case "CH":
		return ComputerRegister{Name: "CX", High: true}
	case "DX":
		return ComputerRegister{Name: "DX", Wide: true}
	case "DL":
		return ComputerRegister{Name: "DX"}
	case "DH":
		return ComputerRegister{Name: "DX", High: true}
	case "SP":
		return ComputerRegister{Name: "SP"}
	case "BP":
		return ComputerRegister{Name: "BP"}
	case "SI":
		return ComputerRegister{Name: "SI"}
	case "DI":
		return ComputerRegister{Name: "DI"}
	}

	return ComputerRegister{}
}
