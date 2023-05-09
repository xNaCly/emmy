package parser

import (
	"strings"

	"github.com/xNaCly/emmy/consts"
)

type Parser struct {
	input []consts.Token
	pos   int
}

func NewParser() *Parser {
	return &Parser{
		input: nil,
		pos:   0,
	}
}

func (p *Parser) NewInput(in []consts.Token) *Parser {
	p.input = in
	return p
}

func String(s []consts.Expression) string {
	b := strings.Builder{}
	for _, v := range s {
		if v != nil {
			b.WriteString(v.String())
		}
	}
	return b.String()
}

func (p *Parser) Parse() []consts.Expression {
	stmts := make([]consts.Expression, 0)
	for !p.isAtEnd() {
		stmts = append(stmts, p.statment())
	}
	return stmts
}

func (p *Parser) statment() consts.Expression {
	p.advance()
	return nil
}

func (p *Parser) peek() consts.Token {
	return p.input[p.pos]
}

func (p *Parser) prev() consts.Token {
	return p.input[p.pos-1]
}

func (p *Parser) advance() {
	if !p.isAtEnd() {
		p.pos++
	}
}

func (p *Parser) isAtEnd() bool {
	return p.pos >= len(p.input)
}
