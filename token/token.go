package token

type TokenType int

const (
	_ = iota
	LPAREN
	RPAREN
	ADD
	SUB
	MUL
	DIV
	MOD
	EXP
	EQ
	EQEQ
	NEQ
	GT
	GTEQ
	LT
	LTEQ
	NOT
	AND
	OR
	XOR
	LOR
	LAND
	LNOT
	LEFT
	RIGHT
	NUM
	IDENT
	ILLEGAL
	EOF
)

type Token struct {
	Type   TokenType
	Lexeme string
}

var toString = map[TokenType]string{
	LPAREN:  "lparen",
	RPAREN:  "rparen",
	ADD:     "addition",
	SUB:     "subtraction",
	MUL:     "multiplication",
	DIV:     "division",
	MOD:     "modulo",
	EXP:     "exponentiation",
	EQ:      "assignment",
	EQEQ:    "equal",
	NEQ:     "not equal",
	GT:      "greater",
	GTEQ:    "greater/equal",
	LT:      "lesser",
	LTEQ:    "lesser/equal",
	NOT:     "bitwise not",
	AND:     "bitwise and",
	OR:      "bitwise or",
	XOR:     "bitwise xor",
	LOR:     "logical or",
	LAND:    "logical and",
	LNOT:    "logical not",
	LEFT:    "bitshift left",
	RIGHT:   "bitshift right",
	NUM:     "number",
	IDENT:   "identifier",
	ILLEGAL: "illegal",
	EOF:     "end",
}

func (tokenType TokenType) String() string {
	return toString[tokenType]
}
