package lexer_test

import (
	"fmt"
	"testing"

	"github.com/juanpablocruz/sim8086/pkg/instruction"
	"github.com/juanpablocruz/sim8086/pkg/lexer"
	"github.com/juanpablocruz/sim8086/pkg/reader"
)

type testStruct struct {
	instruction instruction.Instruction
	str         string
}
type instructionTest struct {
	input []byte
	want  []testStruct
}

func validateInstructions(t *testing.T, tests []instructionTest) {
	t.Helper()
	for _, tt := range tests {
		r := reader.Reader{}
		r.Data = tt.input
		l := lexer.New(&r)
		i := 0
		for i < len(tt.want) {
			t.Run(tt.want[i].str, func(t *testing.T) {
				got := l.NextInstruction()
				if got.Op != tt.want[i].instruction.Op {
					t.Errorf("NextToken() invalid opcode. got=%08b want=%08b", got.Op, tt.want[i].instruction.Op)
				}
				if got.Direction != tt.want[i].instruction.Direction {
					t.Errorf("NextToken() invalid direction. got=%v want=%v", got.Direction, tt.want[i].instruction.Direction)
				}
				if got.Wide != tt.want[i].instruction.Wide {
					t.Errorf("NextToken() invalid wide. got=%v want=%v", got.Wide, tt.want[i].instruction.Wide)
				}
				if got.Mode != tt.want[i].instruction.Mode {
					t.Errorf("NextToken() invalid Mode. got=%v want=%v", got.Mode, tt.want[i].instruction.Mode)
				}
				if got.String() != tt.want[i].str {
					t.Errorf("NextToken() invalid String representation. got=%s want=%s", got.String(), tt.want[i].str)
				}
				if got.Reg.Type != tt.want[i].instruction.Reg.Type {
					t.Errorf("NextToken() invalid Reg Type. got=%v want=%v", got.Reg.Type, tt.want[i].instruction.Reg.Type)
				}
				if got.Reg.Name != tt.want[i].instruction.Reg.Name {
					t.Errorf("NextToken() invalid Reg. got=%v want=%v", got.Reg.Name, tt.want[i].instruction.Reg.Name)
				}
				if got.RM.Type != tt.want[i].instruction.RM.Type {
					t.Errorf("NextToken() invalid RM Type. got=%v want=%v", got.RM.Type, tt.want[i].instruction.RM.Type)
				}
				switch got.RM.Type {
				case instruction.Operand_Register:
					if got.RM.Name != tt.want[i].instruction.RM.Name {
						t.Errorf("NextToken() invalid RM. got=%v want=%v", got.RM.Name, tt.want[i].instruction.RM.Name)
					}
				case instruction.Operand_Immediate:
					if got.RM.Value != tt.want[i].instruction.RM.Value {
						t.Errorf("NextToken() invalid RM. got=%v want=%v", got.RM.Value, tt.want[i].instruction.RM.Value)
						fmt.Printf("got RM: %v Reg: %v\n", got.RM.Value, got.Reg)
					}
				case instruction.Operand_Memory:
					if got.RM.Displacement != tt.want[i].instruction.RM.Displacement {
						t.Errorf("NextToken() invalid RM Displacement. got=%v want=%v", got.RM.Displacement, tt.want[i].instruction.RM.Displacement)
					}
					if got.RM.DisplacementValue != tt.want[i].instruction.RM.DisplacementValue {
						t.Errorf("NextToken() invalid RM Displacement Value. got=%v want=%v", got.RM.DisplacementValue, tt.want[i].instruction.RM.DisplacementValue)
					}
					if len(got.RM.Terms) != len(tt.want[i].instruction.RM.Terms) {
						t.Errorf("NextToken() invalid RM Terms length. got=%v want=%v", len(got.RM.Terms), len(tt.want[i].instruction.RM.Terms))
					}
					for i2, term := range got.RM.Terms {
						if term.Code > 0 && tt.want[i].instruction.RM.Terms[i2].Code > 0 && term.Name != tt.want[i].instruction.RM.Terms[i2].Name {
							t.Errorf("NextToken() invalid RM Term %d. got=%v want=%v", i2, term, tt.want[i].instruction.RM.Terms[i2])
						}
					}
				}
			})
			if t.Failed() {
				fmt.Println("Failed")
				t.Fatal("Stop")
			}
			i++
		}
		if i != len(tt.want) {
			t.Errorf("NextToken() invalid number of instructions. got=%d want=%d", i, len(tt.want))
		}
	}
}

