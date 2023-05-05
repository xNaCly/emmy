package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xNaCly/emmy/consts"
	"github.com/xNaCly/emmy/lexer"
	"github.com/xNaCly/emmy/parser"
)

func run(in string, scanner *lexer.Scanner, parser *parser.Parser) {
	t := scanner.NewInput(os.Args[1]).Start()
	for _, v := range t {
		fmt.Printf("[%s][%d][%v]\n", consts.KIND_LOOKUP[v.Kind], v.Pos, v.Content)
	}
	parser.NewInput(t)
}

func main() {
	l := lexer.NewScanner()
	p := parser.NewParser()
	if len(os.Args) > 1 {
		run(os.Args[1], l, p)
		return
	}
	log.Println("welcome to the emmy repl")
	prompt := "ε> "
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		run(line, l, p)
	}
}
