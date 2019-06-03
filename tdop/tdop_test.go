package tdop

import (
	"expression-parsing/lexer"
	"testing"
)

func TestTdopParser(tester *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "a"},
		{"5", "5"},
		{"~a", "(~ a)"},
		{"-a", "(- a)"},
		{"!a", "(! a)"},
		{"a|b", "(a | b)"},
		{"a	^b", "(a ^ b)"},
		{"a&b", "(a & b)"},
		{"a+b", "(a + b)"},
		{"a-b", "(a - b)"},
		{"a<<b", "(a << b)"},
		{"a>>b", "(a >> b)"},
		{"a*b", "(a * b)"},
		{"a/b", "(a / b)"},
		{"a**b", "(a ** b)"},
		{"x = 5", "(x = 5)"},
		{"a + b + c", "((a + b) + c)"},
		{"a - b - c", "((a - b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a / b / c", "((a / b) / c)"},
		{"a % b % c", "((a % b) % c)"},
		{"a + b * c", "(a + (b * c))"},
		{"a + b * c / d", "(a + ((b * c) / d))"},
		{"a & b + c", "(a & (b + c))"},
		{"a<<b + c<<d", "((a << b) + (c << d))"},
		{"a * (b + c)", "(a * (b + c))"},
		{"a - (b - c) * d", "(a - ((b - c) * d))"},
		{"-a ** b", "(- (a ** b))"},
		{"a ** b ** c", "(a ** (b ** c))"},
		{"a**2 + 2*a*b + b**2", "(((a ** 2) + ((2 * a) * b)) + (b ** 2))"},
		{"a = b = 5", "(a = (b = 5))"},
		{"a - b == c * d", "((a - b) == (c * d))"},
		{"a == b > c == d", "((a == (b > c)) == d)"},
	}

	for index, tt := range tests {
		l := lexer.Lexer{Source: tt.input}
		parser := NewParser(l)
		expr := parser.ParseExpression(LOWEST)

		if tt.expected != expr.String() {
			tester.Errorf("test[%d], wrong operator precedence, expected = %s, got = %s",
				index, tt.expected, expr.String())
		}
	}
}

func benchmarkExpr(input string, b *testing.B) {
	l := lexer.Lexer{Source: input}
	parser := NewParser(l)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		parser.ParseExpression(LOWEST)
		parser.Reset()
	}
}

func BenchmarkOr(b *testing.B) {
	benchmarkExpr("a | b", b)
	benchmarkExpr("a | b | c", b)
	benchmarkExpr("a | b | c | d | e | f", b)
}

func BenchmarkXor(b *testing.B) {
	benchmarkExpr("a ^ b", b)
	benchmarkExpr("a ^ b ^ c", b)
	benchmarkExpr("a ^ b ^ c ^ d ^ e ^ f", b)
}

func BenchmarkAnd(b *testing.B) {
	benchmarkExpr("a & b", b)
	benchmarkExpr("a & b & c", b)
	benchmarkExpr("a & b & c & d & e & f", b)
}

func BenchmarkSum(b *testing.B) {
	benchmarkExpr("a + b", b)
	benchmarkExpr("a - b - c", b)
	benchmarkExpr("a + b - c + d - e + f", b)
}

func BenchmarkShift(b *testing.B) {
	benchmarkExpr("a >> b", b)
	benchmarkExpr("a << b << c", b)
	benchmarkExpr("a >> b >> c << d << e >> f", b)
}

func BenchmarkProduct(b *testing.B) {
	benchmarkExpr("a * b", b)
	benchmarkExpr("a / b % c", b)
	benchmarkExpr("a * b % c * d / e / f", b)
}

func BenchmarkPrefix(b *testing.B) {
	benchmarkExpr("-1000", b)
	benchmarkExpr("~true", b)
}

func BenchmarkExp(b *testing.B) {
	benchmarkExpr("a ** b", b)
	benchmarkExpr("a ** b ** c", b)
	benchmarkExpr("a ** b ** c ** d ** e ** f", b)
}

func BenchmarkPrimary(b *testing.B) {
	benchmarkExpr("whatever", b)
	benchmarkExpr("12412", b)
	benchmarkExpr("sigh", b)
}

func BenchmarkGroup(b *testing.B) {
	benchmarkExpr("(a+b)", b)
	benchmarkExpr("((a+b))", b)
	benchmarkExpr("(((a+b)))", b)
}
