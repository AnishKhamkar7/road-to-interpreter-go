package parser

import (
	"fmt"
	"go-int/src/ast"
	"go-int/src/lexer"
	"go-int/src/tokens"
	"strconv"
)

type (
	preFixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

const (
	_int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	l              *lexer.Lexer
	curToken       tokens.Token
	peekToken      tokens.Token
	errors         []string
	preFixParseFns map[tokens.TokenType]preFixParseFn
	infixParseFns  map[tokens.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tn tokens.TokenType, fn preFixParseFn) {
	p.preFixParseFns[tn] = fn
}

func (p *Parser) registerInfix(tn tokens.TokenType, fn infixParseFn) {
	p.infixParseFns[tn] = fn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.preFixParseFns = make(map[tokens.TokenType]preFixParseFn)
	p.registerPrefix(tokens.IDENT, p.parseIdentifier)
	p.registerPrefix(tokens.INT, p.parseIntegerLiteral)
	p.registerPrefix(tokens.BANG, p.parsePrefix)
	p.registerPrefix(tokens.MINUS, p.parsePrefix)
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parsePrefix() ast.Expression {
	exp := &ast.PrefixExpression{Operator: p.curToken.Literal, Token: p.curToken}

	p.NextToken()

	exp.Right = p.parseExpression(PREFIX)
	return exp

}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {

	program := &ast.Program{}

	program.Statements = []ast.Statement{}

	for !p.curTokenIs(tokens.EOF) {
		iter := p.parseStatement()
		if iter != nil {
			program.Statements = append(program.Statements, iter)
		}

		p.NextToken()

	}

	return program

}

func (p *Parser) parseStatement() ast.Statement {

	switch p.curToken.Type {
	case tokens.LET:
		return p.parseLetStatement()

	case tokens.RETURN:
		return p.parseReturnStatement()

	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(tokens.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) noPrefixParseFnError(t tokens.Token) {

	msg := fmt.Sprintf("no Prefix parse functions for %s found", t)
	p.errors = append(p.errors, msg)

}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.preFixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken)
	}

	leftExp := prefix()

	return leftExp

}

func (p *Parser) parseLetStatement() *ast.LetStatement {

	stmp := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	stmp.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(tokens.ASSIGN) {
		return nil
	}

	//TODO: skip expressions for now

	for !p.curTokenIs(tokens.SEMICOLON) {
		p.NextToken()
	}

	return stmp

}

func (p *Parser) curTokenIs(t tokens.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t tokens.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t tokens.TokenType) bool {

	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t tokens.TokenType) {
	msg := fmt.Sprintf("Expected Next Token to be %s, got %s instead", t, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.NextToken()

	for !(p.curTokenIs(tokens.SEMICOLON)) {
		p.NextToken()
	}

	return stmt
}
