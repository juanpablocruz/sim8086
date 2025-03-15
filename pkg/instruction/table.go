package instruction

import (
	"bytes"
	"fmt"

	"github.com/juanpablocruz/sim8086/pkg/reader"
)

func New8086InstructionTable() InstructionTable {
	return InstructionTable{
		EncodingCount:           len(InstructionTable8086),
		Encodings:               InstructionTable8086,
		MaxInstructionByteCount: 15,
	}
}

func (it *InstructionTable) DecodeInstruction(r *reader.Reader) (Instruction, error) {
	instr := Instruction{}

	startingAddress := r.SegmentOffset
	startingByte := r.Curr

	for _, encoding := range it.Encodings {
		in, err := it.TryDecode(encoding, r)
		if err != nil {
			return instr, err
		}
		instr = in

		if OperationType(instr.Op) != Op_None {
			break
		} else {
			r.SegmentOffset = startingAddress
			r.Curr = startingByte
		}
	}
	return instr, nil
}

func (it *InstructionTable) TryDecode(encoding InstructionEncoding, r *reader.Reader) (Instruction, error) {
	// fmt.Printf("TryDecode: %08b\n", r.Curr)
	instr := Instruction{}

	bitIndx := 0

	bits := make([]byte, Bits_Count)
	has := make([]bool, Bits_Count)
	valid := true
	for _, bit := range encoding.Bits {
		if bit.Usage == Bits_End {
			break
		}
		if bit.Usage == Bits_Literal {
			if bit.Value == 0 {
				bitIndx += int(bit.BitCount)
				continue
			}
			masked := r.Curr >> (8 - bit.BitCount)

			if bit.Value&masked == masked {
				instr.Op = (encoding.Op)
				bitIndx += int(bit.BitCount)
				valid = true
			} else {
				valid = false
			}
		} else {
			// we have already parsed bitIndx bits.
			// example 100010, so we want to test the next
			// bit.BitCount
			//
			// 10001011
			//       ^
			mask := 0
			for range bit.BitCount {
				if bitIndx >= 8 {
					bitIndx -= 8
					r.ReadByte()
				}
				mask |= 1 << (8 - bitIndx - 1)
				bitIndx++
			}

			// fmt.Printf("%d - %08b\n", bit.Usage, r.Curr)
			bits[bit.Usage] |= (r.Curr & byte(mask)) >> (8 - byte(bitIndx))
			has[bit.Usage] = true

		}
	}
	if !valid {
		return Instruction{}, nil
	}

	debugBits(bits)
	mod := bits[Bits_MOD]
	rm := bits[Bits_RM]
	w := bits[Bits_W] == 1
	s := bits[Bits_S] == 1
	d := bits[Bits_D] == 1

	hasDirectAddr := (mod == 0b00) && (rm == 0b110)
	has[Bits_Disp] = ((has[Bits_Disp]) || (mod == 0b10) || hasDirectAddr)

	var regOperand InstructionOperand
	if has[Bits_REG] {
		regOperand, _ = it.ResolveRegister(bits[Bits_REG], w)
	}

	var modOperand InstructionOperand
	if has[Bits_MOD] {
		// register mode, no displacement
		if mod == byte(Reg) {
			modOperand, _ = it.ResolveRegister(rm, w || (bits[Bits_RMRegAlwaysW] == 1))
		} else {
			// Memory mode
			mem, _ := it.ResolveMemoryAddress(Mode(mod), bits[Bits_RM])
			tmp := mem.DisplacementValue
			if hasDirectAddr {
				c, _ := r.ReadByte()
				tmp = int(c)
				if w {
					c, _ := r.ReadByte()
					tmp |= int(c) << 8
				}
			}
			for dis := 0; dis < mem.Displacement; dis += 8 {
				c, _ := r.ReadByte()
				tmp |= (int(c) << dis)
			}
			switch mem.Displacement {
			case 8:
				tmp = int(int8(tmp))
			case 16:
				tmp = int(int16(tmp))
			}
			mem.DisplacementValue = tmp
			modOperand = mem
		}
	}

	// fmt.Printf("data: %v disp: %v mod: %v\n", has[Bits_Data], has[Bits_Disp], has[Bits_MOD])
	if has[Bits_Data] && has[Bits_Disp] && !has[Bits_MOD] {
		fmt.Printf("Entro\n")
	} else {
		if has[Bits_Data] {
			data := it.ParseDataValue(r, has[Bits_Data], w, s)
			r.PrintInstruction()
			flags := int(0)
			if bits[Bits_W] == 1 {
				flags |= int(Bits_W)
			}

			imm, _ := it.ResolveImmediate(data, flags)

			// If we have already modOperand, then we are moving a literal to a EA,
			// therefore, the literal goes to reg
			if modOperand.Type == Operand_None {
				modOperand = imm
			} else {
				regOperand = imm
			}
		}
	}

	switch modOperand.Type {
	case Operand_Immediate:
		if d {
			instr.RM = regOperand
			instr.Reg = modOperand
		} else {
			instr.Reg = regOperand
			instr.RM = modOperand
		}
	case Operand_Memory, Operand_Register:
		if !d {
			instr.RM = regOperand
			instr.Reg = modOperand
		} else {
			instr.Reg = regOperand
			instr.RM = modOperand
		}
	}

	instr.Mode = Mode(mod)
	instr.Direction = d
	instr.Wide = w

	r.EndInstructionAndPrint()
	fmt.Printf("%v\n", instr)
	fmt.Println("")
	fmt.Println("")
	return instr, nil
}

