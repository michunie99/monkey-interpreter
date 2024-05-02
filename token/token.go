package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers, numbers
	IDENT = "IDENT" // x, y, foo etc..
	INT   = "INT"   // 1,2,3 etc..

	// Operators
	ASSIGN = "="
	SUM    = "+"

	// Delimiters
	SEMICOLON = ";"
	COMMA     = ","

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Key work
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
