package lexer

import "expression-parsing/token"

type Lexer struct {
	Source string
	start  int
	curr   int
}

func isLetter(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func (l Lexer) isAtEnd() bool {
	return l.curr >= len(l.Source)
}

func (l Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.Source[l.curr]
}

func (l *Lexer) consume() byte {
	if l.isAtEnd() {
		return 0
	}
	l.curr += 1
	return l.Source[l.curr-1]
}

func (l *Lexer) match(expected byte) bool {
	if l.peek() == expected {
		l.consume()
		return true
	}
	return false
}

func (l *Lexer) skipWhitespace() {
	for !l.isAtEnd() {
		ch := l.peek()
		if ch == ' ' || ch == '\r' || ch == '\t' || ch == '\n' {
			l.consume()
			continue
		}
		break
	}
}

func (l *Lexer) identToken() token.Token {
	for isLetter(l.peek()) || isDigit(l.peek()) {
		l.consume()
	}
	return token.Token{token.IDENT, l.Source[l.start:l.curr]}
}

func (l *Lexer) numeric() token.Token {
	for isDigit(l.peek()) {
		l.consume()
	}
	if l.match('.') {
		if !isDigit(l.peek()) {
			return token.Token{token.ILLEGAL, l.Source[l.start:l.curr]}
		}
	}
	for isDigit(l.peek()) {
		l.consume()
	}
	return token.Token{token.NUM, l.Source[l.start:l.curr]}
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	l.start = l.curr
	switch ch := l.consume(); ch {
	case '(':
		return token.Token{token.LPAREN, string(ch)}
	case ')':
		return token.Token{token.RPAREN, string(ch)}
	case '+':
		return token.Token{token.ADD, string(ch)}
	case '-':
		return token.Token{token.SUB, string(ch)}
	case '/':
		return token.Token{token.DIV, string(ch)}
	case '%':
		return token.Token{token.MOD, string(ch)}
	case '=':
		return token.Token{token.EQ, string(ch)}
	case '~':
		return token.Token{token.NOT, string(ch)}
	case '&':
		return token.Token{token.AND, string(ch)}
	case '|':
		return token.Token{token.OR, string(ch)}
	case '^':
		return token.Token{token.XOR, string(ch)}
	case '*':
		if l.match('*') {
			return token.Token{token.EXP, l.Source[l.start:l.curr]}
		} else {
			return token.Token{token.MUL, string(ch)}
		}
	case '<':
		if l.match('<') {
			return token.Token{token.LEFT, l.Source[l.start:l.curr]}
		} else {
			return token.Token{token.ILLEGAL, string(ch)}
		}
	case '>':
		if l.match('>') {
			return token.Token{token.RIGHT, l.Source[l.start:l.curr]}
		} else {
			return token.Token{token.ILLEGAL, string(ch)}
		}
	case 0:
		return token.Token{token.EOF, ";"}
	default:
		if isDigit(ch) {
			return l.numeric()
		} else if isLetter(ch) {
			return l.identToken()
		} else {
			return token.Token{token.ILLEGAL, string(ch)}
		}
	}
}
