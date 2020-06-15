package main

import (
	"./lexer"
	"./parser"
	"log"
)

func main() {
	phaseBuilder()
}

func phaseBuilder() {
	// Lexical Analysis
	tokenList, err := lexer.Run("eval_lex_1.txt")
	if err != nil {
		//panic(err.ToString())
		tokenList.Print()
		log.Fatal(err.ToString())
	}

	symbolTable := lexer.NewSymbolTable()
	symbolTable.Fill(tokenList)

	// Syntax Analysis
	parser.RunParser(tokenList)
}




