package main

import (
	"fmt"
	"math_engine/lexer"
)

func main() {
	source := "  +   +     +  "
	lexer := lexer.NewLexer(source)
	tokens := lexer.Lex()

	// TODO: Implement tests
	// TODO: Upload to GitHub

	fmt.Println(tokens)
}
