package lexer

import "testing"

func TestSymbolTable_Insert(t *testing.T) {
	testSymbol := Symbol{
		ID:   0,
		Type: 0,
		Word: "",
	}
	symbolTable := NewSymbolTable()
	symbolTable.Insert(testSymbol)
	if len(symbolTable.Map) != 0 {
		t.Errorf("Insert(Symbol); Expected an empty Map, got Map with length of %d", len(symbolTable.Map))
	}

	testSymbol = Symbol{
		ID:   0,
		Type: 0,
		Word: "TEST",
	}
	symbolTable.Insert(testSymbol)
	if symbolTable.Map["TEST"] != testSymbol {
		t.Errorf("Insert(Symbol); Expected testSymbol at key=TEST, got %s", testSymbol.Word)
	}

	testSymbol = Symbol{
		ID:   0,
		Type: 1,
		Word: "TEST",
	}
	symbolTable.Insert(testSymbol)
	if symbolTable.Map["TEST"].Type != uint8(0) {
		t.Errorf("Insert(Symbol); Expected symbol at TEST to be of type 0, got %d", symbolTable.Map["TEST"].Type)
	}

}

func TestSymbolTable_Fill(t *testing.T) {
	symbolTable := NewSymbolTable()
	tokenList, _ := Run("../file_test.txt")

	symbolTable.Fill(tokenList)

	if len(symbolTable.Map) != 3 {
		t.Errorf("Insert(Symbol); Expected Map of length 3, got Map with length of %d", len(symbolTable.Map))
	}
}
