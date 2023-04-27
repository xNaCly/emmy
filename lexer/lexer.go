package lexer

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/xNaCly/emmy/consts"
)

type Scanner struct {
	in []rune // input string
	p  int    // current position in input
	cc rune   // current character in input
}

func NewScanner() *Scanner {
	return &Scanner{
		p:  0,
		in: nil,
		cc: 0,
	}
}

func (s *Scanner) NewInput(input string) *Scanner {
	s.p = 0
	s.in = []rune(input)
	if len(input) == 0 {
		s.cc = 0
	} else {
		s.cc = s.in[s.p]
	}
	return s
}

func (s *Scanner) advance() {
	if !s.isAtEnd() && s.p+1 < len(s.in) {
		s.p++
		s.cc = s.in[s.p]
	} else {
		s.cc = 0
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.p == len(s.in) || s.cc == 0
}

func (s *Scanner) error(lexem string, msg string) {
	pos := s.p - len(lexem)
	if pos < 0 {
		pos = 0
	} else {
		pos += 1
	}

	arrows := len(lexem)

	log.Printf("error: unexpected '%s' at position %d", lexem, pos)
	fmt.Printf("\n\t%s\n\t%s%s\n\t%s%s\n\n",
		string(s.in),
		strings.Repeat(" ", pos),
		strings.Repeat("^", arrows),
		strings.Repeat(" ", pos),
		msg,
	)
	s.cc = 0
}

// returns all runes matching the matcher function as a string, returns a found type
func (s *Scanner) matchWhile(matcher func(rune) bool) string {
	b := strings.Builder{}
	for matcher(s.cc) {
		b.WriteRune(s.cc)
		s.advance()
	}
	return b.String()
}

func (s *Scanner) Start() []consts.Token {
	if len(s.in) == 0 {
		return []consts.Token{}
	}
	token := make([]consts.Token, 0)
	for !s.isAtEnd() {
		var kind int
		var val any
		var raw string
		switch s.cc {
		case '+':
			kind = consts.PLUS
		case '-':
			kind = consts.MINUS
		case '/':
			kind = consts.DIVISION
		case '*':
			kind = consts.MULTIPLICATION
		case '%':
			kind = consts.MODULUS
		case '^':
			kind = consts.POWER
		case ',':
			kind = consts.COMMA
		case ';':
			kind = consts.SEMICOLON
		case '(':
			kind = consts.PARENOPEN
		case ')':
			kind = consts.PARENCLOSE
		case '[':
			kind = consts.BRACKETOPEN
		case ']':
			kind = consts.BRACKETCLOSE
		case '{':
			kind = consts.BRACEOPEN
		case '}':
			kind = consts.BRACECLOSE
		default:
			if unicode.IsDigit(s.cc) {
				kind = consts.NUMBER

				raw = s.matchWhile(func(r rune) bool {
					return unicode.IsDigit(r) || r == '.' || r == 'e' || r == '_'
				})

				t, err := strconv.ParseFloat(raw, 64)
				if err != nil {
					s.error(raw, "not a float: "+err.Error())
				}
				val = t
			} else if unicode.IsLetter(s.cc) || s.cc == '@' {
				val = s.matchWhile(func(r rune) bool {
					return unicode.IsLetter(s.cc) || s.cc == '@'
				})
				raw = val.(string)

				if res, ok := consts.BUILD_INS[val.(string)]; ok {
					kind = res
				} else {
					s.error(val.(string), "unknown identifier, view https://github.com/xnacly/emmy for the complete reference")
					return []consts.Token{}
				}
			} else {
				s.error(string(s.cc), "unexpected character")
				return []consts.Token{}
			}
		}

		// TODO: fix position calculation
		pos := s.p

		token = append(token, consts.Token{
			Pos:     pos,
			Kind:    kind,
			Content: val,
		})

		if kind != consts.IDENTIFIER && kind != consts.NUMBER {
			s.advance()
		}
	}
	return token
}
