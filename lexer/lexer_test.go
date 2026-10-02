package lexer

import (
	"fmt"
	"testing"
)

func TestOperators(t *testing.T) {
	source := "+"

	lexer := NewLexer(source)
	tokens := lexer.Lex()
	tokenLength := len(tokens)
	fmt.Println(tokenLength)

	if tokenLength != 1 {
		t.Errorf("Expected %d token, got %d", 1, tokenLength)
	}
}
