package scanner

type Scanner struct {
	pos   uint
	input []rune
	cc    rune
}

func NewScanner(input string) *Scanner {
	in := []rune(input)
	return &Scanner{
		pos:   0,
		input: in,
		cc:    in[0],
	}
}

func (s *Scanner) NewInput(input string) {
	s.input = []rune(input)
}
