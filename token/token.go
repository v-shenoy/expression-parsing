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
	NOT
	AND
	OR
	XOR
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
	NOT:     "bitwise not",
	AND:     "bitwise and",
	OR:      "bitwise or",
	XOR:     "bitwise xor",
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
