package lexer

import (
	"testing"

	"github.com/xNaCly/emmy/consts"
)

// iterates over key value pairs in m, calls testing.Errorf if the matcher function returns false for ok
func testHelper(t *testing.T, m map[string]any, matcher func(k string, v any) (bool, any)) {
	for k, v := range m {
		if ok, val := matcher(k, v); !ok {
			t.Errorf("wanted %v, got %v", v, val)
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
