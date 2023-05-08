package consts

type Token struct {
	Pos     int
	Kind    int
	Content any // either float64 or string, is this a codesmell?
	Raw     string
}

var BUILD_INS = map[string]int{
	"@sin":  SIN,
	"@cos":  COS,
	"@tan":  TAN,
	"@lb":   LB,
	"@ln":   LN,
	"@lg":   LG,
	"@pi":   PI,
	"@e":    E,
	"@phi":  PHI,
	"@sqrt": SQRT,
}

const (
	UNKNOWN    = iota + 1
	NUMBER     // 0-9
	IDENTIFIER // a-Z

	// single char
	PLUS           // +
	MINUS          // -
	DIVISION       // /
	MULTIPLICATION // *
	MODULUS        // %
	POWER          // ^
	COMMA          // ,
	SEMICOLON      // ;
	PARENOPEN      // (
	PARENCLOSE     // )
	BRACKETOPEN    // [
	BRACKETCLOSE   // ]
	BRACEOPEN      // {
	BRACECLOSE     // }
	COLON          // :
	EQUAL          // =

	// built in functions
	SIN
	SQRT // square root
	E    // natural number
	PHI
	PI // π
	COS
	TAN
	LB // log2
	LN // log e
	LG // log 10
)

var KIND_LOOKUP = map[int]string{
	UNKNOWN:        "UNKNOWN",
	COLON:          "COLON",
	EQUAL:          "EQUAL",
	NUMBER:         "NUMBER",
	PLUS:           "PLUS",
	MINUS:          "MINUS",
	DIVISION:       "DIVISION",
	MULTIPLICATION: "MULTIPLICATION",
	POWER:          "POWER",
	COMMA:          "COMMA",
	SEMICOLON:      "SEMICOLON",
	PARENOPEN:      "PARENOPEN",
	PARENCLOSE:     "PARENCLOSE",
	BRACKETOPEN:    "BRACKETOPEN",
	BRACKETCLOSE:   "BRACKETCLOSE",
	BRACEOPEN:      "BRACEOPEN",
	BRACECLOSE:     "BRACECLOSE",
	IDENTIFIER:     "IDENTIFIER",
	SIN:            "@SIN",
	SQRT:           "@SQRT",
	PI:             "@PI",
	E:              "@E",
	PHI:            "@PHI",
	COS:            "@COS",
	TAN:            "@TAN",
	LB:             "@LB",
	LN:             "@LN",
	LG:             "@LG",
}
