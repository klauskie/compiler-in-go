package lexer

import (
	"../aux/constant"
	"fmt"
)

type Symbol struct {
	ID int
	Type uint8
	Word string
}

type SymbolTable struct {
	Map map[string]Symbol
}

func (table *SymbolTable) Insert(s Symbol) {
	table.Map[s.Word] = s
}

func (table *SymbolTable) Print() {
	for key, element := range table.Map {
		fmt.Printf("Key: %s => Element: %s\n", key, element.getTypeToString())
	}
}

func (t *Symbol) getTypeToString() string {
	var stype string
	switch t.Type {
	case constant.H_WORD:
		stype = "word"
		break
	case constant.H_NUMBER:
		stype = "number"
		break
	}
	return stype
}