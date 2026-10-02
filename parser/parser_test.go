package parser

import (
	"math_engine/lexer"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		source string
		want   float64
	}{
		{"2 + 3 * 4", 14},
		{"10 / 2 - 1.5", 3.5},
		{"42", 42},
	}

	for _, test := range tests {
		t.Run(test.source, func(t *testing.T) {
			got, err := NewParser(lexer.NewLexer(test.source).Lex()).Parse()
			if err != nil {
				t.Fatalf("Parse() returned an error: %v", err)
			}
			if got != test.want {
				t.Errorf("Parse() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	for _, source := range []string{"", "2 +", "2 / 0"} {
		t.Run(source, func(t *testing.T) {
			if _, err := NewParser(lexer.NewLexer(source).Lex()).Parse(); err == nil {
				t.Error("Parse() returned nil error")
			}
		})
	}
}
