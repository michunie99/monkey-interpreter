package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParserProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return  838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParserProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		stmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected *ast.ReturnStatement. Got %T",
				stmt)
			continue
		}
		if stmt.TokenLiteral() != "return" {
			t.Errorf("Expected token literal to be 'return'. Got %q",
				stmt.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Got=%q", s.TokenLiteral())
		return false
	}

	// NOTE: funny thinks happen here can nill be cast
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. Got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. Got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral not %s. Got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()

	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("len(program.Statements) expected 1. Got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] expected type ast.ExpressionStatement. Got %T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression expected type *ast.Identifier. Got %T",
			stmt.Expression)
	}

	if ident.String() != "foobar" {
		t.Errorf("ident.String() expected 'foobar'. Got %q", ident.String())
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() expected 'foobar'. Got %q", ident.TokenLiteral())
	}
}

func TestIntegerLiteralEpxression(t *testing.T) {
	input := "2137;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()

	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("len(program.Statements) expected 1. Got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] expected type ast.ExpressionStatement. Got %T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Expression expected type *ast.IntegerLiteral. Got %T",
			stmt.Expression)
	}

	if literal.Value != 2137 {
		t.Errorf("ident.String() expected 'foobar'. Got %q", literal.String())
	}
	if literal.TokenLiteral() != "2137" {
		t.Errorf("ident.TokenLiteral() expected 'foobar'. Got %q", literal.TokenLiteral())
	}
}

func TestPrefixEpxression(t *testing.T) {
	prefixTest := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTest {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParserProgram()
		checkParseErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("len(program.Statements) expected 1. Got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] expected type ast.ExpressionStatement. Got %T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression expected type *ast.PrefixExpression. Got %T",
				stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Errorf("exp.TokenLiteral() expected %q. Got %q", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}

}

func testIntegerLiteral(t *testing.T, il ast.Expression, integerValue int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. Got %T", il)
		return false
	}

	if integ.Value != integerValue {
		t.Errorf("expected integ.Value %d. Got %d", integerValue, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", integerValue) {
		t.Errorf("integ.TokenLiteral not %d. Got %s", integerValue, integ.TokenLiteral())
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser encountered %d errors ", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
