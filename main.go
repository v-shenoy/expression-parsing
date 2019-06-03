package main

import (
	"bufio"
	"expression-parsing/descent"
	"expression-parsing/eval"
	"expression-parsing/lexer"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		if text == "exit\n" {
			return
		}
		l := lexer.Lexer{Source: text}
		p := descent.NewParser(l)
		expr := p.Expression()
		if value, err := eval.Evaluate(expr); err != nil {
			fmt.Printf("%s\n", err)
		} else {
			fmt.Println(value)
		}
	}
}
