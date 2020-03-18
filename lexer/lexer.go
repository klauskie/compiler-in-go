package lexer

import (
	"../aux/constant"
	"bufio"
	"log"
	"os"
	"strings"
)

func GetTokens(filename string) TokenList {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tokenList := TokenList{[]Token{}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		line = append(line, ' ')
		getTokensInLine(line, &tokenList)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tokenList.Print()
	return tokenList
}

func getTokensInLine(line []byte, tokenList *TokenList) {
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
				tokenList.Add(newToken(state, word[:len(word)-1]))
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

	return Token{state, w}
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