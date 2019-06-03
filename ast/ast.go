package ast

import "expression-parsing/token"

type Expression interface {
	String() string
}

type Infix struct {
	Left  Expression
	Op    token.Token
	Right Expression
}

func (in Infix) String() string {
	return "(" + in.Left.String() + " " + in.Op.Lexeme + " " + in.Right.String() + ")"
}

type Prefix struct {
	Op    token.Token
	Right Expression
}

func (pre Prefix) String() string {
	return "(" + pre.Op.Lexeme + " " + pre.Right.String() + ")"
}

type Grouped struct {
	Group Expression
}

func (grp Grouped) String() string {
	return grp.Group.String()
}

type Literal struct {
	Token token.Token
	Value float64
}

func (lit Literal) String() string {
	return lit.Token.Lexeme
}

type Identifier struct {
	Token token.Token
}

func (id Identifier) String() string {
	return id.Token.Lexeme
}

type Assignment struct {
	Token token.Token
	Value Expression
}

func (as Assignment) String() string {
	return "(" + as.Token.Lexeme + " = " + as.Value.String() + ")"
}
