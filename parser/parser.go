package parser

import "github.com/xNaCly/emmy/consts"

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

func (p *Parser) Parse() []consts.Expression {
	stmts := make([]consts.Expression, 0)
	for !p.isAtEnd() {
		stmts = append(stmts, p.statment())
	}
	return stmts
}

func (p *Parser) statment() consts.Expression {
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
