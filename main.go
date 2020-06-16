package main

import (
	"./lexer"
	"./parser"
	"fmt"
)

func main() {
	fmt.Println()
	phaseBuilder("eval_lex_1.txt")
}

func phaseBuilder(filename string) {
	// Lexical Analysis
	tokenList, err := lexer.Run(filename)
	if err != nil {
		//panic(err.ToString())
		//tokenList.Print()
		fmt.Println(err.ToString())
		//log.Fatal(err.ToString())
		return
	}

	symbolTable := lexer.NewSymbolTable()
	symbolTable.Fill(tokenList)

	// Syntax Analysis
	parserError := parser.RunParser(filename, tokenList)
	if parserError != nil {
		fmt.Println(parserError.ToString())
		//log.Fatal(parserError.ToString())
		return
	}
}




