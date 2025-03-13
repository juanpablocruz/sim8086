package lexer

import (
	"github.com/juanpablocruz/sim8086/pkg/instruction"
	"github.com/juanpablocruz/sim8086/pkg/reader"
)

type Lexer struct {
	readPosition int
	position     int
	ch           byte

	r     *reader.Reader
	table instruction.InstructionTable
}

func New(r *reader.Reader) *Lexer {
	table := instruction.New8086InstructionTable()

	l := &Lexer{r: r, table: table}
	l.readByte()
	return l
}

func (l *Lexer) readByte() {
	l.r.ReadByte()
}

func (l *Lexer) NextInstruction() instruction.Instruction {
	var tok instruction.Instruction

	if l.r.Curr == 0 {
		return tok
	}

	in, err := l.table.DecodeInstruction(l.r)
	if err != nil {
		return tok
	}
	l.readByte()

	return in
}
