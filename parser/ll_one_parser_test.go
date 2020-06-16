package parser

import (
	"../aux/constant"
	"../lexer"
	"testing"
)

// Change parsing table directory
func TestRunParser(t *testing.T) {
	// Var setup
	fn := "../test_files/test4.txt"
	tokenListTest_1, _ := lexer.Run(fn)
	cleanVariables()

	// Test
	res_1 := RunParser(fn, tokenListTest_1)
	if res_1 != nil {
		t.Errorf("RunParser; expexted no errors, got %s", res_1.ToString())
	}

}

func TestPushRulesToStack(t *testing.T) {
	tkStack := RuleStack{}
	prodRule := ProductionRule{
		Name:     "fun_declaration_void",
		NameID:   constant.R_FUN_DECLARATION_VOID,
		RuleList: []uint8{constant.S_OPEN_PARENTHESIS, constant.R_PARAMS, constant.S_CLOSE_PARENTHESIS, constant.R_COMPOUND_STMT_VOID_INIT},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS},
	}
	pushRulesToStack(&tkStack, prodRule)
	if tkStack.top() != constant.S_OPEN_PARENTHESIS {
		t.Errorf("PushRulesToStack; After pushing production rule, top of stack should be %d and got %d", constant.S_OPEN_PARENTHESIS, tkStack.top())
	}

	tkStack.pop()
	if tkStack.top() != constant.R_PARAMS {
		t.Errorf("PushRulesToStack; After pop, top of stack should be %d and got %d", constant.R_PARAMS, tkStack.top())
	}

	tkStack.pop()
	if tkStack.top() != constant.S_CLOSE_PARENTHESIS {
		t.Errorf("PushRulesToStack; After pop, top of stack should be %d and got %d", constant.S_CLOSE_PARENTHESIS, tkStack.top())
	}

	tkStack.pop()
	if tkStack.top() != constant.R_COMPOUND_STMT_VOID_INIT {
		t.Errorf("PushRulesToStack; After pop, top of stack should be %d and got %d", constant.R_COMPOUND_STMT_VOID_INIT, tkStack.top())
	}
}

func TestNextToken(t *testing.T) {
	// Var setup
	/*
	fn := "../test_files/test4.txt"
	tokenListTest_1, _ := lexer.Run(fn)
	cleanVariables()
	RunParser(fn, tokenListTest_1)

	// Test
	nextToken()
	if tokenIndex == 0 {
		t.Errorf("NextToken; expected token index to increment; got %d", tokenIndex)
	}
	 */
}

