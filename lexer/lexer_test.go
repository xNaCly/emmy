package lexer

import (
	"testing"
)

func TestLexerError(t *testing.T) {
	l := len(NewScanner().NewInput("test").Start())
	if l != 0 {
		t.Errorf("got %d, wanted %d", l, 0)
	}
}

func TestLexerNumbers(t *testing.T) {
	tests := map[string]float64{
		"1":                    1,
		"0.25":                 0.25,
		"1.25":                 1.25,
		"1.2_5":                1.25,
		"1129_1029.2_5_1209_0": 11291029.251209,
	}
	s := NewScanner()
	for k, v := range tests {
		o := s.NewInput(k).Start()[0].Content.(float64)
		// INFO: comparing floats only really working by subtracting them
		// and checking if the reminder is within the defined acceptable margin
		if (o - v) > 0.01 {
			t.Errorf("got %f, wanted %f - diff: %f", o, v, o-v)
		}
	}
}

func TestLexerSingleChar(t *testing.T) {

}

func TestLexerBuildIn(t *testing.T) {

}