func (it *InstructionTable) ParseDataValue(r *reader.Reader, exists, wide, signedExtended bool) int {
	if !exists {
		return 0
	}
	if wide {
		dataL, _ := r.ReadByte()
		dataH, _ := r.ReadByte()

		valInt := (int(dataH) << 8) | int(dataL)
		return int(int16(valInt))
	} else {
		dataL, _ := r.ReadByte()

		if signedExtended {
			return int(int8(dataL))
		}

		return int(dataL)
	}
}

func debugBits(bits []byte) {
	var out bytes.Buffer
	for t, bit := range bits {
		switch t {
		case int(Bits_Count):
			out.WriteString("Bits_Count: ")
		case int(Bits_D):
			out.WriteString("Bits_D: ")
		case int(Bits_W):
			out.WriteString("Bits_W: ")
		case int(Bits_S):
			out.WriteString("Bits_S: ")
		case int(Bits_RelJMPDisp):
			out.WriteString("Bits_RelJMPDisp: ")
		case int(Bits_WMakesDataW):
			out.WriteString("Bits_WMakesDataW: ")
		case int(Bits_SR):
			out.WriteString("Bits_SR: ")
		case int(Bits_Z):
			out.WriteString("Bits_Z: ")
		case int(Bits_End):
			out.WriteString("Bits_End: ")
		case int(Bits_V):
			out.WriteString("Bits_V: ")
		case int(Bits_Data):
			out.WriteString("Bits_Data: ")
		case int(Bits_Disp):
			out.WriteString("Bits_Disp: ")
		case int(Bits_DispAlwaysW):
			out.WriteString("Bits_DispAlwaysW: ")
		case int(Bits_Far):
			out.WriteString("Bits_Far: ")
		case int(Bits_Literal):
			out.WriteString("Bits_Literal: ")
		case int(Bits_MOD):
			out.WriteString("Bits_MOD: ")
		case int(Bits_REG):
			out.WriteString("Bits_REG: ")
		case int(Bits_RM):
			out.WriteString("Bits_RM: ")
		case int(Bits_RMRegAlwaysW):
			out.WriteString("Bits_RMRegAlwaysW: ")
		default:
			out.WriteString(fmt.Sprintf("[%d]: ", t))
		}
		out.WriteString(fmt.Sprintf("%08b (%d)\n", bit, bit))
	}
	fmt.Println(out.String())
}
