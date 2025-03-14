package instruction

func (it InstructionTable) ResolveRegister(b byte, w bool) (InstructionOperand, bool) {
	regs := map[bool]map[uint8]InstructionOperand{
		true: {
			0: {Register: Register{Name: "AX", Code: 0}, Type: Operand_Register},
			1: {Register: Register{Name: "CX", Code: 1}, Type: Operand_Register},
			2: {Register: Register{Name: "DX", Code: 2}, Type: Operand_Register},
			3: {Register: Register{Name: "BX", Code: 3}, Type: Operand_Register},
			4: {Register: Register{Name: "SP", Code: 4}, Type: Operand_Register},
			5: {Register: Register{Name: "BP", Code: 5}, Type: Operand_Register},
			6: {Register: Register{Name: "SI", Code: 6}, Type: Operand_Register},
			7: {Register: Register{Name: "DI", Code: 7}, Type: Operand_Register},
		},
		false: {
			0: {Register: Register{Name: "AL", Code: 0}, Type: Operand_Register},
			1: {Register: Register{Name: "CL", Code: 1}, Type: Operand_Register},
			2: {Register: Register{Name: "DL", Code: 2}, Type: Operand_Register},
			3: {Register: Register{Name: "BL", Code: 3}, Type: Operand_Register},
			4: {Register: Register{Name: "AH", Code: 4}, Type: Operand_Register},
			5: {Register: Register{Name: "CH", Code: 5}, Type: Operand_Register},
			6: {Register: Register{Name: "DH", Code: 6}, Type: Operand_Register},
			7: {Register: Register{Name: "BH", Code: 7}, Type: Operand_Register},
		},
	}

	subSet := regs[w]

	reg, ok := subSet[b]
	return reg, ok
}

func (it InstructionTable) ResolveMemoryAddress(mod Mode, rm byte) (InstructionOperand, bool) {
	memTables := map[Mode]map[byte]InstructionOperand{
		Memory: {
			0b000: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
						{Name: "SI", Code: 6},
					},
				},
			},
			0b001: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
						{Name: "DI", Code: 7},
					},
				},
			},
			0b010: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
						{Name: "SI", Code: 6},
					},
				},
			},
			0b011: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
						{Name: "DI", Code: 7},
					},
				},
			},
			0b100: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
					Terms: [2]Register{
						{Name: "SI", Code: 6},
					},
				},
			},
			0b101: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
					Terms: [2]Register{
						{Name: "DI", Code: 7},
					},
				},
			},
			0b110: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
				},
			},
			0b111: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 0,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
					},
				},
			},
		},
		Displ8: {
			0b000: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
						{Name: "SI", Code: 6},
					},
				},
			},
			0b001: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
						{Name: "DI", Code: 7},
					},
				},
			},
			0b010: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
						{Name: "SI", Code: 6},
					},
				},
			},
			0b011: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
						{Name: "DI", Code: 7},
					},
				},
			},
			0b100: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "SI", Code: 6},
					},
				},
			},
			0b101: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "DI", Code: 7},
					},
				},
			},
			0b110: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
					},
				},
			},
			0b111: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 8,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
					},
				},
			},
		},

		Displ16: {
			0b000: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
						{Name: "SI", Code: 6},
					},
				},
			},
			0b001: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
						{Name: "DI", Code: 7},
					},
				},
			},
			0b010: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
						{Name: "SI", Code: 6},
					},
				},
			},
			0b011: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
						{Name: "DI", Code: 7},
					},
				},
			},
			0b100: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "SI", Code: 6},
					},
				},
			},
			0b101: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "DI", Code: 7},
					},
				},
			},
			0b110: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "BP", Code: 5},
					},
				},
			},
			0b111: {
				Type: Operand_Memory,
				EffectiveAddressExpression: EffectiveAddressExpression{
					Displacement: 16,
					Terms: [2]Register{
						{Name: "BX", Code: 3},
					},
				},
			},
		},
	}

	vals, ok := memTables[mod][rm]
	return vals, ok
}

func (it InstructionTable) ResolveImmediate(b int, flags int) (InstructionOperand, bool) {
	var val int
	if flags&int(Bits_W) == int(Bits_W) {
		val = int(int16(b))
	} else {
		val = int(int8(b))
	}
	return InstructionOperand{
		Type: Operand_Immediate,
		Immediate: Immediate{
			Value: val,
			Flags: flags,
		},
	}, true
}
