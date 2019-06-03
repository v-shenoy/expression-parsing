package eval

import (
	"expression-parsing/descent"
	"expression-parsing/lexer"
	"testing"
)

func TestEvaluation(tester *testing.T) {
	tests := []struct {
		input string
		value float64
	}{
		{"1+2", 3},
		{"4*5", 20},
		{"10/2", 5},
		{"27%4", 3},
		{"2**3", 8},
		{"100-64", 36},
		{"5 + 2 * 4", 13},
		{"- (2 + 20)", -22},
		{"(100 + 200)", 300},
		{"~5", ^5},
		{"3 | 12", 3 | 12},
		{"3 ^ 5", 3 ^ 5},
		{"3 & 5", 3 & 5},
		{"1 << 2", 4},
		{"16 >> 1", 8},
		{"x  * 2 + 5", 25},
	}

	variables["x"] = 10

	for index, tt := range tests {
		l := lexer.Lexer{Source: tt.input}
		p := descent.NewParser(l)
		expr := p.Expression()
		value, _ := Evaluate(expr)
		if tt.value != value {
			tester.Errorf("test[%d], incorrect value, expected = %f, got = %f",
				index, tt.value, value)
		}
	}
}
