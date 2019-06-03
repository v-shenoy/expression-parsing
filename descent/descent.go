package descent

import (
	"errors"
	"expression-parsing/ast"
	"expression-parsing/lexer"
	"expression-parsing/token"
	"strconv"
)

type Parser struct {
	l    lexer.Lexer
	curr token.Token
}

func (parser *Parser) move() {
	parser.curr = parser.l.NextToken()
}

func NewParser(l lexer.Lexer) *Parser {
	parser := &Parser{l: l}

	parser.move()
	return parser
}

func (parser *Parser) Reset() {
	parser.l.Reset()
	parser.move()
}

func (parser *Parser) Parse() (ast.Expression, error) {
	expr := parser.Expression()
	if parser.curr.Type != token.EOF {
		return nil, errors.New("Incorrect input format.")
	}
	return expr, nil
}

func (parser *Parser) Expression() ast.Expression {
	return parser.assign()
}

func (parser *Parser) assign() ast.Expression {
	left := parser.lor()
	if parser.curr.Type == token.EQ {
		parser.move()
		as, ok := left.(*ast.Identifier)
		if !ok {
			return nil
		}
		value := parser.Expression()
		return &ast.Assignment{as.Token, value}
	}
	return left
}

func (parser *Parser) lor() ast.Expression {
	expr := parser.land()
	for parser.curr.Type == token.LOR {
		op := parser.curr
		parser.move()
		right := parser.land()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) land() ast.Expression {
	expr := parser.or()
	for parser.curr.Type == token.LAND {
		op := parser.curr
		parser.move()
		right := parser.or()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) or() ast.Expression {
	expr := parser.xor()
	for parser.curr.Type == token.OR {
		op := parser.curr
		parser.move()
		right := parser.xor()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) xor() ast.Expression {
	expr := parser.and()
	for parser.curr.Type == token.XOR {
		op := parser.curr
		parser.move()
		right := parser.and()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) and() ast.Expression {
	expr := parser.eq()
	for parser.curr.Type == token.AND {
		op := parser.curr
		parser.move()
		right := parser.eq()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) eq() ast.Expression {
	expr := parser.compare()
	for parser.curr.Type == token.EQEQ || parser.curr.Type == token.NEQ {
		op := parser.curr
		parser.move()
		right := parser.compare()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) compare() ast.Expression {
	expr := parser.sum()
	for parser.curr.Type == token.GT || parser.curr.Type == token.GTEQ || parser.curr.Type == token.LT || parser.curr.Type == token.LTEQ {
		op := parser.curr
		parser.move()
		right := parser.sum()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) sum() ast.Expression {
	expr := parser.shift()
	for parser.curr.Type == token.ADD || parser.curr.Type == token.SUB {
		op := parser.curr
		parser.move()
		right := parser.shift()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) shift() ast.Expression {
	expr := parser.product()
	for parser.curr.Type == token.LEFT || parser.curr.Type == token.RIGHT {
		op := parser.curr
		parser.move()
		right := parser.product()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) product() ast.Expression {
	expr := parser.prefix()
	for parser.curr.Type == token.MUL || parser.curr.Type == token.DIV || parser.curr.Type == token.MOD {
		op := parser.curr
		parser.move()
		right := parser.prefix()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) prefix() ast.Expression {
	if parser.curr.Type == token.NOT || parser.curr.Type == token.SUB || parser.curr.Type == token.LNOT {
		op := parser.curr
		parser.move()
		right := parser.exponent()
		return &ast.Prefix{op, right}
	}
	return parser.exponent()
}

func (parser *Parser) exponent() ast.Expression {
	expr := parser.primary()
	if parser.curr.Type == token.EXP {
		op := parser.curr
		parser.move()
		right := parser.exponent()
		expr = &ast.Infix{expr, op, right}
	}
	return expr
}

func (parser *Parser) primary() ast.Expression {
	tok := parser.curr
	parser.move()

	if tok.Type == token.LPAREN {
		grp := parser.Expression()
		parser.move()
		return &ast.Grouped{grp}
	}

	if tok.Type == token.NUM {
		num, _ := strconv.ParseFloat(tok.Lexeme, 64)
		return &ast.Literal{tok, num}
	}

	if tok.Type == token.IDENT {
		return &ast.Identifier{tok}
	}

	return nil
}
