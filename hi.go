package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	K_IF uint8 = iota + 10
	K_ELSE
	K_INT
	K_RETURN
	K_VOID
	K_WHILE
	K_INPUT
	K_OUTPUT
	S_SUM
	S_SUBTRACT
	S_ASTERISK // 20
	S_FORWARD_SLASH
	S_LESS
	S_LESS_EQUAL
	S_MORE
	S_MORE_EQUAL
	S_EQUAL_EQUAL
	S_NOT_EQUAL
	S_EQUAL
	S_SEMICOLON
	S_COMMA // 30
	S_OPEN_PARENTHESIS
	S_CLOSE_PARENTHESIS
	S_OPEN_SQR_BRACKET
	S_CLOSE_SQR_BRACKET
	S_OPEN_CURLY_BRACKET
	S_CLOSE_CURLY_BRACKET
	S_OPEN_COMMENT_BLOCK
	S_CLOSE_COMMENT_BLOCK
	H_WORD
	H_NUMBER // 40
	D_SPACE
	D_NEWLINE
	D_TAB
)

type Token struct {
	_type uint8
	_word string
}

type TokenList struct {
	elements []Token
}

func (list *TokenList) add(t Token) {
	list.elements = append(list.elements, t)
}

func (list *TokenList) print() {
	for i:= 0; i < len(list.elements); i++ {
		fmt.Printf("i: %d, t: %d, w: %s\n", i, list.elements[i]._type, list.elements[i]._word )
	}
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tokenList := TokenList{elements: []Token{}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		callTokensRec(&tokenList, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tokenList.print()
}

func isLetter(c byte) bool {
	if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
		return false
	}
	return true
}

func isDigit(c byte) bool {
	if c < '0' || c > '9' {
		return false
	}
	return true
}

func getTransitionTable() [][]uint8 {
	return [][]uint8{
		{1, 2, S_SUM, S_SUBTRACT, 6, 7, 3, 4, 5, S_SEMICOLON, S_COMMA, S_OPEN_PARENTHESIS, S_CLOSE_PARENTHESIS, S_OPEN_SQR_BRACKET, S_CLOSE_SQR_BRACKET, S_OPEN_CURLY_BRACKET, S_CLOSE_CURLY_BRACKET, 8, D_SPACE}, // S - 0 | A	B	SUM	SUBTRACT	F	G	D	D	E	SEMICOLON	COMA	OPEN_PARENTHESIS	CLOSE_PARENTHESIS	OPEN_SQUARE_BRACKET	CLOSE_SQUARE_BRACKET	OPEN_CURLY_BRACKET	CLOSE_CURLY_BRACKET	H	-
		{1, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD, H_WORD}, 														// A - 1 |
		{H_NUMBER, 2, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER, H_NUMBER}, 					// B - 2 |
		{S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS_EQUAL, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS, S_LESS}, 							// C - 3 |
		{0, 0, 0, 0, 0, 0, 0, 0, S_MORE_EQUAL, 0, 0, 0, 0, 0, 0, 0, 0, 0, S_MORE}, 							// D - 4 |
		{0, 0, 0, 0, 0, 0, 0, 0, S_EQUAL_EQUAL, 0, 0, 0, 0, 0, 0, 0, 0, 0, S_EQUAL}, 						// E - 5 |
		{0, 0, 0, 0, 0, S_CLOSE_COMMENT_BLOCK, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, S_ASTERISK}, 			// F - 6 |
		{0, 0, 0, 0, S_OPEN_COMMENT_BLOCK, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, S_FORWARD_SLASH}, 		// G - 7 |
		{0, 0, 0, 0, 0, 0, 0, 0, S_NOT_EQUAL, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 								// H - 8 |
//		 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18
	}
}

func getIndexForChar(c byte) int {
	x := 0
	switch c {
	case '+':
		x = 2
		break
	case '-':
		x = 3
		break
	case '*':
		x = 4
		break
	case '/':
		x = 5
		break
	case '<':
		x = 6
		break
	case '>':
		x = 7
		break
	case '=':
		x = 8
		break
	case ';':
		x = 9
		break
	case ',':
		x = 10
		break
	case '(':
		x = 11
		break
	case ')':
		x = 12
		break
	case '[':
		x = 13
		break
	case ']':
		x = 14
		break
	case '{':
		x = 15
		break
	case '}':
		x = 16
		break
	case '!':
		x = 17
		break
	case ' ':
		x = 18
		break
	case '\n':
		x = 18
		break
	default:
		if isLetter(c) {
			x = 0
		} else if isDigit(c) {
			x = 1
		}
		break
	}
	return x
}

func getTokensRec(x int, y int, line []byte, index int, tokenList *TokenList, word []byte) {
	if index == len(line) {
		return
	}

	newX := 0
	if index+1 < len(line) {
		newX = getIndexForChar(line[index+1])
	}

	word = append(word, line[index])
	state := getTransitionTable()[y][x]

	if state >= 10 {
		// TODO : if state == D_SPACE then don' add it to the list
		tokenList.add(Token{_type: state, _word: string(word)})
		word = word[:0]
		state = 0
	}
	getTokensRec(newX, int(state), line, index+1, tokenList, word)
}

func callTokensRec(tokenList *TokenList, b []byte) {
	word := []byte{}
	getTokensRec(getIndexForChar(b[0]), 0, b, 0, tokenList, word)
}
