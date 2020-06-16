package lexer

import (
	"../aux"
	"../aux/constant"
	"bufio"
	"log"
	"os"
	"strings"
)

var globalState uint8 = 0

// Entry point for the lexer
func Run(filename string) (*TokenList, aux.FoulError) {
	globalState = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	tokenList := TokenList{[]Token{}}
	transitionTable := getTransitionTable()
	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		line := scanner.Bytes()
		line = append(line, ' ')

		if tokenError := getTokensInLine(line, &tokenList, transitionTable); len(tokenError) > 0 {
			return &tokenList, aux.NewFoul(aux.UNKNOWN_TOKEN, i, tokenError)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if globalState == 7 {
		return &tokenList, aux.NewFoul(aux.UNEXPECTED_EOF, "/*")
	}

	return &tokenList, nil
}

// Find tokens in line and append them to the tokenList
func getTokensInLine(line []byte, tokenList *TokenList, transitionTable [][]uint8) string {
	var word []byte
	y := int(globalState)

	for i := 0; i < len(line); i++  {
		x := getIndexForChar(line[i])
		state := transitionTable[y][x]
		y = int(state)

		if state == 7 {
			globalState = state
		} else {
			globalState = 0
		}

		word = append(word, line[i])

		// Check if state is a terminal state
		if state >= 20 {
			if state == constant.S_ERROR {
				return strings.TrimSpace(string(word))
			}
			// add word and state into slice
			if state != constant.D_SPACE {
				tokenList.Add(NewToken(state, word[:len(word)-1]))
			}
			// clean word
			word = word[:0]
			y = 0
		}
		if isRecursiveToken(state) {
			i -= 1
		}
	}

	return ""
}

func GetFileLineForToken(filename string, target int) int {
	globalState = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	tokenList := TokenList{[]Token{}}
	transitionTable := getTransitionTable()
	scanner := bufio.NewScanner(file)

	counted := 0

	for i := 1; scanner.Scan(); i++ {
		line := scanner.Bytes()
		line = append(line, ' ')
		if tokenError := getTokensInLine(line, &tokenList, transitionTable); len(tokenError) > 0 {
			return -1
		}
		if target >= counted && target <= len(tokenList.Elements) {
			return i
		}
		counted = len(tokenList.Elements)
	}
	return -1
}

func getTransitionTable() [][]uint8 {
	return [][]uint8{
		{1, 2, constant.S_SUM, constant.S_SUBTRACT, constant.S_ASTERISK, 6, 3, 4, 5, constant.S_SEMICOLON, constant.S_COMMA, constant.S_OPEN_PARENTHESIS, constant.S_CLOSE_PARENTHESIS, constant.S_OPEN_SQR_BRACKET, constant.S_CLOSE_SQR_BRACKET, constant.S_OPEN_CURLY_BRACKET, constant.S_CLOSE_CURLY_BRACKET, 9, constant.D_SPACE, constant.S_ERROR},
		{1, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.H_WORD, constant.S_ERROR},
		{constant.H_NUMBER, 2, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER, constant.H_NUMBER},
		{constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS_EQUAL, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_LESS, constant.S_ERROR},
		{constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE_EQUAL, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_MORE, constant.S_ERROR},
		{constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_EQUAL, constant.S_ERROR},
		{constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, 8, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_FORWARD_SLASH, constant.S_ERROR},
		{7,7,7,7,8,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7},
		{7,7,7,7,7,0,7,7,7,7,7,7,7,7,7,7,7,7,7,7},
		{constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_NOT_EQUAL, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR, constant.S_ERROR},
	}
}

// Map a char to a column index
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
		} else {
			x = 19
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

func isRecursiveToken(state uint8) bool {
	if
		state == constant.H_WORD ||
		state == constant.H_NUMBER ||
		state == constant.S_FORWARD_SLASH ||
		state == constant.S_LESS ||
		state == constant.S_MORE ||
		state == constant.S_EQUAL {
		return true
	}
	return false
}
