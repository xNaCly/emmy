package consts

import (
	"fmt"
	"strings"
)

var commands = [][]string{
	{".exit", "quits the repl"},
	{".debug", "toggles debug mode"},
	{".help", "prints this message"},
}

var HELP_MSG = func() string {
	b := strings.Builder{}
	for _, c := range commands {
		b.WriteString(fmt.Sprintf("%-10s %-10s\n", c[0], c[1]))
	}
	return b.String()
}()
