package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xNaCly/emmy/consts"
	"github.com/xNaCly/emmy/lexer"
)

func debug_printToken(t []consts.Token) {
	for _, v := range t {
		fmt.Printf("[%s][%d][%v]\n", consts.KIND_LOOKUP[v.Kind], v.Pos, v.Content)
	}
}

func main() {
	if len(os.Args) > 1 {
		debug_printToken(lexer.NewScanner().NewInput(os.Args[1]).Start())
		return
	}
	log.Println("welcome to the emmy repl")
	prompt := "ε> "
	reader := bufio.NewReader(os.Stdin)
	l := lexer.NewScanner()
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		debug_printToken(l.NewInput(strings.TrimSpace(line)).Start())
	}
}
