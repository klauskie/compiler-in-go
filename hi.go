package main

import (
	"./aux/constant"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		fmt.Printf("i: %d, t: %s, w: %s\n", i, list.elements[i].getTypeToString(), list.elements[i]._word )
	}
}

func (t *Token) getTypeToString() string {
	var stype string
	switch t._type {
	case constant.S_SUM:
		stype = "+"
		break
	case constant.S_SUBTRACT:
		stype = "-"
		break
	case constant.S_ASTERISK:
		stype = "*"
		break
	case constant.S_FORWARD_SLASH:
		stype = "/"
		break
	case constant.S_LESS:
		stype = "<"
		break
	case constant.S_LESS_EQUAL:
		stype = "<="
		break
	case constant.S_MORE:
		stype = ">"
		break
	case constant.S_MORE_EQUAL:
		stype = ">="
		break
	case constant.S_EQUAL_EQUAL:
		stype = "=="
		break
	case constant.S_NOT_EQUAL:
		stype = "!="
		break
	case constant.S_EQUAL:
		stype = "="
		break
	case constant.S_SEMICOLON:
		stype = ";"
		break
	case constant.S_COMMA:
		stype = ","
		break
	case constant.S_OPEN_PARENTHESIS:
		stype = "("
		break
	case constant.S_CLOSE_PARENTHESIS:
		stype = ")"
		break
	case constant.S_OPEN_SQR_BRACKET:
		stype = "["
		break
	case constant.S_CLOSE_SQR_BRACKET:
		stype = "]"
		break
	case constant.S_OPEN_CURLY_BRACKET:
		stype = "{"
		break
	case constant.S_CLOSE_CURLY_BRACKET:
		stype = "}"
		break
	case constant.S_OPEN_COMMENT_BLOCK:
		stype = "/*"
		break
	case constant.S_CLOSE_COMMENT_BLOCK:
		stype = "*/"
		break
	case constant.K_IF_ID:
		stype = "if"
		break
	case constant.K_ELSE_ID:
		stype = "else"
		break
	case constant.K_INT_ID:
		stype = "int"
		break
	case constant.K_RETURN_ID:
		stype = "return"
		break
	case constant.K_VOID_ID:
		stype = "void"
		break
	case constant.K_WHILE_ID:
		stype = "while"
		break
	case constant.K_INPUT_ID:
		stype = "input"
		break
	case constant.K_OUTPUT_ID:
		stype = "output"
		break
	case constant.H_WORD:
		stype = "word"
		break
	case constant.H_NUMBER:
		stype = "number"
		break
	case constant.D_SPACE:
		stype = "space"
		break
	case constant.D_NEWLINE:
		stype = "newline"
		break
	case constant.D_TAB:
		stype = "tab"
		break
	}
	return stype
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
		line = append(line, ' ')
		//callTokensRec(&tokenList, line)
		getTokensIter(line, &tokenList)
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
		{1, 2, constant.S_SUM, constant.S_SUBTRACT, 6, 7, 3, 4, 5, constant.S_SEMICOLON, constant.S_COMMA, constant.S_OPEN_PARENTHESIS, constant.S_CLOSE_PARENTHESIS, constant.S_OPEN_SQR_BRACKET, constant.S_CLOSE_SQR_BRACKET, constant.S_OPEN_CURLY_BRACKET, constant.S_CLOSE_CURLY_BRACKET, 8, constant.D_SPACE}, // S - 0 | A	B	SUM	SUBTRACT	F	G	D	D	E	SEMICOLON	COMA	OPEN_PARENTHESIS	CLOSE_PARENTHESIS	OPEN_SQUARE_BRACKET	CLOSE_SQUARE_BRACKET	OPEN_CURLY_BRACKET	CLOSE_CURLY_BRACKET	H	-
		{1, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD}, 														// A - 1 |
		{constant.H_NUMBER, 2, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER}, 					// B - 2 |
		{constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS_EQUAL, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS}, 							// C - 3 |
		{constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE_EQUAL, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE}, 							// D - 4 |
		{constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL}, 						// E - 5 |
		{constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_CLOSE_COMMENT_BLOCK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK, constant.S_ASTERISK}, 			// F - 6 |
		{constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_OPEN_COMMENT_BLOCK, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH}, 		// G - 7 |
		{0, 0, 0, 0, 0, 0, 0, 0, constant.S_NOT_EQUAL, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 								// H - 8 |
//		 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18
	}
}

func isKeyword(word string) (uint8, bool) {
	var id uint8
	isKeyword := true

	switch word {
	case constant.K_IF:
		id = constant.K_IF_ID
		break
	case constant.K_ELSE:
		id = constant.K_ELSE_ID
		break
	case constant.K_INT:
		id = constant.K_INT_ID
		break
	case constant.K_RETURN:
		id = constant.K_RETURN_ID
		break
	case constant.K_VOID:
		id = constant.K_VOID_ID
		break
	case constant.K_WHILE:
		id = constant.K_WHILE_ID
		break
	case constant.K_INPUT:
		id = constant.K_INPUT_ID
		break
	case constant.K_OUTPUT:
		id = constant.K_OUTPUT_ID
		break
	default:
		id = constant.H_WORD
		isKeyword = false
	}

	return id, isKeyword
}

func newToken(state uint8, word []byte) Token {
	var w string
	if state == constant.H_WORD || state == constant.H_NUMBER {
		w = strings.TrimSpace(string(word))
	} else {
		w = string(word[:0])
	}

	if state == constant.H_WORD {
		state, _ = isKeyword(w)
	}

	return Token{_type: state, _word: w}
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

func getTokensIter(line []byte, tokenList *TokenList) {
	var word []byte
	y := 0

	for i := 0; i < len(line); i++  {
		x := getIndexForChar(line[i])
		state := getTransitionTable()[y][x]
		y = int(state)
		word = append(word, line[i])

		if state >= 20 {
			// add word and state into slice
			if state != constant.D_SPACE {
				tokenList.add(newToken(state, word[:len(word)-1]))
			}
			// clean word
			word = word[:0]
			y = 0
		}

		if state == constant.H_WORD || state == constant.H_NUMBER {
			i -= 1
		}
	}
}

func getTokensRec(x int, y int, line []byte, index int, tokenList *TokenList, word []byte) {
	if index == len(line) {
		return
	}

	newX := -1
	if index+1 < len(line) {
		newX = getIndexForChar(line[index+1])
	}

	word = append(word, line[index])
	state := getTransitionTable()[y][x]
	fmt.Printf("out: %s | state: %d\n", word, state)

	if state >= 20 {
		// TODO : if state == D_SPACE then don' add it to the list
		newWord := word[len(word)-1:]
		tokenList.add(newToken(state, word[:len(word)-1]))
		word = newWord
		state = 0
	}
	getTokensRec(newX, int(state), line, index+1, tokenList, word)
}

func callTokensRec(tokenList *TokenList, b []byte) {
	var word []byte
	getTokensRec(getIndexForChar(b[0]), 0, b, 0, tokenList, word)
}