func TestLexer_Listing37(t *testing.T) {
	tests := []instructionTest{
		{input: []byte{0x89, 0xd9}, want: []testStruct{
			{
				str: "mov cx, bx",
				instruction: instruction.Instruction{
					Op:        instruction.Op_mov,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
					RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
				},
			},
		}},
	}
	validateInstructions(t, tests)
}

func TestLexer_Listing38(t *testing.T) {
	tests := []instructionTest{
		{
			input: []byte{0x89, 0xd9, 0x88, 0xe5, 0x89, 0xda, 0x89, 0xde, 0x89, 0xfb, 0x88, 0xc8, 0x88, 0xed, 0x89, 0xc3, 0x89, 0xf3, 0x89, 0xfc, 0x89, 0xc5},
			want: []testStruct{
				{
					str: "mov cx, bx",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov ch, ah",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      false,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "AH"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov dx, bx",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov si, bx",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "SI"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov bx, di",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "DI"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov al, cl",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      false,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AL"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CL"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov ch, ch",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      false,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov bx, ax",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "AX"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov bx, si",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "SI"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov sp, di",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "SP"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "DI"}, Type: instruction.Operand_Register},
					},
				},
				{
					str: "mov bp, ax",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "BP"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "AX"}, Type: instruction.Operand_Register},
					},
				},
			},
		},
	}
	validateInstructions(t, tests)
}

func TestLexer_Listing39(t *testing.T) {
	tests := []instructionTest{
		{
			input: []byte{
				0x89, 0xde, 0x88, 0xc6, 0xb1, 0xc, 0xb5, 0xf4, 0xb9, 0xc, 0x0, 0xb9, 0xf4, 0xff,
				0xba, 0x6c, 0xf, 0xba, 0x94, 0xf0, 0x8a, 0x0, 0x8b, 0x1b, 0x8b, 0x56, 0x0, 0x8a,
				0x60, 0x4, 0x8a, 0x80, 0x87, 0x13, 0x89, 0x9, 0x88, 0xa, 0x88, 0x6e, 0x0,
			},
			want: []testStruct{
				// Register-to-register
				// mov si, bx
				{
					str: "mov si, bx",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "SI"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "BX"}, Type: instruction.Operand_Register},
					},
				},
				// mov dh, al
				{
					str: "mov dh, al",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      false,
						Mode:      instruction.Reg,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DH"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "AL"}, Type: instruction.Operand_Register},
					},
				},
				// 8-bit immediate-to-register
				// mov cl, 12
				{
					str: "mov cl, 12",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      false,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CL"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 12}},
					},
				},
				// mov ch, -12
				{
					str: "mov ch, -12",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      false,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: -12}},
					},
				},

				// 16-bit immediate-to-register
				// mov cx, 12
				{
					str: "mov cx, 12",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 12}},
					},
				},
				// mov cx, -12
				{
					str: "mov cx, -12",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: -12}},
					},
				},
				// mov dx, 3948
				{
					str: "mov dx, 3948",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 3948}},
					},
				},
				{
					// mov dx, -3948 {
					str: "mov dx, -3948",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: -3948}},
					},
				},

				// Source address calculation
				// mov al, [bx+ si]
				{
					str: "mov al, [bx + si]",
					instruction: instruction.Instruction{
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
				},
				// mov bx, [bp + di]
				{
					str: "mov bx, [bp + di]",
					instruction: instruction.Instruction{
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
				},
				// mov dx, [bp]
				{
					str: "mov dx, [bp]",
					instruction: instruction.Instruction{
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
				},

				// Source address calculation plus 8-bit displacement
				// mov ah, [bx+ si + 4]
				{
					str: "mov ah, [bx + si + 4]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      false,
						Mode:      instruction.Displ8,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AH"}, Type: instruction.Operand_Register},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      8,
							DisplacementValue: 4,
							Terms: [2]instruction.Register{
								{Name: "BX"},
								{Name: "SI"},
							},
						}},
					},
				},
				// Source address calculation plus 16-bit displacement
				// mov al, [bx+ si + 4999]
				{
					str: "mov al, [bx + si + 4999]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      false,
						Mode:      instruction.Displ16,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AL"}, Type: instruction.Operand_Register},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      16,
							DisplacementValue: 4999,
							Terms: [2]instruction.Register{
								{Name: "BX"},
								{Name: "SI"},
							},
						}},
					},
				},
				// Dest address calculation
				// mov [bx+ di], cx
				{
					str: "mov [bx + di], cx",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Memory,
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement: 0,
							Terms: [2]instruction.Register{
								{Name: "BX"},
								{Name: "DI"},
							},
						}},
					},
				},
				// mov [bp + si], cl
				{
					str: "mov [bp + si], cl",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      false,
						Mode:      instruction.Memory,
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CL"}, Type: instruction.Operand_Register},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement: 0,
							Terms: [2]instruction.Register{
								{Name: "BP"},
								{Name: "SI"},
							},
						}},
					},
				},
				// mov [bp], ch
				{
					str: "mov [bp], ch",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      false,
						Mode:      instruction.Displ8,
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CH"}, Type: instruction.Operand_Register},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement: 0,
							Terms: [2]instruction.Register{
								{Name: "BP"},
							},
						}},
					},
				},
			},
		},
	}
	validateInstructions(t, tests)
}

