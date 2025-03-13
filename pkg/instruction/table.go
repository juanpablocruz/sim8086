package instruction

import (
	"github.com/juanpablocruz/sim8086/pkg/reader"
)

type OperationType uint

const (
	Op_None OperationType = iota
	Op_mov
	Op_push
	Op_pop
	Op_xchg
	Op_in
	Op_out
	Op_xlat
	Op_lea
	Op_lds
	Op_les
	Op_lahf
	Op_sahf
	Op_pushf
	Op_popf

	Op_add
	Op_adc
	Op_sub
	Op_inc
	Op_sbb
	Op_aaa
	Op_daa
	Op_dec
	Op_neg
	Op_cmp
	Op_aas
	Op_das
	Op_mul
	Op_imul
	Op_aam
	Op_div
	Op_idiv
	Op_aad
	Op_cbw
	Op_cwd

	Op_not
	Op_shl
	Op_shr
	Op_sar
	Op_rol
	Op_ror
	Op_rcl
	Op_rcr

	Op_and
	Op_test
	Op_or
	Op_xor
	Op_rep
	Op_movs
	Op_cmps
	Op_scas
	Op_lods
	Op_stos
	Op_call
	Op_jmp
	Op_ret
	Op_je
	Op_jl
	Op_jle
	Op_jb
	Op_jbe
	Op_jp
	Op_js
	Op_jne
	Op_jnl
	Op_jg
	Op_jnb
	Op_ja
	Op_jnp
	Op_jno
	Op_jns
	Op_loop
	Op_loopz
	Op_loopnz
	Op_jcxz
	Op_int
	Op_int3
	Op_into
	Op_iret
	Op_clc
	Op_cmc
	Op_stc
	Op_cld
	Op_std
	Op_cli
	Op_sti
	Op_hlt
	Op_wait
	Op_esc
	Op_lock
	Op_segment
	Op_Count
)

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

type InstructionTable struct {
	Encodings               []InstructionEncoding
	EncodingCount           int
	MaxInstructionByteCount int
}

