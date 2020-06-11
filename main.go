package main

import (
	"flag"
	"fmt"
	"os"

	"stepwise.com/infix/convert"
)

func printHelp() {
	programName, _ := os.Executable()
	fmt.Printf("Usage: %s -e \"mathmatical expression\"\n", programName)
}

func parseArgs() string {
	expressionPtr := flag.String("e", "", "A mathmatical expression to be converted to infix notation")

	flag.Parse()

	expression := *expressionPtr
	if expression == "" {
		flag.PrintDefaults()
		return ""
	}

	return expression
}

func main() {
	fmt.Println("Received args: ", os.Args)
	expressionToConvert := parseArgs()
	if expressionToConvert == "" {
		os.Exit(1)
		return
	}

	fmt.Println(convert.Convert2(expressionToConvert))
}
