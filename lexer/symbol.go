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

type SymbolTableNumbers struct {
	index int
	Elements []Symbol
}

type SymbolTableStrings struct {
	index int
	Elements []Symbol
}

type SymbolTable interface {
	Add(symbol Symbol)
}

func (list *SymbolTableNumbers) Add(s Symbol) {
	list.index += 1
	s.ID = list.index
	list.Elements = append(list.Elements, s)
}

func (list *SymbolTableStrings) Add(s Symbol) {
	list.index += 1
	s.ID = list.index
	list.Elements = append(list.Elements, s)
}

func (list *SymbolTableNumbers) Print() {
	for i:= 0; i < len(list.Elements); i++ {
		fmt.Printf("i: %d| id: %d, t: %s, w: %s\n", i, list.Elements[i].ID, list.Elements[i].getTypeToString(), list.Elements[i].Word )
	}
}

func (list *SymbolTableStrings) Print() {
	for i:= 0; i < len(list.Elements); i++ {
		fmt.Printf("i: %d| id: %d, t: %s, w: %s\n", i, list.Elements[i].ID, list.Elements[i].getTypeToString(), list.Elements[i].Word )
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