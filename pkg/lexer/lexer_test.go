package lexer_test

import (
	"testing"

	"github.com/juanpablocruz/sim8086/pkg/instruction"
	"github.com/juanpablocruz/sim8086/pkg/lexer"
	"github.com/juanpablocruz/sim8086/pkg/reader"
)

func TestLexer_NextInstruction(t *testing.T) {
	tests := []struct {
		input []byte
		want  []instruction.Instruction
	}{
		{input: []byte{0x89, 0xd9}, want: []instruction.Instruction{
			{
				Op:        instruction.Op_mov,
				Direction: false,
				Wide:      true,
				Mode:      instruction.Reg,
				Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
				RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
			},
		}},
		{
			input: []byte{0x89, 0xd9, 0x88, 0xe5, 0x89, 0xda, 0x89, 0xde, 0x89, 0xfb, 0x88, 0xc8, 0x88, 0xed, 0x89, 0xc3, 0x89, 0xf3, 0x89, 0xfc, 0x89, 0xc5},
			want: []instruction.Instruction{
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AH"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "SI"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DI"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CL"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "AL"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "SI"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DI"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "SP"}, Type: instruction.Operand_Register},
				},
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BP"}, Type: instruction.Operand_Register},
				},
			},
		},
		{
			input: []byte{
				0x89, 0xde, 0x88, 0xc6, 0xb1, 0x0c, 0xb5, 0xf4, 0xb9, 0x0c, 0x00, 0xb9, 0xf4,
				0xff, 0xba, 0x6c, 0x0f, 0xba, 0x94, 0xf0, 0x8a, 0x00, 0x8b, 0x1b, 0x8b, 0x56,
				0x00, 0x8a, 0x60, 0x04, 0x8a, 0x80, 0x87, 0x13, 0x89, 0x09, 0x88, 0x0a, 0x88,
				0x6e, 0x00,
			},
			want: []instruction.Instruction{
				// Register-to-register
				// mov si, bx
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "SI"}, Type: instruction.Operand_Register},
				},
				// mov dh, al
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AL"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "DH"}, Type: instruction.Operand_Register},
				},
				// 8-bit immediate-to-register
				// mov cl, 12
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CL"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 12}},
				},
				// mov ch, -12
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: -12}},
				},

				// 16-bit immediate-to-register
				// mov cx, 12
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 12}},
				},
				// mov cx, -12
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: -12}},
				},
				// mov dx, 3948
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 3948}},
				},
				// mov dx, -3948
				{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: -3948}},
				},

				// Source address calculation
				// mov al, [bx + si]
				{
					Op:        instruction.Op_mov,
					Direction: true,
					Wide:      false,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AL"}, Type: instruction.Operand_Register},
					RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
						Displacement: 0,
						Terms: [2]instruction.Register{
							{Name: "BX", Code: 3},
							{Name: "SI", Code: 6},
						},
					}},
				},
				// mov bx, [bp + di]
				{
					Op:        instruction.Op_mov,
					Direction: true,
					Wide:      true,
					Mode:      instruction.Memory,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
						Displacement: 0,
						Terms: [2]instruction.Register{
							{Name: "BP"},
							{Name: "DI"},
						},
					}},
				},
				// mov dx, [bp]
				{
					Op:        instruction.Op_mov,
					Direction: true,
					Wide:      true,
					Mode:      instruction.Displ8,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
					RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
						Displacement: 8,
						Terms: [2]instruction.Register{
							{Name: "BP"},
						},
					}},
				},

				// Source address calculation plus 8-bit displacement
				// mov ah, [bx + si + 4]
				/*
					{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Displ8,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AH"}, Type: instruction.Operand_Register},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement: 8,
							Terms: [2]instruction.Register{
								{Name: "BX"},
								{Name: "SI"},
							},
						}},
					},*/
				// Source address calculation plus 16-bit displacement
				// mov al, [bx + si + 4999]

				// Dest address calculation
				// mov [bx + di], cx
				// mov [bp + si], cl
				// mov [bp], ch
			},
		},
	}

	for _, tt := range tests {

		r := reader.Reader{}
		r.Data = tt.input
		l := lexer.New(&r)
		i := 0
		for i < len(tt.want) {
			got := l.NextInstruction()
			if got.Op != tt.want[i].Op {
				t.Errorf("NextToken() invalid opcode. got=%08b want=%08b", got.Op, tt.want[i].Op)
			}
			if got.Direction != tt.want[i].Direction {
				t.Errorf("NextToken() invalid direction. got=%v want=%v", got.Direction, tt.want[i].Direction)
			}
			if got.Wide != tt.want[i].Wide {
				t.Errorf("NextToken() invalid wide. got=%v want=%v", got.Wide, tt.want[i].Wide)
			}
			if got.Mode != tt.want[i].Mode {
				t.Errorf("NextToken() invalid Mode. got=%v want=%v", got.Mode, tt.want[i].Mode)
			}
			if got.Reg.Type != tt.want[i].Reg.Type {
				t.Errorf("NextToken() invalid Reg Type. got=%v want=%v", got.Reg.Type, tt.want[i].Reg.Type)
			}
			if got.Reg.Name != tt.want[i].Reg.Name {
				t.Errorf("NextToken() invalid Reg. got=%v want=%v", got.Reg.Name, tt.want[i].Reg.Name)
			}
			if got.RM.Type != tt.want[i].RM.Type {
				t.Errorf("NextToken() invalid RM Type. got=%v want=%v", got.RM.Type, tt.want[i].RM.Type)
			}
			switch got.RM.Type {
			case instruction.Operand_Register:
				if got.RM.Name != tt.want[i].RM.Name {
					t.Errorf("NextToken() invalid RM. got=%v want=%v", got.RM.Name, tt.want[i].RM.Name)
				}
			case instruction.Operand_Immediate:
				if got.RM.Value != tt.want[i].RM.Value {
					t.Errorf("NextToken() invalid RM. got=%v want=%v", got.RM.Value, tt.want[i].RM.Value)
				}
			case instruction.Operand_Memory:
				if got.RM.Displacement != tt.want[i].RM.Displacement {
					t.Errorf("NextToken() invalid RM Displacement. got=%v want=%v", got.RM.Displacement, tt.want[i].RM.Displacement)
				}
				if len(got.RM.Terms) != len(tt.want[i].RM.Terms) {
					t.Errorf("NextToken() invalid RM Terms length. got=%v want=%v", len(got.RM.Terms), len(tt.want[i].RM.Terms))
				}
				for i2, term := range got.RM.Terms {
					if term.Code > 0 && tt.want[i].RM.Terms[i2].Code > 0 && term.Name != tt.want[i].RM.Terms[i2].Name {
						t.Errorf("NextToken() invalid RM Term %d. got=%v want=%v", i2, term, tt.want[i].RM.Terms[i2])
					}
				}
			}
			i++
		}
		if i != len(tt.want) {
			t.Errorf("NextToken() invalid number of instructions. got=%d want=%d", i, len(tt.want))
		}
	}
}
