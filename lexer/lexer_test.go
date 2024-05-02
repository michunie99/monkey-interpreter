package lexer

import (
	"michunie/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := ",;(){}=+"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.ASSIGN, "="},
		{token.SUM, "+"},
	}

	// Create new lexer
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - wrong tokentype, expected %q, got %q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal, expected %q, got %q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
