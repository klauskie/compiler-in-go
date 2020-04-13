package lexer

import (
	"../aux/constant"
	"testing"
)

func TestAdd(t *testing.T) {
	tokenList := TokenList{[]Token{}}

	token := Token{
		Type: 0,
		Word: "test",
	}

	tokenList.Add(token)


	if len(tokenList.Elements) == 0 {
		t.Errorf("Add(Token) Expected length: %d, got: %d", 1, len(tokenList.Elements))
	}

	if tokenList.Elements[0] != token {
		t.Errorf("Add(Token) Expected token at pos 0")
	}

}

func TestNewToken(t *testing.T) {
	t1 := NewToken(constant.S_EQUAL, []byte("="))
	if t1.Word == "=" {
		t.Errorf("Expected word = ''; got '%s' ", t1.Word)
	}

	t2 := NewToken(constant.H_WORD, []byte("test"))
	if t2.Word != "test" {
		t.Errorf("Expected word = 'test'; got '%s' ", t2.Word)
	}

	t3 := NewToken(constant.H_NUMBER, []byte("321"))
	if t3.Word != "321" {
		t.Errorf("Expected word = '321'; got '%s' ", t3.Word)
	}

	t4 := NewToken(constant.H_WORD, []byte("if"))
	if t4.Word == "if" {
		t.Errorf("Expected word = ''; got '%s' ", t4.Word)
	}
	if t4.Type != constant.K_IF_ID {
		t.Errorf("Expected type = '%d'; got '%d' ",constant.K_IF_ID, t4.Type)
	}
}
