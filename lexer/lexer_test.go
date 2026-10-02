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