var InstructionTable8086 = []InstructionEncoding{
	{Op_mov, []InstructionBits{
		{Bits_Literal, 6, 0, 0b100010},
		{Usage: Bits_D, BitCount: 1},
		{Usage: Bits_W, BitCount: 1},
		{Usage: Bits_MOD, BitCount: 2},
		{Usage: Bits_REG, BitCount: 3},
		{Usage: Bits_RM, BitCount: 3},
	}},
	{Op_mov, []InstructionBits{
		{Bits_Literal, 7, 0, 0b1100011},
		{Usage: Bits_W, BitCount: 1},
		{Usage: Bits_MOD, BitCount: 2},
		{Bits_Literal, 3, 0, 0b000},
		{Usage: Bits_RM, BitCount: 3},
		{Bits_Data, 0, 0, 0},
		{Bits_WMakesDataW, 0, 0, 1},
		{Bits_D, 0, 0, 0},
	}},
	{
		Op_mov, []InstructionBits{
			{Bits_Literal, 4, 0, 0b1011},
			{Usage: Bits_W, BitCount: 1},
			{Usage: Bits_REG, BitCount: 3},
			{Bits_Data, 0, 0, 0},
			{Bits_WMakesDataW, 0, 0, 1},
			{Bits_D, 0, 0, 1},
		},
	},
	{
		Op_mov, []InstructionBits{
			{Bits_Literal, 7, 0, 0b1010000},
			{Usage: Bits_W, BitCount: 1},
			{Bits_Disp, 0, 0, 0},
			{Bits_DispAlwaysW, 0, 0, 0b1},
			{Bits_REG, 0, 0, 0},
			{Bits_MOD, 0, 0, 0},
			{Bits_RM, 0, 0, 0b110},
			{Bits_D, 0, 0, 0b1},
		},
	},
	{
		Op_mov, []InstructionBits{
			{Bits_Literal, 7, 0, 0b1010001},
			{Usage: Bits_W, BitCount: 1},
			{Bits_Disp, 0, 0, 0},
			{Bits_DispAlwaysW, 0, 0, 0b1},
			{Bits_REG, 0, 0, 0},
			{Bits_MOD, 0, 0, 0},
			{Bits_RM, 0, 0, 0b110},
			{Bits_D, 0, 0, 0b0},
		},
	},
	{
		Op_mov, []InstructionBits{
			{Bits_Literal, 6, 0, 0b100011},
			{Usage: Bits_D, BitCount: 1},
			{Bits_Literal, 1, 0, 0b0},
			{Bits_MOD, 0, 0, 0},
			{Bits_Literal, 1, 0, 0b0},
			{Usage: Bits_SR, BitCount: 2},
			{Bits_W, 0, 0, 0b1},
		},
	},
	/*
		{Op_push, []InstructionBits{}},
		{Op_pop, []InstructionBits{}},
		{Op_xchg, []InstructionBits{}},
		{Op_in, []InstructionBits{}},
		{Op_out, []InstructionBits{}},
		{Op_xlat, []InstructionBits{}},
		{Op_lea, []InstructionBits{}},
		{Op_lds, []InstructionBits{}},
		{Op_les, []InstructionBits{}},
		{Op_lahf, []InstructionBits{}},
		{Op_sahf, []InstructionBits{}},
		{Op_pushf, []InstructionBits{}},
		{Op_popf, []InstructionBits{}},
		{Op_add, []InstructionBits{}},
		{Op_adc, []InstructionBits{}},
		{Op_sub, []InstructionBits{}},
		{Op_inc, []InstructionBits{}},
		{Op_sbb, []InstructionBits{}},
		{Op_aaa, []InstructionBits{}},
		{Op_daa, []InstructionBits{}},
		{Op_dec, []InstructionBits{}},
		{Op_neg, []InstructionBits{}},
		{Op_cmp, []InstructionBits{}},
		{Op_aas, []InstructionBits{}},
		{Op_das, []InstructionBits{}},
		{Op_mul, []InstructionBits{}},
		{Op_imul, []InstructionBits{}},
		{Op_aam, []InstructionBits{}},
		{Op_div, []InstructionBits{}},
		{Op_idiv, []InstructionBits{}},
		{Op_aad, []InstructionBits{}},
		{Op_cbw, []InstructionBits{}},
		{Op_cwd, []InstructionBits{}},

		{Op_not, []InstructionBits{}},
		{Op_shl, []InstructionBits{}},
		{Op_shr, []InstructionBits{}},
		{Op_sar, []InstructionBits{}},
		{Op_rol, []InstructionBits{}},
		{Op_ror, []InstructionBits{}},
		{Op_rcl, []InstructionBits{}},
		{Op_rcr, []InstructionBits{}},
		{Op_and, []InstructionBits{}},
		{Op_test, []InstructionBits{}},
		{Op_or, []InstructionBits{}},
		{Op_xor, []InstructionBits{}},
		{Op_rep, []InstructionBits{}},
		{Op_movs, []InstructionBits{}},
		{Op_cmps, []InstructionBits{}},
		{Op_scas, []InstructionBits{}},
		{Op_lods, []InstructionBits{}},
		{Op_stos, []InstructionBits{}},
		{Op_call, []InstructionBits{}},
		{Op_jmp, []InstructionBits{}},
		{Op_ret, []InstructionBits{}},
		{Op_je, []InstructionBits{}},
		{Op_jl, []InstructionBits{}},
		{Op_jle, []InstructionBits{}},
		{Op_jb, []InstructionBits{}},
		{Op_jbe, []InstructionBits{}},
		{Op_jp, []InstructionBits{}},
		{Op_js, []InstructionBits{}},
		{Op_jne, []InstructionBits{}},
		{Op_jnl, []InstructionBits{}},
		{Op_jg, []InstructionBits{}},
		{Op_jnb, []InstructionBits{}},
		{Op_ja, []InstructionBits{}},
		{Op_jnp, []InstructionBits{}},
		{Op_jno, []InstructionBits{}},
		{Op_jns, []InstructionBits{}},
		{Op_loop, []InstructionBits{}},
		{Op_loopz, []InstructionBits{}},
		{Op_loopnz, []InstructionBits{}},
		{Op_jcxz, []InstructionBits{}},
		{Op_int, []InstructionBits{}},
		{Op_int3, []InstructionBits{}},
		{Op_into, []InstructionBits{}},
		{Op_iret, []InstructionBits{}},
		{Op_clc, []InstructionBits{}},
		{Op_cmc, []InstructionBits{}},
		{Op_stc, []InstructionBits{}},
		{Op_cld, []InstructionBits{}},
		{Op_std, []InstructionBits{}},
		{Op_cli, []InstructionBits{}},
		{Op_sti, []InstructionBits{}},
		{Op_hlt, []InstructionBits{}},
		{Op_wait, []InstructionBits{}},
		{Op_esc, []InstructionBits{}},
		{Op_lock, []InstructionBits{}},
		{Op_segment, []InstructionBits{}},*/
}

func New8086InstructionTable() InstructionTable {
	return InstructionTable{
		EncodingCount:           len(InstructionTable8086),
		Encodings:               InstructionTable8086,
		MaxInstructionByteCount: 15,
	}
}

