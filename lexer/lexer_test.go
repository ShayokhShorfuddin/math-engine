package lexer

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []Token
	}{
		{"no source", "", []Token{}},

		{"source with whitespace and tabs", " 	 	", []Token{}},

		{"operators", "+-*/", []Token{
			*NewToken("+", Plus),
			*NewToken("-", Minus),
			*NewToken("*", Asterisk),
			*NewToken("/", Slash)},
		},

		{"operators with whitespace", "  + - * /  ", []Token{
			*NewToken("+", Plus),
			*NewToken("-", Minus),
			*NewToken("*", Asterisk),
			*NewToken("/", Slash)},
		},

		{"whole numbers and decimal numbers", "123 456 789 3.1416 123.456", []Token{
			*NewToken("123", Number),
			*NewToken("456", Number),
			*NewToken("789", Number),
			*NewToken("3.1416", Number),
			*NewToken("123.456", Number)},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewLexer(test.source).Lex()

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Source: %q\nGot: %#v\nWant: %#v", test.source, got, test.want)
			}
		})
	}
}
