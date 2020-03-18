package main

import (
	"./lexer"
	"log"
)

func main() {
	_, err := lexer.GetTokens("file.txt")
	if err != nil {
		//panic(err.ToString())
		log.Fatal(err.ToString())
	}
}




