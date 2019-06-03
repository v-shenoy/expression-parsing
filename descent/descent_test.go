package descent

import (
	"expression-parsing/lexer"
	"testing"
)

func TestDescentParser(tester *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "a"},
		{"5", "5"},
		{"~a", "(~ a)"},
		{"-a", "(- a)"},
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
	}

	for index, tt := range tests {
		l := lexer.Lexer{Source: tt.input}
		parser := NewParser(l)
		expr := parser.Expression()

		if tt.expected != expr.String() {
			tester.Errorf("test[%d], wrong operator precedence, expected = %s, got = %s",
				index, tt.expected, expr.String())
		}
	}
}

func benchmarkExpr(input string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := lexer.Lexer{Source: input}
		parser := NewParser(l)
		parser.Expression()
	}
}

func BenchmarkComplex(b *testing.B) {
	benchmarkExpr("a**2 + 2*a*b + b**2", b)
	benchmarkExpr("a*b + c*d + e*f", b)
	benchmarkExpr("a>>b - c**3 - e/f", b)
	benchmarkExpr("a&2 + b%3", b)
}

func BenchmarkOr(b *testing.B) {
	benchmarkExpr("a | b | c", b)
	benchmarkExpr("a | b | c | d", b)
}

func BenchmarkXor(b *testing.B) {
	benchmarkExpr("a ^ b ^ c", b)
	benchmarkExpr("a ^ b ^ c ^ d", b)
}

func BenchmarkAnd(b *testing.B) {
	benchmarkExpr("a & b & c", b)
	benchmarkExpr("a & b & c & d", b)
}

func BenchmarkSum(b *testing.B) {
	benchmarkExpr("a + b - c", b)
	benchmarkExpr("a + b + c + d", b)
}

func BenchmarkShift(b *testing.B) {
	benchmarkExpr("a >> 2", b)
	benchmarkExpr("a << 2", b)
}

func BenchmarkProduct(b *testing.B) {
	benchmarkExpr("a * b / c", b)
	benchmarkExpr("a % b * c / d", b)
}

func BenchmarkPrefix(b *testing.B) {
	benchmarkExpr("- 1000", b)
	benchmarkExpr("~true", b)
}

func BenchmarkExp(b *testing.B) {
	benchmarkExpr("2 ** 3 ** 4", b)
	benchmarkExpr("2 ** 3 ** 4 ** 5", b)
}

func BenchmarkPrimary(b *testing.B) {
	benchmarkExpr("1024", b)
	benchmarkExpr("volume", b)
}

func BenchmarkGroup(b *testing.B) {
	benchmarkExpr("a - (b - c)", b)
	benchmarkExpr("a * (b + c)", b)
}
