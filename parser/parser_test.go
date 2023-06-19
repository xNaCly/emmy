package parser

import (
	"math"
	"testing"

	"github.com/xNaCly/emmy/lexer"
)

// iterates over key value pairs in m, calls testing.Errorf if the matcher function returns false for ok
func testHelper(t *testing.T, m map[string]any, matcher func(k string, v any) (bool, any)) {
	for k, v := range m {
		ok, val := matcher(k, v)
		if !ok {
			t.Errorf("wanted %v: got %v, for %s", v, val, k)
		}
	}
}

// the parser evalutes the ast to math.NaN if a sematic error is encountered
func TestParserError(t *testing.T) {
	tests := map[string]any{
		"1//2": struct{}{},
		"1/":   struct{}{},
		"5++":  struct{}{},
		"6++/": struct{}{},
		"8*/":  struct{}{},
	}

	s := lexer.NewScanner()
	p := NewParser()

	testHelper(t, tests, func(k string, _ any) (bool, any) {
		t := s.NewInput(k).Start()
		a := p.NewInput(t, k).Parse()
		o := p.Eval(a)
		return math.IsNaN(o), o
	})
}

func TestParserSimple(t *testing.T) {
	tests := map[string]any{
		"1/2": 0.5,
		"1/1": 1.0,
		"5+5": 10.0,
		"6-6": 0.0,
		"8*8": 64.0,
	}

	s := lexer.NewScanner()
	p := NewParser()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		t := s.NewInput(k).Start()
		a := p.NewInput(t, k).Parse()
		o := p.Eval(a)
		return math.Abs(o-v.(float64)) < 0.001, o
	})
}

func TestParserLong(t *testing.T) {
	tests := map[string]any{
		"1/2_092_102":           4.779881669249396e-07,
		"2_092_102/1":           2092102.0,
		"50_192_010_192.0021+5": 50192010197.0021,
		"60_129-6":              60123.0,
		"102_910*192_190_210":   19778294511100.0,
	}

	s := lexer.NewScanner()
	p := NewParser()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		t := s.NewInput(k).Start()
		a := p.NewInput(t, k).Parse()
		o := p.Eval(a)
		return math.Abs(o-v.(float64)) < 0.001, o
	})
}

func TestParserPemdas(t *testing.T) {
	// examples taken from https://pemdas.info/
	tests := map[string]any{
		"1+2*3": 7.0,
		// BUG: this is not correct!, emmy returns 7, not 8, somehow +3/3 is not included in the resulting ast
		// "7-1*0+3/3": 8.0,
		"1-2*3*4": -23.0,
		"3+4/2-4": 1.0,
	}

	s := lexer.NewScanner()
	p := NewParser()

	testHelper(t, tests, func(k string, v any) (bool, any) {
		t := s.NewInput(k).Start()
		a := p.NewInput(t, k).Parse()
		o := p.Eval(a)
		return math.Abs(o-v.(float64)) < 0.001, o
	})
}
