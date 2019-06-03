package tdop

import (
	"errors"
	"expression-parsing/ast"
	"expression-parsing/lexer"
	"expression-parsing/token"
	"strconv"
)

const (
	LOWEST = iota
	ASSIGN
	LOR
	LAND
	OR
	XOR
	AND
	EQ
	COMPARE
	SUM
	SHIFT
	PRODUCT
	PREFIX
	EXP
	GROUP
)

type (
	nud func() ast.Expression
	led func(ast.Expression) ast.Expression
)

type Parser struct {
	l    lexer.Lexer
	curr token.Token

	nuds         map[token.TokenType]nud
	leds         map[token.TokenType]led
	bindingPower map[token.TokenType]int
}

func (parser *Parser) move() {
	parser.curr = parser.l.NextToken()
}

func NewParser(l lexer.Lexer) *Parser {
	parser := &Parser{l: l}
	parser.register()
	parser.move()
	return parser
}

func (parser *Parser) register() {
	parser.nuds = map[token.TokenType]nud{
		token.IDENT:  parser.parseIdentifier,
		token.NUM:    parser.parseLiteral,
		token.SUB:    parser.parsePrefix,
		token.NOT:    parser.parsePrefix,
		token.LNOT:   parser.parsePrefix,
		token.LPAREN: parser.parseGroup,
	}

	parser.leds = map[token.TokenType]led{
		token.EQ:    parser.parseEq,
		token.LOR:   parser.parseInfix,
		token.LAND:  parser.parseInfix,
		token.OR:    parser.parseInfix,
		token.XOR:   parser.parseInfix,
		token.AND:   parser.parseInfix,
		token.EQEQ:  parser.parseInfix,
		token.NEQ:   parser.parseInfix,
		token.GT:    parser.parseInfix,
		token.GTEQ:  parser.parseInfix,
		token.LT:    parser.parseInfix,
		token.LTEQ:  parser.parseInfix,
		token.LEFT:  parser.parseInfix,
		token.RIGHT: parser.parseInfix,
		token.ADD:   parser.parseInfix,
		token.SUB:   parser.parseInfix,
		token.MUL:   parser.parseInfix,
		token.DIV:   parser.parseInfix,
		token.MOD:   parser.parseInfix,
		token.EXP:   parser.parseInfix,
	}

	parser.bindingPower = map[token.TokenType]int{
		token.EQ:     ASSIGN,
		token.LOR:    LOR,
		token.LAND:   LAND,
		token.OR:     OR,
		token.XOR:    XOR,
		token.AND:    AND,
		token.EQEQ:   EQ,
		token.NEQ:    EQ,
		token.GT:     COMPARE,
		token.GTEQ:   COMPARE,
		token.LT:     COMPARE,
		token.LTEQ:   COMPARE,
		token.LEFT:   SHIFT,
		token.RIGHT:  SHIFT,
		token.ADD:    SUM,
		token.SUB:    SUM,
		token.MUL:    PRODUCT,
		token.DIV:    PRODUCT,
		token.MOD:    PRODUCT,
		token.NOT:    PREFIX,
		token.LNOT:   PREFIX,
		token.EXP:    EXP,
		token.LPAREN: GROUP,
	}
}

func (parser *Parser) Reset() {
	parser.l.Reset()
	parser.move()
}

func (parser Parser) getPrecedence() int {
	if val, ok := parser.bindingPower[parser.curr.Type]; ok {
		return val
	}
	return LOWEST
}

func (parser *Parser) parseIdentifier() ast.Expression {
	tok := parser.curr
	parser.move()
	return &ast.Identifier{tok}
}

func (parser *Parser) parseLiteral() ast.Expression {
	tok := parser.curr
	parser.move()

	num, _ := strconv.ParseFloat(tok.Lexeme, 64)
	return &ast.Literal{tok, num}
}

func (parser *Parser) parsePrefix() ast.Expression {
	expr := &ast.Prefix{Op: parser.curr}
	parser.move()
	expr.Right = parser.ParseExpression(LOWEST)
	return expr
}

func (parser *Parser) parseInfix(left ast.Expression) ast.Expression {
	expr := &ast.Infix{Left: left, Op: parser.curr}
	precedence := parser.getPrecedence()
	parser.move()
	if expr.Op.Type == token.EQ || expr.Op.Type == token.EXP {
		precedence -= 1
	}
	expr.Right = parser.ParseExpression(precedence)
	return expr
}

func (parser *Parser) parseEq(left ast.Expression) ast.Expression {
	parser.move()
	as, ok := left.(*ast.Identifier)
	if !ok {
		return nil
	}
	value := parser.ParseExpression(LOWEST)
	return &ast.Assignment{as.Token, value}
}

func (parser *Parser) parseGroup() ast.Expression {
	parser.move()
	expr := parser.ParseExpression(LOWEST)
	parser.move()
	return expr
}

func (parser *Parser) ParseExpression(precedence int) ast.Expression {
	tokenNud := parser.nuds[parser.curr.Type]
	if tokenNud == nil {
		return nil
	}

	expr := tokenNud()

	for parser.curr.Type != token.EOF && precedence < parser.getPrecedence() {
		tokenLed := parser.leds[parser.curr.Type]
		if tokenLed == nil {
			return expr
		}
		expr = tokenLed(expr)
	}
	return expr
}

func (parser *Parser) Parse() (ast.Expression, error) {
	expr := parser.ParseExpression(LOWEST)
	if parser.curr.Type != token.EOF {
		return nil, errors.New("Incorrect input format.")
	}
	return expr, nil
}
