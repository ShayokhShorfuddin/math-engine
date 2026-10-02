package lexer

type Token struct {
	text string
	kind TokenType
}

func NewToken(text string, kind TokenType) *Token {
	return &Token{text: text, kind: kind}
}

type TokenType int

const (
	EOF TokenType = iota

	// Operators
	Plus
	Minus
	Asterisk
	Slash
)
