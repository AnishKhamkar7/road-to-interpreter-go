package ast

import "go-int/src/tokens"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {

	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token tokens.Token
	Value Expression
	Name  *Identifier
}

func (Is *LetStatement) statementNode() {

}

func (Is *LetStatement) TokenLiteral() string {
	return Is.Token.Literal
}

type Identifier struct {
	Token tokens.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
