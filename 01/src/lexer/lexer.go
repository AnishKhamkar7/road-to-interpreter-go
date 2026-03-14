package lexer

import (
	"go-int/src/tokens"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()

	return l

}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0

	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1

}
func (l *Lexer) NextToken() tokens.Token {
	var tok tokens.Token

	l.skipWhitespace()

	switch l.ch {

	case '=':
		tok = newToken(tokens.ASSIGN, l.ch)
	case ';':
		tok = newToken(tokens.SEMICOLON, l.ch)
	case '(':
		tok = newToken(tokens.LPAREN, l.ch)
	case ')':
		tok = newToken(tokens.RPAREN, l.ch)
	case ',':
		tok = newToken(tokens.COMMA, l.ch)
	case '+':
		tok = newToken(tokens.PLUS, l.ch)
	case '{':
		tok = newToken(tokens.LBRACE, l.ch)
	case '}':
		tok = newToken(tokens.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = tokens.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = tokens.LookUpIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(tokens.ILLEGAL, l.ch)
		}

	}

	l.readChar()

	return tok

}

func newToken(tokenType tokens.TokenType, ch byte) tokens.Token {
	return tokens.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}
