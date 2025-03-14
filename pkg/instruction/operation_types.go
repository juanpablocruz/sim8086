package instruction

import "fmt"

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
	mnemonic := GetMnemonic(op)

	if mnemonic == "" {
		mnemonic = fmt.Sprintf("ERROR(OP): %d", op)
	}
	return mnemonic
}
