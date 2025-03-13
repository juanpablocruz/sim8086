package lexer_test

import (
	"testing"

	"github.com/juanpablocruz/aware/pkg/instruction"
	"github.com/juanpablocruz/aware/pkg/lexer"
	"github.com/juanpablocruz/aware/pkg/reader"
	"github.com/juanpablocruz/aware/pkg/token"
)

func TestLexer_NextInstruction(t *testing.T) {
	tests := []struct {
		input []byte
		want  []instruction.Instruction
	}{
		{input: []byte{0x89, 0xd9}, want: []instruction.Instruction{
			{
				Op:        instruction.Op_mov,
				OpCode:    token.MOV_R,
				Direction: false,
				Wide:      true,
				Mode:      instruction.Reg,
				Reg:       instruction.Register{Name: "BX"},
				RM:        instruction.Register{Name: "CX"},
			},
		}},
		{
			input: []byte{0x89, 0xd9, 0x88, 0xe5, 0x89, 0xda, 0x89, 0xde, 0x89, 0xfb, 0x88, 0xc8, 0x88, 0xed, 0x89, 0xc3, 0x89, 0xf3, 0x89, 0xfc, 0x89, 0xc5},
			want: []instruction.Instruction{
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "BX"},
					RM:        instruction.Register{Name: "CX"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "AH"},
					RM:        instruction.Register{Name: "CH"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "BX"},
					RM:        instruction.Register{Name: "DX"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "BX"},
					RM:        instruction.Register{Name: "SI"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "DI"},
					RM:        instruction.Register{Name: "BX"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "CL"},
					RM:        instruction.Register{Name: "AL"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      false,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "CH"},
					RM:        instruction.Register{Name: "CH"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "AX"},
					RM:        instruction.Register{Name: "BX"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "SI"},
					RM:        instruction.Register{Name: "BX"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "DI"},
					RM:        instruction.Register{Name: "SP"},
				},
				{
					Op:        instruction.Op_mov,
					OpCode:    token.MOV_R,
					Direction: false,
					Wide:      true,
					Mode:      instruction.Reg,
					Reg:       instruction.Register{Name: "AX"},
					RM:        instruction.Register{Name: "BP"},
				},
			},
		},
	}

	for _, tt := range tests {

		r := reader.Reader{}
		r.Data = tt.input
		l := lexer.New(&r)
		i := 0
		for {
			got := l.NextInstruction()
			if got.Op == 0 {
				break
			}
			if i >= len(tt.want) {
				break
			}

			if got.Op != tt.want[i].Op {
				t.Errorf("NextToken() invalid opcode. got=%08b want=%08b", got.OpCode, tt.want[i].OpCode)
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
			if got.Reg.Name != tt.want[i].Reg.Name {
				t.Errorf("NextToken() invalid Reg. got=%v want=%v", got.Reg.Name, tt.want[i].Reg.Name)
			}
			if got.RM.Name != tt.want[i].RM.Name {
				t.Errorf("NextToken() invalid RM. got=%v want=%v", got.RM.Name, tt.want[i].RM.Name)
			}
			i++
		}
		if i != len(tt.want) {
			t.Errorf("NextToken() invalid number of instructions. got=%d want=%d", i, len(tt.want))
		}
	}
}
