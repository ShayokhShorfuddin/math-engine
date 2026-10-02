package parser

import (
	"fmt"
	"math_engine/lexer"
	"strconv"
)

type Parser struct {
	tokens        []lexer.Token
	currentToken  lexer.Token
	peekToken     lexer.Token
	tokenPosition int
}

func NewParser(tokens []lexer.Token) *Parser {
	tokens = append(tokens, lexer.Token{Kind: lexer.EOF})
	parser := &Parser{tokens: tokens, tokenPosition: -1}
	parser.nextToken()

	return parser
}

// Parse evaluates one arithmetic expression.
func (parser *Parser) Parse() (float64, error) {
	if len(parser.tokens) == 1 {
		return 0, fmt.Errorf("expression is empty")
	}

	value, err := parser.parseExpression()
	if err != nil {
		return 0, err
	}
	if !parser.checkCurrentToken(lexer.EOF) {
		return 0, fmt.Errorf("unexpected token %q", parser.currentToken.Text)
	}
	return value, nil
}

func (parser *Parser) parseExpression() (float64, error) {
	value, err := parser.parseTerm()
	if err != nil {
		return 0, err
	}

	for parser.checkCurrentToken(lexer.Plus) || parser.checkCurrentToken(lexer.Minus) {
		operator := parser.currentToken.Kind
		parser.nextToken()
		right, err := parser.parseTerm()
		if err != nil {
			return 0, err
		}
		if operator == lexer.Plus {
			value += right
		} else {
			value -= right
		}
	}
	return value, nil
}

func (parser *Parser) parseTerm() (float64, error) {
	value, err := parser.parseNumber()
	if err != nil {
		return 0, err
	}

	for parser.checkCurrentToken(lexer.Asterisk) || parser.checkCurrentToken(lexer.Slash) {
		operator := parser.currentToken.Kind

		parser.nextToken()

		right, err := parser.parseNumber()
		if err != nil {
			return 0, err
		}
		if operator == lexer.Asterisk {
			value *= right
		} else {
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			value /= right
		}
	}
	return value, nil
}

func (parser *Parser) parseNumber() (float64, error) {
	if !parser.checkCurrentToken(lexer.Number) {
		return 0, fmt.Errorf("expected number, got %q", parser.currentToken.Text)
	}

	value, err := strconv.ParseFloat(parser.currentToken.Text, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", parser.currentToken.Text, err)
	}
	parser.nextToken()
	return value, nil
}

func (parser *Parser) checkCurrentToken(kind lexer.TokenType) bool {
	return kind == parser.currentToken.Kind
}

func (parser *Parser) checkPeekToken(kind lexer.TokenType) bool {
	return kind == parser.peekToken.Kind
}

// Try to match current token. If not, error. Advances the current token.
func (parser *Parser) match(kind lexer.TokenType) {
	if !parser.checkCurrentToken(kind) {
		message := fmt.Sprintf("Expected %#v go %#v", kind, parser.currentToken.Kind)
		parser.abort(message)
	}

	parser.nextToken()
}

// Advances the current token
func (parser *Parser) nextToken() {
	parser.tokenPosition++

	if parser.tokenPosition >= len(parser.tokens) {
		parser.currentToken = lexer.Token{Kind: lexer.EOF}
		parser.peekToken = parser.currentToken
		return
	}

	parser.currentToken = parser.tokens[parser.tokenPosition]

	if parser.tokenPosition+1 < len(parser.tokens) {
		parser.peekToken = parser.tokens[parser.tokenPosition+1]
	} else {
		parser.peekToken = lexer.Token{Kind: lexer.EOF}
	}
}

func (lexer *Parser) abort(message string) {
	panic(message)
}