func TestLexer_Listing40(t *testing.T) {
	tests := []instructionTest{
		{
			input: []byte{
				0x8b, 0x41, 0xdb, 0x89, 0x8c, 0xd4, 0xfe, 0x8b, 0x57, 0xe0, 0xc6, 0x3, 0x7, 0xc7, 0x85,
				0x85, 0x3, 0x5b, 0x1, 0x8b, 0x2e, 0x5, 0x0, 0x8b, 0x1e, 0x82, 0xd, 0xa1, 0xfb, 0x9, 0xa1,
				0x10, 0x0, 0xa3, 0xfa, 0x9, 0xa3, 0xf, 0x0,
			},
			want: []testStruct{
				// Signed displacements
				// mov ax, [bx + di - 37]
				{
					str: "mov ax, [bx + di - 37]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Displ8,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "AX"}, Type: instruction.Operand_Register},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      8,
							DisplacementValue: -37,
							Terms: [2]instruction.Register{
								{Name: "BX"},
								{Name: "DI"},
							},
						}},
					},
				},
				// mov [si - 300], cx
				{
					str: "mov [si - 300], cx",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Displ16,
						RM:        instruction.InstructionOperand{Register: instruction.Register{Name: "CX"}, Type: instruction.Operand_Register},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      8,
							DisplacementValue: -300,
							Terms: [2]instruction.Register{
								{Name: "SI"},
							},
						}},
					},
				},
				// mov dx, [bx - 32]
				{
					str: "mov dx, [bx - 32]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Displ8,
						Reg:       instruction.InstructionOperand{Register: instruction.Register{Name: "DX"}, Type: instruction.Operand_Register},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      8,
							DisplacementValue: -32,
							Terms: [2]instruction.Register{
								{Name: "BX"},
							},
						}},
					},
				},
				// Explicit sizes
				// mov [bp + di], byte 7
				{
					str: "mov [bp + di], byte 7",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      false,
						Mode:      instruction.Memory,
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 7}},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      0,
							DisplacementValue: 0,
							Terms: [2]instruction.Register{
								{Name: "BP"},
								{Name: "DI"},
							},
						}},
					},
				},
				// mov [di + 901], word 347
				{
					str: "mov [di + 901], word 347",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Displ16,
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Immediate, Immediate: instruction.Immediate{Value: 347}},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      16,
							DisplacementValue: 901,
							Terms: [2]instruction.Register{
								{Name: "DI"},
							},
						}},
					},
				},
				// Direct address
				// mov bp, [5]
				{
					str: "mov bp, [5]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Type: instruction.Operand_Register, Register: instruction.Register{Name: "BP"}},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      0,
							DisplacementValue: 5,
						}},
					},
				},
				// mov bx, [3458]
				{
					str: "mov bx, [3458]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Type: instruction.Operand_Register, Register: instruction.Register{Name: "BX"}},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      0,
							DisplacementValue: 3458,
						}},
					},
				},
				// Memory-to-accumulator
				// mov ax, [2555]
				{
					str: "mov ax, [2555]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Type: instruction.Operand_Register, Register: instruction.Register{Name: "AX"}},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      0,
							DisplacementValue: 2555,
						}},
					},
				},
				// mov ax, [16]
				{
					str: "mov ax, [16]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Type: instruction.Operand_Register, Register: instruction.Register{Name: "AX"}},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      0,
							DisplacementValue: 16,
						}},
					},
				},
				// Accumulator-to-memory
				// mov [2554], ax
				{
					str: "mov [2554], ax",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Memory,
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Register, Register: instruction.Register{Name: "AX"}},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      0,
							DisplacementValue: 2554,
						}},
					},
				},
				// mov [15], ax

				{
					str: "mov [15], ax",
					instruction: instruction.Instruction{
						Op:        instruction.Op_mov,
						Direction: false,
						Wide:      true,
						Mode:      instruction.Memory,
						RM:        instruction.InstructionOperand{Type: instruction.Operand_Register, Register: instruction.Register{Name: "AX"}},
						Reg: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							Displacement:      0,
							DisplacementValue: 15,
						}},
					},
				},
			},
		},
	}
	validateInstructions(t, tests)
}

