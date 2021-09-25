package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ---- TOKENS ----
type TokenType int

const (
	T_KEYWORD TokenType = iota
	T_OPERATION
	T_LITERAL
)

type Token struct {
	Type  TokenType
	Value interface{}
}

// ---- KEYWORDS ----
type KeywordType int

const (
	K_VAR KeywordType = iota
)

var Keywords = map[string]Token{
	"var": {T_KEYWORD, K_VAR},
}

// ---- OPERATIONS ----
type OperationType int

const (
	O_ASSIGN OperationType = iota
)

var Operations = map[string]Token{
	"=": {T_OPERATION, O_ASSIGN},
}

// ---- LEXING ----
func lexLines(lines []string) [][]Token {
	var tokens [][]Token

	for _, line := range lines {

		words := strings.Split(line, " ")
		var tokenLine []Token

		for _, word := range words {
			token, ok := Keywords[word]
			if ok {
				tokenLine = append(tokenLine, token)
				continue
			}

			token, ok = Operations[word]
			if ok {
				tokenLine = append(tokenLine, token)
				continue
			}

			if word, err := strconv.Atoi(word); err == nil {
				token = Token{T_LITERAL, word}
				tokenLine = append(tokenLine, token)
				continue
			}

		}

		tokens = append(tokens, tokenLine)
	}

	return tokens
}

func readLinesFromFile(filename string) []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(contents), "\n")

	// Make sure to remove the empty line, that doesnt really exist,
	// behind the last newline character
	return lines[:len(lines)-1]
}

func main() {
	lines := readLinesFromFile("./test.spraak")
	tokens := lexLines(lines)

	fmt.Println(tokens)
}
