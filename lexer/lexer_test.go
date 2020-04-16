package lexer

import (
	"../aux/constant"
	"testing"
)

func TestRun(t *testing.T) {
	tokens, err := Run("../file_test.txt")
	if err == nil {
		t.Errorf("Run; Expected an error; got nil")
	}

	if len(tokens.Elements) != 13 {
		t.Errorf("Run; Expected 13 tokens detected; got %d", len(tokens.Elements))
	}

	for i, token := range tokens.Elements {
		switch i {
		case 0:
			if token.Type != constant.K_INT_ID {
				t.Errorf("Run; Expected an INT TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 1:
			if token.Type != constant.H_WORD {
				t.Errorf("Run; Expected a WORD TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 2:
			if token.Type != constant.S_OPEN_PARENTHESIS {
				t.Errorf("Run; Expected an OPEN PARENTHESIS TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 3:
			if token.Type != constant.S_CLOSE_PARENTHESIS {
				t.Errorf("Run; Expected a CLOSE PARENTHESIS TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 4:
			if token.Type != constant.S_OPEN_CURLY_BRACKET {
				t.Errorf("Run; Expected an OPEN CURLY BRACKET TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 5:
			if token.Type != constant.K_OUTPUT_ID {
				t.Errorf("Run; Expected an OUTPUT TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 6:
			if token.Type != constant.S_OPEN_SQR_BRACKET {
				t.Errorf("Run; Expected an OPEN SQUARE BRACKET TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 7:
			if token.Type != constant.H_NUMBER {
				t.Errorf("Run; Expected a NUMBER TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 8:
			if token.Type != constant.S_ASTERISK {
				t.Errorf("Run; Expected an ASTERISK TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 9:
			if token.Type != constant.H_WORD {
				t.Errorf("Run; Expected a WORD TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 10:
			if token.Type != constant.S_CLOSE_SQR_BRACKET {
				t.Errorf("Run; Expected a CLOSE SQUARE BRACKET TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 11:
			if token.Type != constant.K_INPUT_ID {
				t.Errorf("Run; Expected an INPUT TYPE; got %s", getTypeToString(token.Type))
			}
			break
		case 12:
			if token.Type != constant.S_OPEN_PARENTHESIS {
				t.Errorf("Run; Expected an OPEN PARENTHESIS TYPE; got %s", getTypeToString(token.Type))
			}
			break
		}
	}
}

func TestGetTokensInLine(t *testing.T) {

	line := "hello" + " "
	tl := TokenList{[]Token{}}
	tokenMock := Token{constant.H_WORD, "hello"}
	transitionTable := getTransitionTable()
	getTokensInLine([]byte(line), &tl, transitionTable)
	if tl.Elements[0] != tokenMock {
		t.Errorf("getTokensInLine; Expected a token with type %d and word = %s; " +
			"got type = %d and word = %s", tokenMock.Type, tokenMock.Word, tl.Elements[0].Type, tl.Elements[0].Word)
	}
	// Clean token list
	tl.Elements = tl.Elements[:0]

	line = "2+a if" + " "
	getTokensInLine([]byte(line), &tl, transitionTable)
	if len(tl.Elements) != 4 {
		t.Errorf("getTokensInLine; Expected 4 tokens; got %d", len(tl.Elements))
	}
	if tl.Elements[0].Type != constant.H_NUMBER {
		t.Errorf("getTokensInLine; Expected a NUMBER token; got %s", getTypeToString(tl.Elements[0].Type))
	}
	if tl.Elements[1].Type != constant.S_SUM {
		t.Errorf("getTokensInLine; Expected a SUM token; got %s", getTypeToString(tl.Elements[1].Type))
	}
	if tl.Elements[2].Type != constant.H_WORD {
		t.Errorf("getTokensInLine; Expected a WORD token; got %s", getTypeToString(tl.Elements[2].Type))
	}
	if tl.Elements[3].Type != constant.K_IF_ID {
		t.Errorf("getTokensInLine; Expected an IF token; got %s", getTypeToString(tl.Elements[3].Type))
	}

	line = "\n" + " "
	tl2 := TokenList{[]Token{}}
	err := getTokensInLine([]byte(line), &tl2, transitionTable)
	if len(tl2.Elements) != 0 {
		t.Errorf("getTokensInLine; Expected no tokens; got %d token(s)", len(tl2.Elements))
	}

	// Clean token list
	tl.Elements = tl.Elements[:0]

	// Error encounters
	line = "if (2+2) $ then" + " "
	err = getTokensInLine([]byte(line), &tl, transitionTable)
	if err != "$" {
		t.Errorf("getTokensInLine; Expected UNKNOWN $ token; got %s", err)
	}

	// Comments
	line = "/* %& UNKNOWN TOKEN-S **/" + " "
	possibleFutureSize := 2
	tl.Elements = tl.Elements[:0]
	err = getTokensInLine([]byte(line), &tl, transitionTable)
	if possibleFutureSize != len(tl.Elements) {
		t.Errorf("getTokensInLine; Expected text to be ignored and 2 tokens detected; got %d token detected", len(tl.Elements))
	}

	line = "/* int main () {}" + " "
	tl.Elements = tl.Elements[:0]
	err = getTokensInLine([]byte(line), &tl, transitionTable)
	if len(tl.Elements) != 1 {
		t.Errorf("getTokensInLine; Expected all text to be ignerd and identify 1 token; got %d token detected", len(tl.Elements))
	}

}

func TestIsLetter(t *testing.T) {
	for _, c := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYX" {
		got := isLetter(byte(c))
		if !got {
			t.Errorf("IsLetter; Expected a letter; got %c", byte(c))
		}
	}

	got := isLetter(byte('1'))
	if got {
		t.Errorf("IsLetter; Expected a number; got %c", byte('1'))
	}

}

func TestIsDigit(t *testing.T) {
	for _, c := range "0123456789" {
		got := isDigit(byte(c))
		if !got {
			t.Errorf("isDigit; Expected a digit; got %c", byte(c))
		}
	}

	got := isDigit(byte('a'))
	if got {
		t.Errorf("isDigit; Expected a letter; got %c", byte('1'))
	}

}

func TestGetIndexForChar(t *testing.T) {
	got := getIndexForChar(byte('+'))
	if got != 2 {
		t.Errorf("getIndexForChar; Expected 2 for [ + ] input; got %c", byte('+'))
	}

	got = getIndexForChar(byte('a'))
	if got != 0 {
		t.Errorf("getIndexForChar; Expected 0 for [ a ] input; got %c", byte('a'))
	}

	got = getIndexForChar(byte('1'))
	if got != 1 {
		t.Errorf("getIndexForChar; Expected 1 for [ 1 ] input; got %c", byte('1'))
	}

	got = getIndexForChar(byte('ñ'))
	if got != 19 {
		t.Errorf("getIndexForChar; Expected 19 for [ ñ ] input; got %c", byte('ñ'))
	}
}