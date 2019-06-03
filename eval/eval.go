package eval

import (
	"errors"
	"math"
	// "fmt"
	"expression-parsing/ast"
	"expression-parsing/token"
)

var variables = make(map[string]float64)

func IsIntegral(value float64) bool {
	return value == float64(int64(value))
}

func Evaluate(expr ast.Expression) (float64, error) {
	if expr == nil {
		return 0, errors.New("Undefined behaviour.")
	}
	switch expr := expr.(type) {
	case *ast.Literal:
		return expr.Value, nil
	case *ast.Identifier:
		value, ok := variables[expr.Token.Lexeme]
		if !ok {
			return 0, errors.New(expr.Token.Lexeme + " not defined.")
		}
		return value, nil
	case *ast.Grouped:
		return Evaluate(expr.Group)
	case *ast.Prefix:
		return EvaluatePrefix(expr)
	case *ast.Infix:
		return EvaluateInfix(expr)
	case *ast.Assignment:
		name := expr.Token.Lexeme
		value, err := Evaluate(expr.Value)
		if err != nil {
			return 0, err
		}
		variables[name] = value
		return value, nil
	default:
		return 0, nil
	}
}

func EvaluatePrefix(expr *ast.Prefix) (float64, error) {
	right, err := Evaluate(expr.Right)
	if err != nil {
		return 0, err
	}

	if expr.Op.Type == token.SUB {
		return -right, nil
	}

	if IsIntegral(right) {
		return float64(^int64(right)), nil
	} else {
		return 0, errors.New("Not operand must be integral.")
	}
}

func EvaluateInfix(expr *ast.Infix) (float64, error) {
	left, errLeft := Evaluate(expr.Left)
	if errLeft != nil {
		return 0, errLeft
	}
	right, errRight := Evaluate(expr.Right)
	if errRight != nil {
		return 0, errRight
	}
	switch expr.Op.Type {
	case token.ADD:
		return left + right, nil
	case token.SUB:
		return left - right, nil
	case token.MUL:
		return left * right, nil
	case token.DIV:
		if right != 0 {
			return left / right, nil
		}
		return 0, errors.New("Division by zero")
	case token.MOD:
		if IsIntegral(left) && IsIntegral(right) {
			if right != 0 {
				return float64(int64(left) % int64(right)), nil
			}
			return 0, errors.New("Division by zero")
		}
		return 0, errors.New("Modulus operand must be integral.")
	case token.EXP:
		return math.Pow(left, right), nil
	case token.OR:
		if IsIntegral(left) && IsIntegral(right) {
			return float64(int64(left) | int64(right)), nil
		}
		return 0, errors.New("Or operand must be integral.")
	case token.XOR:
		if IsIntegral(left) && IsIntegral(right) {
			return float64(int64(left) ^ int64(right)), nil
		}
		return 0, errors.New("Xor operand must be integral.")
	case token.AND:
		if IsIntegral(left) && IsIntegral(right) {
			return float64(int64(left) & int64(right)), nil
		}
		return 0, errors.New("And operand must be integral.")
	case token.LEFT:
		if IsIntegral(left) && IsIntegral(right) {
			if right >= 0 {
				return float64(int64(left) << uint64(right)), nil
			}
			return 0, errors.New("Left shift operand must be integral. Shift value must be unsigned.")
		}
		return 0, errors.New("Left shift operand must be integral. Shift value must be unsigned.")
	case token.RIGHT:
		if IsIntegral(left) && IsIntegral(right) {
			if right >= 0 {
				return float64(int64(left) >> uint64(right)), nil
			}
			return 0, errors.New("Right shift operand must be integral. Shift value must be unsigned.")
		}
		return 0, errors.New("Right shift operand must be integral. Shift value must be unsigned.")
	default:
		return 0, nil
	}
}
