package lexer

import "strings"

type Lexer struct {
	source           []rune
	currentCharacter rune
	currentPosition  int
}

func NewLexer(source string) *Lexer {
	source = strings.TrimSpace(source)
	lexer := &Lexer{source: []rune(source), currentPosition: -1, currentCharacter: 0} // Since rune cannot be '', 0 must be used.
	lexer.nextChar()

	return lexer
}

func (lexer *Lexer) Lex() []Token {
	var tokens []Token

	// Return empty splice if source is blank ("")
	if len(lexer.source) == 0 {
		return []Token{}
	}

	for lexer.currentCharacter != -1 {
		lexer.skipWhitespace()

		if lexer.currentCharacter == -1 {
			break
		}

		tokens = append(tokens, lexer.getToken())
	}

	return tokens
}

func (lexer *Lexer) nextChar() {
	lexer.currentPosition++

	if lexer.currentPosition >= len(lexer.source) {
		lexer.currentCharacter = -1 // EOF.
	} else {
		lexer.currentCharacter = lexer.source[lexer.currentPosition]
	}
}

func (lexer *Lexer) peek() rune {
	if lexer.currentPosition+1 >= len(lexer.source) {
		return -1
	} else {
		return lexer.source[lexer.currentPosition+1]
	}
}

func (lexer *Lexer) getToken() Token {
	var token Token

	switch lexer.currentCharacter {
	case '+':
		token = *NewToken("+", Plus)
	case '-':
		token = *NewToken("-", Minus)
	case '*':
		token = *NewToken("*", Asterisk)
	case '/':
		token = *NewToken("/", Slash)
	default:
		lexer.abort("Unknown rune: " + string(lexer.currentCharacter))
	}

	lexer.nextChar()

	return token
}

func (lexer *Lexer) abort(message string) {
	panic(message)
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.currentCharacter == ' ' || lexer.currentCharacter == '\t' || lexer.currentCharacter == '\r' {
		lexer.nextChar()
	}
}
