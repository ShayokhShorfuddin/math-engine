package lexer

type Token struct {
	Text string
	Kind TokenType
}

func NewToken(text string, kind TokenType) *Token {
	return &Token{Text: text, Kind: kind}
}

type TokenType int

const (
	EOF TokenType = iota
	Number

	// Operators
	Plus
	Minus
	Asterisk
	Slash
)
