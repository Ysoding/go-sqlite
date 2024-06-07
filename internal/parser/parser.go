package parser

import (
	"github.com/Ysoding/go-sqlite/internal/ast"
	"github.com/Ysoding/go-sqlite/internal/lexer"
	"github.com/Ysoding/go-sqlite/internal/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	return p
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.SELECT:
		return p.parseSelectStatement()
	default:
		return nil
	}
}

func (p *Parser) parseSelectStatement() ast.Statement {
	// select * from st;
	stmt := &ast.SelectStatement{Token: p.curToken}

	return stmt
}

func (p *Parser) curTokenIs(typ token.Type) bool {
	return p.curToken.Type == typ
}

func (p *Parser) peekTokenIs(typ token.Type) bool {
	return p.peekToken.Type == typ
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