func TestLexer_Listing41(t *testing.T) {
	tests := []instructionTest{
		{
			input: []byte{
				0x3, 0x18, 0x3, 0x5e, 0x0, 0x83, 0xc6, 0x2, 0x83, 0xc5, 0x2, 0x83, 0xc1, 0x8, 0x3, 0x5e, 0x0, 0x3, 0x4f, 0x2, 0x2, 0x7a, 0x4, 0x3, 0x7b, 0x6, 0x1, 0x18, 0x1, 0x5e, 0x0, 0x1, 0x5e, 0x0, 0x1, 0x4f, 0x2, 0x0, 0x7a, 0x4, 0x1, 0x7b, 0x6, 0x80, 0x7, 0x22, 0x83, 0x82, 0xe8, 0x3, 0x1d, 0x3, 0x46, 0x0, 0x2, 0x0, 0x1, 0xd8, 0x0, 0xe0, 0x5, 0xe8, 0x3, 0x4, 0xe2, 0x4, 0x9, 0x2b, 0x18, 0x2b, 0x5e, 0x0, 0x83, 0xee, 0x2, 0x83, 0xed, 0x2, 0x83, 0xe9, 0x8, 0x2b, 0x5e, 0x0, 0x2b, 0x4f, 0x2, 0x2a, 0x7a, 0x4, 0x2b, 0x7b, 0x6, 0x29, 0x18, 0x29, 0x5e, 0x0, 0x29, 0x5e, 0x0, 0x29, 0x4f, 0x2, 0x28, 0x7a, 0x4, 0x29, 0x7b, 0x6, 0x80, 0x2f, 0x22, 0x83, 0x29, 0x1d, 0x2b, 0x46, 0x0, 0x2a, 0x0, 0x29, 0xd8, 0x28, 0xe0, 0x2d, 0xe8, 0x3, 0x2c, 0xe2, 0x2c, 0x9, 0x3b, 0x18, 0x3b, 0x5e, 0x0, 0x83, 0xfe, 0x2, 0x83, 0xfd, 0x2, 0x83, 0xf9, 0x8, 0x3b, 0x5e, 0x0, 0x3b, 0x4f, 0x2, 0x3a, 0x7a, 0x4, 0x3b, 0x7b, 0x6, 0x39, 0x18, 0x39, 0x5e, 0x0, 0x39, 0x5e, 0x0, 0x39, 0x4f, 0x2, 0x38, 0x7a, 0x4, 0x39, 0x7b, 0x6, 0x80, 0x3f, 0x22, 0x83, 0x3e, 0xe2, 0x12, 0x1d, 0x3b, 0x46, 0x0, 0x3a, 0x0, 0x39, 0xd8, 0x38, 0xe0, 0x3d, 0xe8, 0x3, 0x3c, 0xe2, 0x3c, 0x9, 0x75, 0x2, 0x75, 0xfc, 0x75, 0xfa, 0x75, 0xfc, 0x74, 0xfe, 0x7c, 0xfc, 0x7e, 0xfa, 0x72, 0xf8, 0x76, 0xf6, 0x7a, 0xf4, 0x70, 0xf2, 0x78, 0xf0, 0x75, 0xee, 0x7d, 0xec, 0x7f, 0xea, 0x73, 0xe8, 0x77, 0xe6, 0x7b, 0xe4, 0x71, 0xe2, 0x79, 0xe0, 0xe2, 0xde, 0xe1, 0xdc, 0xe0, 0xda, 0xe3, 0xd8,
			},
			want: []testStruct{
				{
					str: "add bx, [bx + si]",
					instruction: instruction.Instruction{
						Op:        instruction.Op_add,
						Direction: true,
						Wide:      true,
						Mode:      instruction.Memory,
						Reg:       instruction.InstructionOperand{Type: instruction.Operand_Register, Register: instruction.Register{Name: "BX"}},
						RM: instruction.InstructionOperand{Type: instruction.Operand_Memory, EffectiveAddressExpression: instruction.EffectiveAddressExpression{
							DisplacementValue: 0,
							Displacement:      0,
							Terms: [2]instruction.Register{
								{Name: "BX"},
								{Name: "SI"},
							},
						}},
					},
				},
			},
		},
	}
	validateInstructions(t, tests)
}
