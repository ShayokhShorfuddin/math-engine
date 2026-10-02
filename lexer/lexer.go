package lexer

import (
	"strings"
	"unicode"
)

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

	switch {
	case lexer.currentCharacter == '+':
		token = *NewToken("+", Plus)
	case lexer.currentCharacter == '-':
		token = *NewToken("-", Minus)
	case lexer.currentCharacter == '*':
		token = *NewToken("*", Asterisk)
	case lexer.currentCharacter == '/':
		token = *NewToken("/", Slash)
	case unicode.IsDigit(lexer.currentCharacter):
		startingPosition := lexer.currentPosition

		for unicode.IsDigit(lexer.peek()) || lexer.peek() == '.' {
			if lexer.peek() == '.' { // Decimal
				lexer.nextChar()

				// There must be at least one digit after the dot.
				if !unicode.IsDigit(lexer.peek()) {
					lexer.abort("Illegal character in number.")
				}

				for unicode.IsDigit(lexer.peek()) {
					lexer.nextChar()
				}
			} else {
				for unicode.IsDigit(lexer.peek()) {
					lexer.nextChar()
				}
			}
		}

		numberText := lexer.source[startingPosition : lexer.currentPosition+1]
		token = *NewToken(string(numberText), Number)

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
