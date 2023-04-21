package consts

type Token struct {
	pos     uint
	kind    uint
	content string
}

const (
	UNKNOWN        = iota + 1
	NUMBER         // 0-9
	PLUS           // +
	MINUS          // -
	DIVISION       // /
	MULTIPLICATION // *
	POWER          // ^
	COMMA          // ,
	SEMICOLON      // ;
	PARENOPEN      // (
	PARENCLOSE     // )
	IDENTIFIER     // a-Z
)
