package token

import "fmt"

type Type uint

const (
	Identifier Type = iota
	Invalid

	// ---- OPERATIONS ----

	Assign
	Plus
	Minus
	Mul
	Div
	Gt
	Lt
	Equals
	GtEquals
	LtEquals

	// ---- KEYWORDS ----

	Var
	If
	Then
	Else
	End
	While
	Do

	// ---- LITERALS ----

	Integer
	String
)

type Token struct {
	Type  Type
	Value string
}

func (t Token) HumanReadable() string {
	tName, ok := humanReadable[t.Type]
	if !ok {
		tName = "NO HUMANREADABLE FROM ADDED"
	}

	return fmt.Sprintf("(%+v : %+v)", tName, t.Value)
}

var Keywords = map[string]Type{
	"var":   Var,
	"if":    If,
	"then":  Then,
	"else":  Else,
	"end":   End,
	"while": While,
	"do":    Do,
}

var Operations = map[string]Type{
	"=":  Assign,
	"+":  Plus,
	"-":  Minus,
	"*":  Mul,
	"/":  Div,
	">":  Gt,
	"<":  Lt,
	"==": Equals,
	">=": GtEquals,
	"<=": LtEquals,
}

var PossibleTwoCharOperation = []rune{'=', '<', '>'}

var humanReadable = map[Type]string{
	Identifier: "IDENTIFIER",
	Invalid:    "INVAILD",
	Assign:     "ASSIGN",
	Plus:       "PLUS",
	Minus:      "MINUS",
	Mul:        "MUL",
	Div:        "DIV",
	Gt:         "GREATER THAN",
	Lt:         "LESS THAN",
	Equals:     "EQUALS",
	GtEquals:  "GREATER THAN OR EQUALS",
	LtEquals:  "LESS THAN OR EQUALS",
	Var:        "VAR",
	If:         "IF",
	Then:       "THEN",
	Else:       "ELSE",
	End:        "END",
	While:      "WHILE",
	Do:         "DO",
	Integer:    "INTEGER",
	String:     "STRING",
}
