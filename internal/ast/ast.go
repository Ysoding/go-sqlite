package ast

import "github.com/Ysoding/go-sqlite/internal/token"

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

type Program struct {
	Statements []Statement
}

type SelectStatement struct {
	Token   token.Token // the token.SELECT token
	Columns []Expression
	From    *FromClause
	Where   Expression
	OrderBy []*OrderClause
	Limit   *LimitClause
}

func (s *SelectStatement) statementNode() {}

func (s *SelectStatement) String() string {
	return ""
}

func (s *SelectStatement) TokenLiteral() string {
	return s.Token.Literal
}

type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type FromClause struct {
	Tables []Identifier
}

type OrderClause struct {
	Column    Expression
	Direction string // "asc" or "desc"
}

type LimitClause struct {
	Count  Expression
	Offset Expression
}

func (l *LimitClause) expressionNode() {}
