package main

import (
	"./lexer"
	"./parser"
	"fmt"
)

func main() {
	fmt.Println()
	phaseBuilder()
}

func phaseBuilder() {
	// Lexical Analysis
	tokenList, err := lexer.Run("eval_lex_1.txt")
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
	parserError := parser.RunParser(tokenList)
	if parserError != nil {
		fmt.Println(parserError.ToString())
		//log.Fatal(parserError.ToString())
		return
	}
}




