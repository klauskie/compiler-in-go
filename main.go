package main

import (
	"./lexer"
	"fmt"
	"log"
)

func main() {
	tokenList, err := lexer.Run("file.txt")
	if err != nil {
		//panic(err.ToString())
		tokenList.Print()
		log.Fatal(err.ToString())
	}

	symbolTable := lexer.NewSymbolTable()
	symbolTable.Fill(tokenList)

	fmt.Println("---- TOKEN LIST -----")
	tokenList.Print()
	fmt.Println("---- SYMBOL TABLE -----")
	symbolTable.Print()
}




