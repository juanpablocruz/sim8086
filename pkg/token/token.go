package token

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}

const (
	MOV_R    = (1 << 7) | (1 << 3)
	MOV_I_RM = (1 << 7) | (1 << 6) | (1 << 2) | (1 << 1)
	MOV_I    = (1 << 7) | (1 << 5) | (1 << 4)
	MOV_M_A  = (1 << 7) | (1 << 5)
	MOV_A_M  = (1 << 7) | (1 << 5) | (1 << 1)
	MOV_R_S  = (1 << 7) | (1 << 3) | (1 << 2) | (1 << 1)
	MOV_S_R  = (1 << 7) | (1 << 3) | (1 << 2)

	PUSH_RM = (1 << 7) | (1 << 6) | (1 << 5) | (1 << 4) | (1 << 3) | (1 << 2) | (1 << 1) | 1
	PUSH_R  = (1 << 6) | (1 << 4)
	PUSH_S  = (1 << 2) | (1 << 1)

	POP_RM = (1 << 7) | (1 << 3) | (1 << 2) | (1 << 1) | 1
	POP_R  = (1 << 6) | (1 << 4) | (1 << 3)
	POP_S  = (1 << 2) | (1 << 1) | 1

	ADD_RM   = 0
	ADD_IM   = (1 << 7)
	ADD_IM_A = (1 << 2)
)

var AllTokens = map[TokenType]TokenType{
	MOV_R:    MOV_R,
	MOV_I_RM: MOV_I_RM,
	MOV_I:    MOV_I,
	MOV_M_A:  MOV_M_A,
	MOV_A_M:  MOV_A_M,
	MOV_R_S:  MOV_R_S,
	MOV_S_R:  MOV_S_R,
	PUSH_RM:  PUSH_RM,
	PUSH_R:   PUSH_R,
	PUSH_S:   PUSH_S,
	POP_RM:   POP_RM,
	POP_R:    POP_R,
	POP_S:    POP_S,
	ADD_RM:   ADD_RM,
	ADD_IM:   ADD_IM,
	ADD_IM_A: ADD_IM_A,
}

func (op TokenType) String() string {
	switch op {
	case MOV_A_M, MOV_I, MOV_I_RM, MOV_M_A, MOV_R, MOV_R_S, MOV_S_R:
		return "mov"
	default:
		return ""
	}
}
