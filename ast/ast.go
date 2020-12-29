package ast

import (
	"bytes"
	"interpreter/token"
)

// TokenLiteral is used only for debugging and testing
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// in monkey almost everything is an expression
// expressions are code segments that produce values
// let x = 5; 5 is an expression; let does not produce a value,
// it simply binds x to 5; thus, it is an statement;

// functions are expressions in monkey :: sublime

// Program consists of Statements
// Program is a concrete type that implements the Node interface
// Program is the root Node of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement utilizes the following structure:
// let <identifier>
// = <expression> <-> the expression that produces the value
type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

// dummy method implemented for::
// TypeErrors(when an expression is used as an statement)
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier represents the user defined Ident
// inside of the LetStatement -> to keep things simple and
// the amount of Node types controlled -> awesome
// it also implements Node interface, making it an expression
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value }

// return <expression>
type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// for standalone expressions that produce a value but
// lack state -> they're not persisted; kind of useless
// for a program :: like -> x + 5; 5 + 5; 10 + 12; etc
// since it's a statement it can be added to Program.Statements
// which is the whole reason for it being a Statement instead of an
// expression
type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}