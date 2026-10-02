package main

import (
	"fmt"
	"math_engine/lexer"
	"math_engine/parser"
)

func main() {
	source := "2 + 3 * 4"
	tokens := lexer.NewLexer(source).Lex()

	result, err := parser.NewParser(tokens).Parse()

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(result)
}
