package token

type TokenType uint

const (
	IDENTIFIER TokenType = iota

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

var Keywords = map[string]TokenType{
	"var": VAR,
}

var Operations = map[string]TokenType{
	"=": ASSIGN,
}
