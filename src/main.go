package main

import (
	"fmt"
	"flag"
)

var ps1, ps2 string

func main() {
	var language string
	var interactive, ourInteractive bool
	var code string
	var quiet bool
	var exp string
	

	flag.StringVar(&language, "l", "python2", "langauge")
	flag.StringVar(&code, "c", "", "code to run")
	flag.StringVar(&exp, "e", "", "expression to print")

	flag.StringVar(&ps1, "ps1", "--> ", "PS1")
	flag.StringVar(&ps2, "ps2", "... ", "PS2")

	flag.BoolVar(&interactive, "i", false, "interactive")
	flag.BoolVar(&ourInteractive, "I", false, "like -i, but never use language REPL")
	flag.BoolVar(&quiet, "q", false, "quiet")
	flag.Parse()

	args := flag.Args()


	lang := GetLanguage(language)
	if !quiet {
		fmt.Println(lang.Version())
	}
	if code != "" {
		lang.Eval(code)
	} 
	if exp != "" {
		lang.EvalAndTryToPrint(exp)
	}
	if len(args) > 0 {
		lang.EvalFile(args[0], args[1:])
	}
	if interactive {
		lang.REPL()
	} else if ourInteractive {
		lang.InternalREPL()
	}

	
}

