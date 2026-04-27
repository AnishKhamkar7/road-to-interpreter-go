package parser

import (
	"fmt"
	"go-int/src/ast"
	"go-int/src/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {

	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned Nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Program statements Does not consists 3 statements, Got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		iter := program.Statements[i]

		if !testLetStatement(t, iter, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {

	if s.TokenLiteral() != "let" {
		t.Fatalf("s.tokenLiteral not 'let'. Got  = %q ", s.TokenLiteral())
		return false
	}

	letIter, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("S not  *ast.LetStatement. Got = %T", s)
		return false
	}

	if letIter.Name.Value != name {
		t.Errorf("LetIter.Name.Value not %s. Got = %s", name, letIter.Name.Value)
		return false
	}

	if letIter.Name.TokenLiteral() != name {
		t.Errorf("Name not %s, Got = %s)", name, letIter.Name)
		return false
	}

	return true

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parse has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser Error: %q", msg)
	}

	t.FailNow()
}

func TestReturnStatement(t *testing.T) {

	input := `
	return 5;
	return 10;
	return 09090;
	`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("Program statements Does not consists 3 statements, Got = %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not *ast.returnStatement, got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}

}

func TestIdentifyExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)

	p := New(l)

	t.Logf("ParseNEw = %v", p)

	parser := p.ParseProgram()

	t.Logf("Parser Statement = %s", parser.Statements[0].String())

	checkParserErrors(t, p)

	if len(parser.Statements) != 1 {
		t.Fatalf("len(Parser.statement) is not enough, got= %d", len(parser.Statements))

	}

	stmt, ok := parser.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statement[0] is not an ExpressionStatement, got= %T", parser.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("stmt.Expression is not an Ast.Identifier, got= %T", stmt.Expression)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() is not foobar, got=%s", ident.TokenLiteral())
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value is not foobar,got= %s", ident.Value)
	}

}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)

	p := New(l)

	parser := p.ParseProgram()

	checkParserErrors(t, p)

	if len(parser.Statements) != 1 {
		t.Fatalf("len(Parser.statement) is not enough, got= %d", len(parser.Statements))
	}

	stmt, ok := parser.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("parser.Statements[0] is not an ast.ExpressionStatement got =%T", parser.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("stmt.Expression is not implemented by IntegerLiteral, got = %T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Fatalf("Literal.value is not 5 got= %d", literal.Value)
	}

	if literal.Token.Literal != "5" {
		t.Fatalf("Literal.Token.Literal is not 5 got = %s", literal.Token.Literal)
	}

}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		intValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)

		p := New(l)

		parse := p.ParseProgram()

		if len(parse.Statements) != 1 {
			t.Fatalf("Len of parse statement is not 1, got = %d", len(parse.Statements))
		}

		stmt, ok := parse.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("parser.Statements[0] is not an ast.ExpressionStatement got =%T", parse.Statements[0])
		}

		pre, ok := stmt.Expression.(*ast.PrefixExpression)

		if !ok {

			t.Fatalf("stmt>Expression is not an ast.PrefixExpression got =%T", stmt.Expression)
		}

		if pre.Operator != tt.operator {
			t.Fatalf("pre.Operator is not the expected one ")
		}

		if !testIntegerLiteral(t, pre.Right, tt.intValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, right ast.Expression, value int64) bool {
	integ, ok := right.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("right is not an ast.IntegerLiteral got= %T", right)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ Value is not %d got = %d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral() is not matching with Value= %d", value)
		return false
	}

	return true

}
