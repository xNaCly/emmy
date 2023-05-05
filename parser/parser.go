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

func (p *Parser) parse() {}
