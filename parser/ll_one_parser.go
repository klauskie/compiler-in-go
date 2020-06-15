package parser

import (
	"../aux/constant"
	"../lexer"
	"errors"
	"fmt"
	"log"
)

var tokens []lexer.Token
var tokenIndex int = 0
var grammarRules GrammarRules

func RunParser(tokenList *lexer.TokenList) {
	initialSetup(tokenList)
	//tests(tokenList)
}

func tests(tokenList *lexer.TokenList) {
	tokens = tokenList.Elements
	tokens = append(tokens, lexer.Token{
		Type: constant.S_ENDLINE,
		Word: "$",
	})
	grammarRules = *initGrammarRules()
	tkStack := RuleStack{}

	tkStack.push(grammarRules.getRuleByID(0)) // $
	tkStack.push(grammarRules.getRuleByID(1))

	// tests
	fmt.Println(tkStack.top())
	tkStack.pop()
	fmt.Println(tkStack.top())
}

func initialSetup(tokenList *lexer.TokenList) {
	tokens = tokenList.Elements
	tokens = append(tokens, lexer.Token{
		Type: constant.S_ENDLINE,
		Word: "$",
	})
	grammarRules = *initGrammarRules()
	tableMap, e := FillParsingTable("parsing_table.csv")
	if e != nil {
		log.Fatal(e)
	}
	tkStack := RuleStack{}

	tkStack.push(grammarRules.getRuleByID(0)) // $
	tkStack.push(grammarRules.getRuleByID(1)) // program RULE

	if executeParse(&tkStack, &tableMap) {
		fmt.Println("Parsing Complete: SUCCESS")
	} else {
		fmt.Println("Parsing Complete: FAIL")
	}

}

func executeParse(tkStack *RuleStack, tableMap *map[string][]int) bool {
	tokenType, tErr := nextToken()
	for {
		if tkStack.isEmpty() {
			return false
		}
		if tkStack.top().GetType() == constant.S_ENDLINE {
			break
		}

		fmt.Printf("TopStack: %d, Current Token: %d\n", tkStack.top().GetType(), tokenType)

		if tkStack.top().GetType() == tokenType {
			tkStack.pop()
			tokenType, tErr = nextToken()
			if tErr != nil {
				// TODO return error
				fmt.Println("ERROR. Next Token error")
				break
			}
		} else if tkStack.top().IsTerminal() {
			// TODO return error
			fmt.Println("ERROR. Token mismatch")
			break
		} else if getTableCell(*tableMap, tkStack, tokenType) == 0 {
			// TODO return error
			fmt.Println("ERROR. Wrong Table Index")
			break
		} else if prodID := getTableCell(*tableMap, tkStack, tokenType); prodID > 0 {
			pushRulesToStack(tkStack, grammarRules.getRuleByID(prodID))
		}
	}

	// Check for correctness
	if tkStack.top().GetType() == constant.S_ENDLINE && tokenType == constant.S_ENDLINE {
		return true
	}
	return false
}

func getTableCell(tableMap map[string][]int, tkStack *RuleStack, currentToken uint8) uint8 {
	if tkStack.top().IsTerminal() {
		fmt.Println("getTableCell: if 1")
		return uint8(0)
	}
	if GetTokenColumn(currentToken) > 30 {
		fmt.Println("getTableCell: if 2")
		return uint8(0)
	}
	// If epsilon. Pop stack
	if len(tableMap[tkStack.top().GetLabel()]) == 0 {
		fmt.Println("getTableCell: if 3")
		tkStack.pop()
		return uint8(200)
	}
	fmt.Printf("CELL PROD: %d\n", uint8(tableMap[tkStack.top().GetLabel()][GetTokenColumn(currentToken)]))
	return uint8(tableMap[tkStack.top().GetLabel()][GetTokenColumn(currentToken)])
}

func pushRulesToStack(tkStack *RuleStack, rule ProductionRule) {
	fmt.Println("Push to stack!")
	tkStack.pop()
	size := len(rule.RuleList)
	for i := size-1; i >= 0; i-- {
		genericElement := GenericElement{
			Type:  rule.RuleList[i],
			Label: grammarRules.getLabelByUID(rule.RuleList[i]),
		}
		fmt.Printf("Rule: %d, ", genericElement.GetType())
		tkStack.push(genericElement)
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

/* Generic Element */
type GenericElement struct {
	Type uint8
	Label string
}

func (ge GenericElement) IsTerminal() bool {
	return ge.Type < 60
}

func (ge GenericElement) GetType() uint8 {
	return ge.Type
}

func (ge GenericElement) GetLabel() string {
	return ge.Label
}
/* END Generic Element */

/* Token Stack */
type RuleStack struct {
	collection []ProductionRuleElement
}

func (s *RuleStack) push(element ProductionRuleElement) {
	s.collection = append(s.collection, element)
}

func (s *RuleStack) pop()  {
	if len(s.collection) == 0 {
		return
	}
	s.collection = s.collection[:len(s.collection)-1]
}

func (s *RuleStack) top() ProductionRuleElement {
	if len(s.collection) == 0 {
		return nil
	}
	return s.collection[len(s.collection)-1]
}

func (s *RuleStack) isEmpty() bool {
	return len(s.collection) == 0
}
/* END Token Stack */
