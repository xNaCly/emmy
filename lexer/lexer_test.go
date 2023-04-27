package lexer

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/xNaCly/emmy/consts"
)

// iterates over key value pairs in m, calls testing.Errorf if the matcher function returns false for ok
func testHelper(t *testing.T, m map[string]any, matcher func(k string, v any) (bool, any)) {
	for k, v := range m {
		if ok, val := matcher(k, v); !ok {
			t.Errorf("wanted %v: got %v", v, val)
		}
	}
}

func TestLexerError(t *testing.T) {
	tests := map[string]any{
		"test":     0,
		"!":        0,
		"=":        0,
		"x":        0,
		"aoadojad": 0,
	}

	s := NewScanner()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		o := len(s.NewInput(k).Start())
		return o == v, o
	})
}

func TestLexerNumbers(t *testing.T) {
	tests := map[string]any{
		"1":                    1.0,
		"0.25":                 0.25,
		"1.25":                 1.25,
		"1.2_5":                1.25,
		"1129_1029.2_5_1209_0": 11291029.251209,
	}

	s := NewScanner()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		o := s.NewInput(k).Start()[0].Content.(float64)
		return (o - v.(float64)) < 0.01, o
	})
}

func TestLexerPositions(t *testing.T) {
	tests := map[string]any{
		"1+1":                        "0,1,2",
		"+0.25-4":                    "0,1,5,6",
		"@sqrt+129_0-":               "0,5,6,11",
		"12091_102910.0129-+9128791": "0,17,18,19",
	}

	s := NewScanner()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		o := s.NewInput(k).Start()
		y := true
		a := make([]int, 0)

		for _, val := range strings.Split(v.(string), ",") {
			i, _ := strconv.ParseInt(val, 10, 64)
			a = append(a, int(i))
		}

		r := make([]string, 0)
		for _, token := range o {
			r = append(r, fmt.Sprint(token.Pos))
		}

		for i, token := range o {
			if token.Pos != a[i] {
				log.Printf("error: %+v.Kind, does not match %d", token, a[i])
				y = false
				break
			}
		}

		return y, strings.Join(r, ",")
	})
}

func TestLexerSingleChar(t *testing.T) {
	tests := map[string]any{
		"0": consts.NUMBER,
		"+": consts.PLUS,
		"-": consts.MINUS,
		"/": consts.DIVISION,
		"*": consts.MULTIPLICATION,
		"^": consts.POWER,
		",": consts.COMMA,
		";": consts.SEMICOLON,
		"(": consts.PARENOPEN,
		")": consts.PARENCLOSE,
		"[": consts.BRACKETOPEN,
		"]": consts.BRACKETCLOSE,
		"{": consts.BRACEOPEN,
		"}": consts.BRACECLOSE,
	}

	s := NewScanner()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		o := s.NewInput(k).Start()[0].Kind
		return o == v, o
	})
}

func TestLexerBuildIn(t *testing.T) {
	tests := map[string]any{
		"@sin":  consts.SIN,
		"@sqrt": consts.SQRT,
		"@e":    consts.E,
		"@phi":  consts.PHI,
		"@pi":   consts.PI,
		"@cos":  consts.COS,
		"@tan":  consts.TAN,
		"@lb":   consts.LB,
		"@ln":   consts.LN,
		"@lg":   consts.LG,
	}

	s := NewScanner()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		o := s.NewInput(k).Start()[0].Kind
		return o == v, o
	})
}

func FuzzLexer(f *testing.F) {
	tests := []string{"129.0", "12_9128.9182_1", "@sin", "@sqrt", "()"}
	for _, tc := range tests {
		f.Add(tc)
	}
	s := NewScanner()
	f.Fuzz(func(t *testing.T, a string) {
		s.NewInput(a).Start()
	})
}
