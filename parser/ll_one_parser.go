package parser

import (
	"../aux/constant"
	"../lexer"
	"errors"
	"fmt"
)

var tokens []lexer.Token
var tokenIndex int = 0
var grammarRules GrammarRules
var tableMap map[uint8][]int

func RunParser(tokenList *lexer.TokenList) {
	initialSetup(tokenList)
}

func initialSetup(tokenList *lexer.TokenList) {
	tokens = tokenList.Elements
	tokens = append(tokens, lexer.Token{
		Type: constant.S_ENDLINE,
		Word: "$",
	})
	grammarRules = *initGrammarRules()
	tableMap, _ = FillParsingTable2("parsing_table_two.csv")
	tkStack := RuleStack{}

	tkStack.push(constant.S_ENDLINE)
	tkStack.push(constant.R_PROGRAM)

	if executeParse(&tkStack) {
		fmt.Println("Parsing Complete: SUCCESS")
	} else {
		fmt.Println("Parsing Complete: FAIL")
	}

}

func executeParse(tkStack *RuleStack) bool {
	tokenType, tErr := nextToken()

	for {
		if tkStack.top() == constant.S_ENDLINE {
			fmt.Println("EMPTY STACK")
			break
		}

		fmt.Printf("TopStack: %d, CurrToken: %d | %s\n", tkStack.top(), tokenType, lexer.GetTypeToString(tokenType))

		if tkStack.top() == tokenType {
			tkStack.pop()
			tokenType, tErr = nextToken()
			if tErr != nil {
				fmt.Println("Error nextToken")
				break
			}
		} else if tkStack.top() < 60 {
			// TODO error mismatch
			fmt.Println("Error token mismatch")
			return false
		} else if tableMap[tkStack.top()][GetTokenColumn(tokenType)] == 0 {
			// TODO error dont correspond
			fmt.Printf("Error dont correspond: prod: %d, col: %d\n", tkStack.top(), GetTokenColumn(tokenType))
			return false
		} else if pRule := tableMap[tkStack.top()][GetTokenColumn(tokenType)]; pRule > 0 {
			tkStack.pop()
			pushRulesToStack(tkStack, grammarRules.getRuleByID(uint8(pRule)))
		}
	}

	if tkStack.top() == constant.S_ENDLINE && tokenType == constant.S_ENDLINE {
		return true
	}
	return false

}


func pushRulesToStack(tkStack *RuleStack, rule ProductionRule) {
	fmt.Printf("Push to stack! FATHER: %s\n", rule.Name)
	size := len(rule.RuleList)
	for i := size-1; i >= 0; i-- {
		fmt.Printf("Rule %d, ", rule.RuleList[i])
		tkStack.push(rule.RuleList[i])
	}
	fmt.Println()
}

func nextToken() (uint8, error) {
	if tokenIndex > len(tokens) {
		return 0, errors.New("index out of range")
	}
	tk := tokens[tokenIndex]
	tokenIndex += 1
	return tk.Type, nil
}


/* Token Stack */
type RuleStack struct {
	collection []uint8
}

func (s *RuleStack) push(element uint8) {
	s.collection = append(s.collection, element)
}

func (s *RuleStack) pop()  {
	if len(s.collection) == 0 {
		return
	}
	s.collection = s.collection[:len(s.collection)-1]
}

func (s *RuleStack) top() uint8 {
	return s.collection[len(s.collection)-1]
}

func (s *RuleStack) isEmpty() bool {
	return len(s.collection) == 0
}
/* END Token Stack */