func (it InstructionTable) DecodeInstruction(r *reader.Reader) (Instruction, error) {
	instr := Instruction{}

	for _, encoding := range it.Encodings {
		in, err := it.TryDecode(encoding, r)
		if err != nil {
			return instr, err
		}
		instr = in

		if OperationType(instr.Op) != Op_None {
			break
		}
	}
	return instr, nil
}

func (it InstructionTable) TryDecode(encoding InstructionEncoding, r *reader.Reader) (Instruction, error) {
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
			masked := r.Curr >> (8 - bit.BitCount)

			if bit.Value&masked == masked {
				instr.Op = (encoding.Op)
				bitIndx += int(bit.BitCount)
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

			bits[bit.Usage] |= (r.Curr & byte(mask)) >> (8 - byte(bitIndx))
			has[bit.Usage] = true

		}
	}

	if valid {
		mod := bits[Bits_MOD]
		rm := bits[Bits_RM]
		w := bits[Bits_W]
		//	s := bits[Bits_S] == 1
		d := bits[Bits_D] == 1

		hasDirectAddr := (mod == 0b00) && (rm == 0b110)
		has[Bits_Disp] = ((has[Bits_Disp]) || (mod == 0b10) || hasDirectAddr)

		//	displacementIsW := ((bits[Bits_DispAlwaysW]) != 0 || (mod == 0b10) || hasDirectAddr)
		//		dataIsW := ((bits[Bits_WMakesDataW] != 0) && !s && (w == 0))

		var regOperand Register
		if has[Bits_REG] {
			regOperand, _ = it.ResolveRegister(bits[Bits_REG], w != 0)
		}

		var modOperand Register
		if has[Bits_MOD] {
			// register mode, no displacement
			if mod == 0b11 {
				modOperand, _ = it.ResolveRegister(rm, (w != 0) || (bits[Bits_RMRegAlwaysW] == 1))
			} else {
			}
		}

		if d {
			instr.RM = regOperand
			instr.Reg = modOperand
		} else {
			instr.Reg = regOperand
			instr.RM = modOperand
		}

		instr.Mode = Mode(mod)
		instr.Direction = d
		instr.Wide = w == 1
	}

	return instr, nil
}

func (it InstructionTable) ResolveRegister(b byte, w bool) (Register, bool) {
	regs := map[bool]map[uint8]Register{
		true: {
			0: {Name: "AX", Code: 0},
			1: {Name: "CX", Code: 1},
			2: {Name: "DX", Code: 2},
			3: {Name: "BX", Code: 3},
			4: {Name: "SP", Code: 4},
			5: {Name: "BP", Code: 5},
			6: {Name: "SI", Code: 6},
			7: {Name: "DI", Code: 7},
		},
		false: {
			0: {Name: "AL", Code: 0},
			1: {Name: "CL", Code: 1},
			2: {Name: "DL", Code: 2},
			3: {Name: "BL", Code: 3},
			4: {Name: "AH", Code: 4},
			5: {Name: "CH", Code: 5},
			6: {Name: "DH", Code: 6},
			7: {Name: "BH", Code: 7},
		},
	}

	subSet := regs[w]

	reg, ok := subSet[b]
	return reg, ok
}

var OpcodeMnemonics = []string{
	"",
	"mov",
	"push",
	"pop",
	"xchg",
	"in",
	"out",
	"xlat",
	"lea",
	"lds",
	"les",
	"lahf",
	"sahf",
	"pushf",
	"popf",
	"add",
	"adc",
	"inc",
	"aaa",
	"daa",
	"sub",
	"sbb",
	"dec",
	"neg",
	"cmp",
	"aas",
	"das",
	"mul",
	"imul",
	"aam",
	"div",
	"idiv",
	"aad",
	"cbw",
	"cwd",
	"not",
	"shl",
	"shr",
	"sar",
	"rol",
	"ror",
	"rcl",
	"rcr",
	"and",
	"test",
	"or",
	"xor",
	"rep",
	"movs",
	"cmps",
	"scas",
	"lods",
	"stos",
	"call",
	"jmp",
	"ret",
	"retf",
	"je",
	"jl",
	"jle",
	"jb",
	"jbe",
	"jp",
	"jo",
	"js",
	"jne",
	"jnl",
	"jg",
	"jnb",
	"ja",
	"jnp",
	"jno",
	"jns",
	"loop",
	"loopz",
	"loopnz",
	"jcxz",
	"int",
	"int3",
	"into",
	"iret",
	"clc",
	"stc",
	"cld",
	"std",
	"cli",
	"sti",
	"hlt",
	"wait",
	"esc",
	"lock",
	"segment",
}

func GetMnemonic(op OperationType) string {
	res := ""
	if int(op) < int(Op_Count) {
		res = OpcodeMnemonics[op]
	}
	return res
}

func (op OperationType) String() string {
	return GetMnemonic(op)
}
