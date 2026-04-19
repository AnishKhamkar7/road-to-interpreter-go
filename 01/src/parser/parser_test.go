package parser

import (
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
