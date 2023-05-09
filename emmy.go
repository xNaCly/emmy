package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xNaCly/emmy/consts"
	"github.com/xNaCly/emmy/lexer"
	"github.com/xNaCly/emmy/parser"
)

var DEBUG = false

func run(in string, scanner *lexer.Scanner, parser *parser.Parser) {
	t := scanner.NewInput(in).Start()
	if DEBUG {
		fmt.Println(lexer.String(t))
	}
	// parser.NewInput(t).Parse()
}

func main() {
	l := lexer.NewScanner()
	p := parser.NewParser()

	if len(os.Args) > 1 {
		run(os.Args[1], l, p)
		return
	}

	fmt.Println("Welcome to the emmy repl")
	prompt := "ε> "
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')

		if line[0] == '.' {
			l := strings.TrimSpace(line)
			switch l {
			case ".exit":
				fmt.Println("Exiting...")
				os.Exit(0)
			case ".help":
				fmt.Println(consts.HELP_MSG)
			case ".debug":
				DEBUG = !DEBUG
				fmt.Printf("Toggled debug mode to: '%v'\n", DEBUG)
			default:
				fmt.Printf("Unknown command: '%s'\n", l)
				continue
			}
		} else {
			run(line, l, p)
		}
	}
}
