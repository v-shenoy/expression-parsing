package lexer

import (
	"expression-parsing/token"
	"testing"
)

func TestOperators(tester *testing.T) {
	input := "( ) + - * / % ** = ~ & | ^ << >>"

	tests := []token.Token{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.ADD, "+"},
		{token.SUB, "-"},
		{token.MUL, "*"},
		{token.DIV, "/"},
		{token.MOD, "%"},
		{token.EXP, "**"},
		{token.EQ, "="},
		{token.NOT, "~"},
		{token.AND, "&"},
		{token.OR, "|"},
		{token.XOR, "^"},
		{token.LEFT, "<<"},
		{token.RIGHT, ">>"},
		{token.EOF, ";"},
	}

	l := Lexer{Source: input}
	for _, tt := range tests {
		tok := l.NextToken()
		if tt.Type != tok.Type {
			tester.Errorf("wrong token type, expected = %s, got = %s", tt.Type, tok.Type)
		}

		if tt.Lexeme != tok.Lexeme {
			tester.Errorf("wrong lexeme, expected = %s, got = %s", tt.Lexeme, tok.Lexeme)
		}
	}
}

func TestOperands(tester *testing.T) {
	input := "1024 21.2 weight 42.3"

	tests := []token.Token{
		{token.NUM, "1024"},
		{token.NUM, "21.2"},
		{token.IDENT, "weight"},
		{token.NUM, "42.3"},
		{token.EOF, ";"},
	}

	l := Lexer{Source: input}
	for _, tt := range tests {
		tok := l.NextToken()
		if tt.Type != tok.Type {
			tester.Errorf("wrong token type, expected = %v, got = %v", tt.Type, tok.Type)
		}

		if tt.Lexeme != tok.Lexeme {
			tester.Errorf("wrong lexeme, expected = %s, got = %s", tt.Lexeme, tok.Lexeme)
		}
	}
}

func TestLexer(tester *testing.T) {
	input := "perimeter = 2*(l+b)"

	tests := []token.Token{
		{token.IDENT, "perimeter"},
		{token.EQ, "="},
		{token.NUM, "2"},
		{token.MUL, "*"},
		{token.LPAREN, "("},
		{token.IDENT, "l"},
		{token.ADD, "+"},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.EOF, ";"},
	}

	l := Lexer{Source: input}
	for _, tt := range tests {
		tok := l.NextToken()
		if tt.Type != tok.Type {
			tester.Errorf("wrong token type, expected = %v, got = %v", tt.Type, tok.Type)
		}

		if tt.Lexeme != tok.Lexeme {
			tester.Errorf("wrong lexeme, expected = %s, got = %s", tt.Lexeme, tok.Lexeme)
		}
	}
}
