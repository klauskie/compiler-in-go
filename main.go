package main

import (
	"./lexer"
	"log"
)

func main() {
	tokens, err := lexer.Run("file.txt")
	if err != nil {
		//panic(err.ToString())
		log.Fatal(err.ToString())
	}

	tokens.Print()
}




