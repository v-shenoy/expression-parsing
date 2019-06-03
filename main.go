package main

import (
	"bufio"
	"expression-parsing/descent"
	"expression-parsing/eval"
	"expression-parsing/lexer"
	"expression-parsing/tdop"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	if len(os.Args) > 2 {
		fmt.Println("Incorrect usage. Sample usage : go run main.go [descent/tdop].")
		return
	}
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "descent") {
		for {
			fmt.Print(">>> ")
			text, _ := reader.ReadString('\n')
			if text == "exit\n" {
				return
			}
			l := lexer.Lexer{Source: text}
			p := descent.NewParser(l)
			if expr, parseErr := p.Parse(); parseErr != nil {
				fmt.Println(parseErr)
			} else {
				if value, err := eval.Evaluate(expr); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(value)
				}
			}
		}
	} else if os.Args[1] == "tdop" {
		for {
			fmt.Print(">>> ")
			text, _ := reader.ReadString('\n')
			if text == "exit\n" {
				return
			}
			l := lexer.Lexer{Source: text}
			p := tdop.NewParser(l)
			if expr, parseErr := p.Parse(); parseErr != nil {
				fmt.Println(parseErr)
			} else {
				if value, err := eval.Evaluate(expr); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(value)
				}
			}
		}
	} else {
		fmt.Println("Parser not found. Available choices : descent, tdop.")
	}
}
