package lexer

import (
	"go-int/src/tokens"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.ASSIGN, "="},
		{tokens.PLUS, "+"},
		{tokens.LPAREN, "("},
		{tokens.RPAREN, ")"},
		{tokens.LBRACE, "{"},
		{tokens.RBRACE, "}"},
		{tokens.COMMA, ","},
		{tokens.SEMICOLON, ";"},
		{tokens.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		println("TokenType -", tok.Type)
		if tok.Type != tt.expectedType {
			println("Why am I here")
			t.Fatalf("tets[%d] - tokenType wrong.  expected=%q,got =%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q,got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {

	input := `let five = 5;
	let ten = 10;
	let add = fn(x,y){
	x + y;
	};
	let res = add(five,ten);
	`

	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.LET, "let"},
		{tokens.IDENT, "five"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "5"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "ten"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "10"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "add"},
		{tokens.ASSIGN, "="},
		{tokens.FUNCTION, "fn"},
		{tokens.LPAREN, "("},
		{tokens.IDENT, "x"},
		{tokens.COMMA, ","},
		{tokens.IDENT, "y"},
		{tokens.RPAREN, ")"},
		{tokens.LBRACE, "{"},
		{tokens.IDENT, "x"},
		{tokens.PLUS, "+"},
		{tokens.IDENT, "y"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "res"},
		{tokens.ASSIGN, "="},
		{tokens.IDENT, "add"},
		{tokens.LPAREN, "("},
		{tokens.IDENT, "five"},
		{tokens.COMMA, ","},
		{tokens.IDENT, "ten"},
		{tokens.RPAREN, ")"},
		{tokens.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		println("TokenType -", tok.Type)
		if tok.Type != tt.expectedType {
			println("Why am I here")
			t.Fatalf("tets[%d] - tokenType wrong.  expected=%q,got =%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q,got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}


func TestNextToken3(t *testing.T){
	input := `let five = 5;
	let ten = 10;
	let add = fn(x,y){
	x + y	
	};
	let result = add(five,ten);
	!-/*5;
	5 < 10 >5;
	`

	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{

	}
	

}