package token

import "fmt"

type TokenType uint

const (
	IDENTIFIER TokenType = iota
	INVALID

	// ---- OPERATIONS ----

	ASSIGN

	// ---- KEYWORDS ----

	VAR

	// ---- LITERALS ----

	INTEGER
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
	"=": ASSIGN,
}

var humanReadable = map[TokenType]string{
	IDENTIFIER: "IDENTIFIER",
	INVALID:    "INVAILD",
	ASSIGN:     "ASSIGN",
	VAR:        "VAR",
	INTEGER:    "INTEGER",
}
