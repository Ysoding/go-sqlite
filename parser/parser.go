package parser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Parser struct {
	input        string
	position     int // start currenct char index
	readPosition int // after currenct char index
	ch           byte
}

func New(inptu string) *Parser {
	l := &Parser{
		input: inptu,
	}
	l.readChar()
	return l
}

func (p *Parser) Parse(out io.Writer) {
	p.skipWhitespace()

	if p.ch == '.' {
		if err := p.doMetaCommand(); err != nil {
			fmt.Fprintf(out, "%s\n", err.Error())
			return
		}
		return
	}

	statement, err := p.prepareStatement()
	if err != nil {
		fmt.Fprintf(out, "%s\n", err.Error())
		return
	}

	p.executeStatement(statement)
	fmt.Fprintf(out, "Executed.\n")
}

func (p *Parser) executeStatement(st *Statement) {

}

func (p *Parser) prepareStatement() (*Statement, error) {
	switch {
	case strings.HasPrefix(p.input[p.position:], "insert"):
		return &Statement{Type: StatementInsert}, nil
	case strings.HasPrefix(p.input[p.position:], "select"):
		return &Statement{Type: StatementSelect}, nil
	default:
		return nil, fmt.Errorf("unrecognized keyword at start of '%s'", p.input[p.position:])
	}
}

func (p *Parser) doMetaCommand() error {
	if p.input[p.position:] == ".exit" {
		os.Exit(0)
	} else {
		return fmt.Errorf("unrecognized command '%s'", p.input)
	}
	return nil
}

func (p *Parser) skipWhitespace() {
	for p.ch == ' ' || p.ch == '\t' || p.ch == '\n' || p.ch == '\r' {
		p.readChar()
	}
}

func (p *Parser) readChar() {
	if p.readPosition >= len(p.input) {
		p.ch = 0
	} else {
		p.ch = p.input[p.readPosition]
	}
	p.position = p.readPosition
	p.readPosition += 1
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
