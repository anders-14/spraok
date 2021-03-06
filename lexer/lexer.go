package lexer

import (
	"github.com/anders-14/spraok/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	char         rune
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) Done() bool {
	return l.nextPosition >= len(l.input)
}

func (l *Lexer) peek() rune {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.nextPosition])
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.char = 0
		return
	}

	l.char = rune(l.input[l.nextPosition])
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

		return token.Token{Type: token.Identifier, Value: word}
	}

	if isInteger(l.char) {
		integer := l.readInteger()
		return token.Token{Type: token.Integer, Value: integer}
	}

	if l.char == '"' {
		string := l.readString()
		l.readChar()
		return token.Token{Type: token.String, Value: string}
	}

	if runeInSlice(l.char, token.PossibleTwoCharOperation) {
		str := string(l.char) + string(l.peek())
		if operation, ok := token.Operations[str]; ok {
			l.readChar()
			return token.Token{Type: operation, Value: str}
		}
	}

	if operation, ok := token.Operations[string(l.char)]; ok {
		return token.Token{Type: operation, Value: string(l.char)}
	}

	return token.Token{Type: token.Invalid, Value: string(l.char)}
}

func runeInSlice(el rune, slice []rune) bool {
	for _, _el := range slice {
		if el == _el {
			return true
		}
	}
	return false
}

func isLetter(char rune) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || char == '_'
}

func isInteger(char rune) bool {
	return '0' <= char && char <= '9'
}
