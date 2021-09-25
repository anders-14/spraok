package lexer

import (
	"github.com/anders-14/spraok/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	char         byte
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) Done() bool {
	return l.nextPosition >= len(l.input)
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.char = 0
		return
	}

	l.char = l.input[l.nextPosition]
	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) readSkippingWhitespace() {
	l.readChar()
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readWord() string {
	start := l.position

	l.readChar()
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) readInteger() string {
	start := l.position

	l.readChar()
	for isInteger(l.char) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) readString() string {
	start := l.position + 1

	l.readChar()
	for l.char != '"' {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) NextToken() token.Token {
	l.readSkippingWhitespace()

	if isLetter(l.char) {
		word := l.readWord()

		if keyword, ok := token.Keywords[word]; ok {
			return token.Token{Type: keyword, Value: word}
		}

		return token.Token{Type: token.IDENTIFIER, Value: word}
	}

	if isInteger(l.char) {
		integer := l.readInteger()
		return token.Token{Type: token.INTEGER, Value: integer}
	}

	if l.char == '"' {
		string := l.readString()
		l.readChar()
		return token.Token{Type: token.STRING, Value: string}
	}

	if operation, ok := token.Operations[string(l.char)]; ok {
		return token.Token{Type: operation, Value: string(l.char)}
	}

	return token.Token{Type: token.INVALID, Value: string(l.char)}
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || char == '_'
}

func isInteger(char byte) bool {
	return '0' <= char && char <= '9'
}
