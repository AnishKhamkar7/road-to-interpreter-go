package lexer

import (
	"go-int/src/tokens"
	"testing"
)

func TestNextToken (t *testing.T){
	input := `=+(){},;`

	tests := []struct {
		expectedType tokens.TokenType
		expectedLiteral string

	} {
		{tokens.ASSIGN,"="}, 
		{tokens.PLUS,"+"},
		{tokens.LPAREN,"("},
		{tokens.RPAREN,")"},
		{tokens.LBRACE,"{"},
		{tokens.RBRACE,"}"},
		{tokens.COMMA,","},
		{tokens.SEMICOLON,";"},
		{tokens.EOF,""},
		
	}

	l := New(input)

	for i,tt := range tests{
		tok := l.NextToken()
		println("TokenType -",tok.Type)
		if tok.Type != tt.expectedType {
			println("Why am I here")
			t.Fatalf("tets[%d] - tokenType wrong.  expected=%q,got =%q",
		i,tt.expectedType,tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q,got=%q",i,tt.expectedLiteral,tok.Literal)
		}
	} 


}
