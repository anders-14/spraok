package token

import "fmt"

type TokenType uint

const (
	IDENTIFIER TokenType = iota
	INVALID

	// ---- OPERATIONS ----

	ASSIGN
	GT
	LT
	EQUALS
	GT_EQUALS
	LT_EQUALS

	// ---- KEYWORDS ----

	VAR

	// ---- LITERALS ----

	INTEGER
	STRING
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) HumanReadable() string {
	tName, ok := humanReadable[t.Type]
	if !ok {
		tName = "NO HUMANREADABLE FROM ADDED"
	}

	return fmt.Sprintf("(%+v : %+v)", tName, t.Value)
}

var Keywords = map[string]TokenType{
	"var": VAR,
}

var Operations = map[string]TokenType{
	"=":  ASSIGN,
	">":  GT,
	"<":  LT,
	"==": EQUALS,
	">=": GT_EQUALS,
	"<=": LT_EQUALS,
}

var PossibleTwoCharOperation = []byte{'=', '<', '>'}

var humanReadable = map[TokenType]string{
	IDENTIFIER: "IDENTIFIER",
	INVALID:    "INVAILD",
	ASSIGN:     "ASSIGN",
	GT:         "GREATER THAN",
	LT:         "LESS THAN",
	EQUALS:     "EQUALS",
	GT_EQUALS:  "GREATER THAN OR EQUALS",
	LT_EQUALS:  "LESS THAN OR EQUALS",
	VAR:        "VAR",
	INTEGER:    "INTEGER",
	STRING:     "STRING",
}
