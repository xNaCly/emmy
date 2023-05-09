package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/xNaCly/emmy/consts"
)

type Parser struct {
	input       []consts.Token
	pos         int
	inputString string
	err         bool
}

func NewParser() *Parser {
	return &Parser{
		input: nil,
		pos:   0,
	}
}

func (p *Parser) NewInput(in []consts.Token, input string) *Parser {
	p.input = in
	p.pos = 0
	p.inputString = input
	return p
}

func String(s []consts.Expression) string {
	out, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Panicln(err)
	}
	return string(out)
}

func (p *Parser) Eval(s []consts.Expression) float64 {
	if p.err {
		log.Println("semantic errors found, no evaluation of ast possible")
		return math.NaN()
	} else {
		var res float64
		for _, s := range s {
			if s == nil {
				continue
			}
			res += s.Eval()
		}
		return res
	}
}

func (p *Parser) Parse() []consts.Expression {
	stmts := make([]consts.Expression, 0)
	stmts = append(stmts, p.statement())
	return stmts
}

func (p *Parser) statement() consts.Expression {
	return p.binaryStmt()
}

// TODO: needs to unwind to stop the repl
func (p *Parser) error() {
	p.err = true
	t := p.peek()
	tk := t.Kind
	var val any
	if t.Content != nil {
		val = t.Content
	} else {
		val = consts.TOKEN_LOOKUP[tk]
	}
	if tk == consts.EOF {
		log.Print("error: unexpected end of expression")
	} else {
		log.Printf("error: unexpected '%v' (%s) at position %d\n", val, consts.KIND_LOOKUP[tk], p.pos)
	}
	fmt.Printf("\n\t%s\n\t%s%s%s%s\n\n",
		p.inputString,
		strings.Repeat(" ", p.pos),
		"^",
		strings.Repeat(" ", p.pos),
		"unexpected token",
	)
}

func (p *Parser) unaryStmt() consts.Expression {
	if p.match(consts.NUMBER) {
		val, ok := p.prev().Content.(float64)
		if !ok {
			log.Fatalln("Token value not of type float64")
		}
		return consts.NumberExpression{Value: float64(val)}
	}
	p.error()
	return nil
}

func (p *Parser) binaryStmt() consts.Expression {
	return p.lineStatement()
}

func (p *Parser) pointStatement() consts.Expression {
	lhs := p.unaryStmt()

	for p.match(consts.DIVISION) {
		lhs = consts.DivisionExpression{
			Lhs: lhs,
			Rhs: p.unaryStmt(),
		}
	}
	for p.match(consts.MULTIPLICATION) {
		lhs = consts.MultiplicationExpression{
			Lhs: lhs,
			Rhs: p.unaryStmt(),
		}
	}

	return lhs
}

func (p *Parser) lineStatement() consts.Expression {
	lhs := p.pointStatement()
	for p.match(consts.PLUS) {
		lhs = consts.PlusExpression{
			Lhs: lhs,
			Rhs: p.pointStatement(),
		}
	}
	for p.match(consts.MINUS) {
		lhs = consts.MinusExpression{
			Lhs: lhs,
			Rhs: p.pointStatement(),
		}
	}

	return lhs
}

func (p *Parser) check(k int) bool {
	return p.input[p.pos].Kind == k
}

func (p *Parser) peek() consts.Token {
	if p.pos >= len(p.input) {
		return consts.Token{
			Pos:     len(p.input),
			Kind:    consts.EOF,
			Content: nil,
			Raw:     "",
		}
	}
	return p.input[p.pos]
}

func (p *Parser) match(t ...int) bool {
	for _, e := range t {
		if p.peek().Kind == e {
			p.advance()
			return true
		}
	}
	return false
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
