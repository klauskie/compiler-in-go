package parser

import (
	"../aux/constant"
)

type ProductionRule struct {
	Name string
	NameID uint8
	RuleList []uint8
	PlusSet []uint8
}

type GrammarRules struct {
	RuleList []ProductionRule
	RuleMap map[uint8] ProductionRule
}

func initGrammarRules() *GrammarRules {
	gr := GrammarRules{}
	gr.RuleMap = make(map[uint8] ProductionRule)
	gr.setRules()
	return &gr
}

func (g *GrammarRules) getRuleByID(id uint8) ProductionRule {
	return g.RuleMap[id]
}

func (g *GrammarRules) getLabelByUID(id uint8) string {
	for i := uint8(0); i >= 89; i++ {
		if g.RuleMap[i].NameID == id {
			return g.RuleMap[i].Name
		}
	}
	return ""
}

func (g *GrammarRules) setRules() {
	g.RuleMap[0] = ProductionRule{
		Name:     "end",
		NameID:   constant.S_ENDLINE,
		RuleList: []uint8{},
		PlusSet:  []uint8{},
	}
	g.RuleMap[1] = ProductionRule{
		Name:     "program",
		NameID:   constant.R_PROGRAM,
		RuleList: []uint8{constant.R_DECLARATION, constant.R_DECLARATION_LIST},
		PlusSet:  []uint8{constant.K_INT_ID, constant.K_VOID_ID},
	}
	g.RuleMap[2] = ProductionRule{
		Name:     "declaration_list",
		NameID:   constant.R_DECLARATION_LIST,
		RuleList: []uint8{constant.R_DECLARATION, constant.R_DECLARATION_LIST},
		PlusSet:  []uint8{constant.K_INT_ID, constant.K_VOID_ID},
	}
	g.RuleMap[3] = ProductionRule{
		Name:     "declaration_list",
		NameID:   constant.R_DECLARATION_LIST,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_ENDLINE},
	}
	g.RuleMap[4] = ProductionRule{
		Name:     "declaration",
		NameID:   constant.R_DECLARATION,
		RuleList: []uint8{constant.K_INT_ID, constant.H_WORD, constant.R_DECLARATION_AUX},
		PlusSet:  []uint8{constant.K_INT_ID},
	}
	g.RuleMap[5] = ProductionRule{
		Name:     "declaration",
		NameID:   constant.R_DECLARATION,
		RuleList: []uint8{constant.K_VOID_ID, constant.H_WORD, constant.R_FUN_DECLARATION_VOID},
		PlusSet:  []uint8{constant.K_VOID_ID},
	}
	g.RuleMap[6] = ProductionRule{
		Name:     "declaration_aux",
		NameID:   constant.R_DECLARATION_AUX,
		RuleList: []uint8{constant.R_VAR_DECLARATION_DECORATION, constant.S_SEMICOLON},
		PlusSet:  []uint8{constant.S_OPEN_SQR_BRACKET, constant.S_SEMICOLON},
	}
	g.RuleMap[7] = ProductionRule{
		Name:     "declaration_aux",
		NameID:   constant.R_DECLARATION_AUX,
		RuleList: []uint8{constant.R_FUN_DECLARATION_TYPED},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS},
	}
	g.RuleMap[8] = ProductionRule{
		Name:     "var_declaration_decoration",
		NameID:   constant.R_VAR_DECLARATION_DECORATION,
		RuleList: []uint8{constant.S_OPEN_SQR_BRACKET, constant.H_NUMBER, constant.S_CLOSE_SQR_BRACKET},
		PlusSet:  []uint8{constant.S_OPEN_SQR_BRACKET},
	}
	g.RuleMap[9] = ProductionRule{
		Name:     "var_declaration_decoration",
		NameID:   constant.R_VAR_DECLARATION_DECORATION,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_SEMICOLON},
	}
	g.RuleMap[10] = ProductionRule{
		Name:     "fun_declaration_void",
		NameID:   constant.R_FUN_DECLARATION_VOID,
		RuleList: []uint8{constant.S_OPEN_PARENTHESIS, constant.R_PARAMS, constant.S_CLOSE_PARENTHESIS, constant.R_COMPOUND_STMT_VOID_INIT},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS},
	}
	g.RuleMap[11] = ProductionRule{
		Name:     "fun_declaration_typed",
		NameID:   constant.R_FUN_DECLARATION_TYPED,
		RuleList: []uint8{constant.S_OPEN_PARENTHESIS, constant.R_PARAMS, constant.S_CLOSE_PARENTHESIS, constant.R_COMPOUND_STMT_TYPED_RETURN},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS},
	}
	g.RuleMap[12] = ProductionRule{
		Name:     "params",
		NameID:   constant.R_PARAMS,
		RuleList: []uint8{constant.R_PARAM_LIST},
		PlusSet:  []uint8{constant.K_INT_ID},
	}
	g.RuleMap[13] = ProductionRule{
		Name:     "params",
		NameID:   constant.R_PARAMS,
		RuleList: []uint8{constant.K_VOID_ID},
		PlusSet:  []uint8{constant.K_VOID_ID},
	}
	g.RuleMap[14] = ProductionRule{
		Name:     "param_list",
		NameID:   constant.R_PARAM_LIST,
		RuleList: []uint8{constant.R_PARAM, constant.R_PARAM_LIST_AUX},
		PlusSet:  []uint8{constant.K_INT_ID},
	}
	g.RuleMap[15] = ProductionRule{
		Name:     "param_list_aux",
		NameID:   constant.R_PARAM_LIST_AUX,
		RuleList: []uint8{constant.S_COMMA, constant.R_PARAM, constant.R_PARAM_LIST_AUX},
		PlusSet:  []uint8{constant.S_COMMA},
	}
	g.RuleMap[16] = ProductionRule{
		Name:     "param_list_aux",
		NameID:   constant.R_PARAM_LIST_AUX,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_CLOSE_PARENTHESIS},
	}
	g.RuleMap[17] = ProductionRule{
		Name:     "param",
		NameID:   constant.R_PARAM,
		RuleList: []uint8{constant.K_INT_ID, constant.H_WORD, constant.R_PARAM_DECORATOR},
		PlusSet:  []uint8{constant.K_INT_ID},
	}
	g.RuleMap[18] = ProductionRule{
		Name:     "param_decorator",
		NameID:   constant.R_PARAM_DECORATOR,
		RuleList: []uint8{constant.S_OPEN_SQR_BRACKET, constant.S_CLOSE_SQR_BRACKET},
		PlusSet:  []uint8{constant.S_OPEN_SQR_BRACKET},
	}
	g.RuleMap[19] = ProductionRule{
		Name:     "param_decorator",
		NameID:   constant.R_PARAM_DECORATOR,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_COMMA, constant.S_CLOSE_PARENTHESIS},
	}
	g.RuleMap[20] = ProductionRule{
		Name:     "compound_stmt_void",
		NameID:   constant.R_COMPOUND_STMT_VOID,
		RuleList: []uint8{constant.S_OPEN_CURLY_BRACKET, constant.R_LOCAL_DECLARATIONS, constant.R_STATEMENT_LIST_VOID, constant.S_CLOSE_CURLY_BRACKET},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET},
	}
	g.RuleMap[21] = ProductionRule{
		Name:     "compound_stmt_void_init",
		NameID:   constant.R_COMPOUND_STMT_VOID_INIT,
		RuleList: []uint8{constant.S_OPEN_CURLY_BRACKET, constant.R_LOCAL_DECLARATIONS, constant.R_STATEMENT_LIST_VOID, constant.S_CLOSE_CURLY_BRACKET},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET},
	}
	g.RuleMap[22] = ProductionRule{
		Name:     "compound_stmt_typed_return",
		NameID:   constant.R_COMPOUND_STMT_TYPED_RETURN,
		RuleList: []uint8{constant.S_OPEN_CURLY_BRACKET, constant.R_LOCAL_DECLARATIONS, constant.R_STATEMENT_LIST_TYPED, constant.R_RETURN_STMT_TYPED, constant.S_CLOSE_CURLY_BRACKET},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET},
	}
	g.RuleMap[23] = ProductionRule{
		Name:     "compound_stmt_typed",
		NameID:   constant.R_COMPOUND_STMT_TYPED,
		RuleList: []uint8{constant.S_OPEN_CURLY_BRACKET, constant.R_LOCAL_DECLARATIONS, constant.R_STATEMENT_LIST_TYPED, constant.S_CLOSE_CURLY_BRACKET},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET},
	}
	g.RuleMap[24] = ProductionRule{
		Name:     "local_declarations",
		NameID:   constant.R_LOCAL_DECLARATIONS,
		RuleList: []uint8{constant.K_INT_ID, constant.H_WORD, constant.R_VAR_DECLARATION_DECORATION, constant.S_SEMICOLON, constant.R_LOCAL_DECLARATIONS},
		PlusSet:  []uint8{constant.K_INT_ID},
	}
	g.RuleMap[25] = ProductionRule{
		Name:     "local_declarations",
		NameID:   constant.R_LOCAL_DECLARATIONS,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET, constant.K_IF_ID, constant.K_WHILE_ID, constant.K_RETURN_ID, constant.H_WORD, constant.K_INPUT_ID, constant.K_OUTPUT_ID, constant.S_CLOSE_CURLY_BRACKET},
	}
	g.RuleMap[26] = ProductionRule{
		Name:     "statement_list_void",
		NameID:   constant.R_STATEMENT_LIST_VOID,
		RuleList: []uint8{constant.R_STATEMENT_VOID, constant.R_STATEMENT_LIST_VOID},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET, constant.K_IF_ID, constant.K_WHILE_ID, constant.K_RETURN_ID, constant.H_WORD, constant.K_INPUT_ID, constant.K_OUTPUT_ID},
	}
	g.RuleMap[27] = ProductionRule{
		Name:     "statement_list_void",
		NameID:   constant.R_STATEMENT_LIST_VOID,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_CLOSE_CURLY_BRACKET},
	}
	g.RuleMap[28] = ProductionRule{
		Name:     "statement_list_typed",
		NameID:   constant.R_STATEMENT_LIST_TYPED,
		RuleList: []uint8{constant.R_STATEMENT_TYPED, constant.R_STATEMENT_LIST_TYPED},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET, constant.K_IF_ID, constant.K_WHILE_ID, constant.K_RETURN_ID, constant.H_WORD, constant.K_INPUT_ID, constant.K_OUTPUT_ID},
	}
	g.RuleMap[29] = ProductionRule{
		Name:     "statement_list_typed",
		NameID:   constant.R_STATEMENT_LIST_TYPED,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_CLOSE_CURLY_BRACKET},
	}
	g.RuleMap[30] = ProductionRule{
		Name:     "statement_void",
		NameID:   constant.R_STATEMENT_VOID,
		RuleList: []uint8{constant.R_COMPOUND_STMT_VOID},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET},
	}
	g.RuleMap[31] = ProductionRule{
		Name:     "statement_void",
		NameID:   constant.R_STATEMENT_VOID,
		RuleList: []uint8{constant.R_SELECTION_STMT_VOID},
		PlusSet:  []uint8{constant.K_IF_ID},
	}
	g.RuleMap[32] = ProductionRule{
		Name:     "statement_void",
		NameID:   constant.R_STATEMENT_VOID,
		RuleList: []uint8{constant.R_ITERATION_STMT_VOID},
		PlusSet:  []uint8{constant.K_WHILE_ID},
	}
	g.RuleMap[33] = ProductionRule{
		Name:     "statement_void",
		NameID:   constant.R_STATEMENT_VOID,
		RuleList: []uint8{constant.R_RETURN_STMT_VOID},
		PlusSet:  []uint8{constant.K_RETURN_ID},
	}
	g.RuleMap[34] = ProductionRule{
		Name:     "statement_void",
		NameID:   constant.R_STATEMENT_VOID,
		RuleList: []uint8{constant.R_STATEMENT_GENERICS},
		PlusSet:  []uint8{constant.H_WORD, constant.K_INPUT_ID, constant.K_OUTPUT_ID},
	}
	g.RuleMap[35] = ProductionRule{
		Name:     "statement_typed",
		NameID:   constant.R_STATEMENT_TYPED,
		RuleList: []uint8{constant.R_COMPOUND_STMT_TYPED},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET},
	}
	g.RuleMap[36] = ProductionRule{
		Name:     "statement_typed",
		NameID:   constant.R_STATEMENT_TYPED,
		RuleList: []uint8{constant.R_SELECTION_STMT_TYPED},
		PlusSet:  []uint8{constant.K_IF_ID},
	}
	g.RuleMap[37] = ProductionRule{
		Name:     "statement_typed",
		NameID:   constant.R_STATEMENT_TYPED,
		RuleList: []uint8{constant.R_ITERATION_STMT_TYPED},
		PlusSet:  []uint8{constant.K_WHILE_ID},
	}
	g.RuleMap[38] = ProductionRule{
		Name:     "statement_typed",
		NameID:   constant.R_STATEMENT_TYPED,
		RuleList: []uint8{constant.R_RETURN_STMT_TYPED},
		PlusSet:  []uint8{constant.K_RETURN_ID},
	}
	g.RuleMap[39] = ProductionRule{
		Name:     "statement_typed",
		NameID:   constant.R_STATEMENT_TYPED,
		RuleList: []uint8{constant.R_STATEMENT_GENERICS},
		PlusSet:  []uint8{constant.H_WORD, constant.K_INPUT_ID, constant.K_OUTPUT_ID},
	}
	g.RuleMap[40] = ProductionRule{
		Name:     "statement_generics",
		NameID:   constant.R_STATEMENT_GENERICS,
		RuleList: []uint8{constant.R_ASSIGNMENT_CALL_STMT},
		PlusSet:  []uint8{constant.H_WORD},
	}
	g.RuleMap[41] = ProductionRule{
		Name:     "statement_generics",
		NameID:   constant.R_STATEMENT_GENERICS,
		RuleList: []uint8{constant.R_INPUT_STMT},
		PlusSet:  []uint8{constant.K_INPUT_ID},
	}
	g.RuleMap[42] = ProductionRule{
		Name:     "statement_generics",
		NameID:   constant.R_STATEMENT_GENERICS,
		RuleList: []uint8{constant.R_OUTPUT_STMT},
		PlusSet:  []uint8{constant.K_OUTPUT_ID},
	}
	g.RuleMap[43] = ProductionRule{
		Name:     "assignment_call_stmt",
		NameID:   constant.R_ASSIGNMENT_CALL_STMT,
		RuleList: []uint8{constant.H_WORD, constant.R_ASSIGNMENT_CALL_STMT_AUX, constant.S_SEMICOLON},
		PlusSet:  []uint8{constant.H_WORD},
	}
	g.RuleMap[44] = ProductionRule{
		Name:     "assignment_call_stmt_aux",
		NameID:   constant.R_ASSIGNMENT_CALL_STMT_AUX,
		RuleList: []uint8{constant.R_VAR_DECORATION, constant.S_EQUAL, constant.R_EXPRESSION},
		PlusSet:  []uint8{constant.S_OPEN_SQR_BRACKET, constant.S_EQUAL},
	}
	g.RuleMap[45] = ProductionRule{
		Name:     "assignment_call_stmt_aux",
		NameID:   constant.R_ASSIGNMENT_CALL_STMT_AUX,
		RuleList: []uint8{constant.S_OPEN_PARENTHESIS, constant.R_ARGS, constant.S_CLOSE_PARENTHESIS},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS},
	}
	g.RuleMap[46] = ProductionRule{
		Name:     "selection_stmt_void",
		NameID:   constant.R_SELECTION_STMT_VOID,
		RuleList: []uint8{constant.K_IF_ID, constant.S_OPEN_PARENTHESIS, constant.R_EXPRESSION, constant.S_CLOSE_PARENTHESIS, constant.R_STATEMENT_VOID, constant.R_SELECTION_STMT_VOID_AUX},
		PlusSet:  []uint8{constant.K_IF_ID},
	}
	g.RuleMap[47] = ProductionRule{
		Name:     "selection_stmt_void_aux",
		NameID:   constant.R_SELECTION_STMT_VOID_AUX,
		RuleList: []uint8{constant.K_ELSE_ID, constant.R_STATEMENT_VOID},
		PlusSet:  []uint8{constant.K_ELSE_ID},
	}
	g.RuleMap[48] = ProductionRule{
		Name:     "selection_stmt_void_aux",
		NameID:   constant.R_SELECTION_STMT_VOID_AUX,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET, constant.K_IF_ID, constant.K_WHILE_ID, constant.K_RETURN_ID, constant.H_WORD, constant.K_INPUT_ID, constant.K_OUTPUT_ID, constant.S_CLOSE_CURLY_BRACKET},
	}
	g.RuleMap[49] = ProductionRule{
		Name:     "selection_stmt_typed",
		NameID:   constant.R_SELECTION_STMT_TYPED,
		RuleList: []uint8{constant.K_IF_ID, constant.S_OPEN_PARENTHESIS, constant.R_EXPRESSION, constant.S_CLOSE_PARENTHESIS, constant.R_STATEMENT_TYPED, constant.R_SELECTION_STMT_TYPED_AUX},
		PlusSet:  []uint8{constant.K_IF_ID},
	}
	g.RuleMap[50] = ProductionRule{
		Name:     "selection_stmt_typed_aux",
		NameID:   constant.R_SELECTION_STMT_TYPED_AUX,
		RuleList: []uint8{constant.K_ELSE_ID, constant.R_STATEMENT_TYPED},
		PlusSet:  []uint8{constant.K_ELSE_ID},
	}
	g.RuleMap[51] = ProductionRule{
		Name:     "selection_stmt_typed_aux",
		NameID:   constant.R_SELECTION_STMT_TYPED_AUX,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_OPEN_CURLY_BRACKET, constant.K_IF_ID, constant.K_WHILE_ID, constant.K_RETURN_ID, constant.H_WORD, constant.K_INPUT_ID, constant.K_OUTPUT_ID, constant.S_CLOSE_CURLY_BRACKET},
	}
	g.RuleMap[52] = ProductionRule{
		Name:     "iteration_stmt_void",
		NameID:   constant.R_ITERATION_STMT_VOID,
		RuleList: []uint8{constant.K_WHILE_ID, constant.S_OPEN_PARENTHESIS, constant.R_EXPRESSION, constant.S_CLOSE_PARENTHESIS, constant.R_STATEMENT_VOID},
		PlusSet:  []uint8{constant.K_WHILE_ID},
	}
	g.RuleMap[53] = ProductionRule{
		Name:     "iteration_stmt_typed",
		NameID:   constant.R_ITERATION_STMT_TYPED,
		RuleList: []uint8{constant.K_WHILE_ID, constant.S_OPEN_PARENTHESIS, constant.R_EXPRESSION, constant.S_CLOSE_PARENTHESIS, constant.R_STATEMENT_TYPED},
		PlusSet:  []uint8{constant.K_WHILE_ID},
	}
	g.RuleMap[54] = ProductionRule{
		Name:     "return_stmt_typed",
		NameID:   constant.R_RETURN_STMT_TYPED,
		RuleList: []uint8{constant.K_RETURN_ID, constant.R_EXPRESSION, constant.S_SEMICOLON},
		PlusSet:  []uint8{constant.K_RETURN_ID},
	}
	g.RuleMap[55] = ProductionRule{
		Name:     "return_stmt_void",
		NameID:   constant.R_RETURN_STMT_VOID,
		RuleList: []uint8{constant.K_RETURN_ID, constant.S_SEMICOLON},
		PlusSet:  []uint8{constant.K_RETURN_ID},
	}
	g.RuleMap[56] = ProductionRule{
		Name:     "input_stmt",
		NameID:   constant.R_INPUT_STMT,
		RuleList: []uint8{constant.K_INPUT_ID, constant.H_WORD, constant.R_VAR_DECORATION, constant.S_SEMICOLON},
		PlusSet:  []uint8{constant.K_INPUT_ID},
	}
	g.RuleMap[57] = ProductionRule{
		Name:     "output_stmt",
		NameID:   constant.R_OUTPUT_STMT,
		RuleList: []uint8{constant.K_OUTPUT_ID, constant.R_EXPRESSION, constant.S_SEMICOLON},
		PlusSet:  []uint8{constant.K_OUTPUT_ID},
	}
	g.RuleMap[58] = ProductionRule{
		Name:     "var_decoration",
		NameID:   constant.R_VAR_DECORATION,
		RuleList: []uint8{constant.S_OPEN_SQR_BRACKET, constant.R_ARITHMETIC_EXPRESSION, constant.S_CLOSE_SQR_BRACKET},
		PlusSet:  []uint8{constant.S_OPEN_SQR_BRACKET},
	}
	g.RuleMap[59] = ProductionRule{
		Name:     "var_decoration",
		NameID:   constant.R_VAR_DECORATION,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_EQUAL, constant.S_SEMICOLON, constant.S_ASTERISK, constant.S_FORWARD_SLASH, constant.S_SUM, constant.S_SUBTRACT, constant.S_CLOSE_SQR_BRACKET, constant.S_LESS_EQUAL, constant.S_LESS, constant.S_MORE, constant.S_MORE_EQUAL, constant.S_EQUAL_EQUAL, constant.S_NOT_EQUAL, constant.S_CLOSE_PARENTHESIS, constant.S_COMMA},
	}
	g.RuleMap[60] = ProductionRule{
		Name:     "expression",
		NameID:   constant.R_EXPRESSION,
		RuleList: []uint8{constant.R_ARITHMETIC_EXPRESSION, constant.R_EXPRESSION_AUX},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS, constant.H_NUMBER, constant.H_WORD},
	}
	g.RuleMap[61] = ProductionRule{
		Name:     "expression_aux",
		NameID:   constant.R_EXPRESSION_AUX,
		RuleList: []uint8{constant.R_RELOP, constant.R_ARITHMETIC_EXPRESSION},
		PlusSet:  []uint8{constant.S_LESS_EQUAL, constant.S_LESS, constant.S_MORE, constant.S_MORE_EQUAL, constant.S_EQUAL_EQUAL, constant.S_NOT_EQUAL},
	}
	g.RuleMap[62] = ProductionRule{
		Name:     "expression_aux",
		NameID:   constant.R_EXPRESSION_AUX,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_SEMICOLON, constant.S_CLOSE_PARENTHESIS},
	}
	g.RuleMap[63] = ProductionRule{
		Name:     "relop",
		NameID:   constant.R_RELOP,
		RuleList: []uint8{constant.S_LESS_EQUAL},
		PlusSet:  []uint8{constant.S_LESS_EQUAL},
	}
	g.RuleMap[64] = ProductionRule{
		Name:     "relop",
		NameID:   constant.R_RELOP,
		RuleList: []uint8{constant.S_LESS},
		PlusSet:  []uint8{constant.S_LESS},
	}
	g.RuleMap[65] = ProductionRule{
		Name:     "relop",
		NameID:   constant.R_RELOP,
		RuleList: []uint8{constant.S_MORE},
		PlusSet:  []uint8{constant.S_MORE},
	}
	g.RuleMap[66] = ProductionRule{
		Name:     "relop",
		NameID:   constant.R_RELOP,
		RuleList: []uint8{constant.S_MORE_EQUAL},
		PlusSet:  []uint8{constant.S_MORE_EQUAL},
	}
	g.RuleMap[67] = ProductionRule{
		Name:     "relop",
		NameID:   constant.R_RELOP,
		RuleList: []uint8{constant.S_EQUAL_EQUAL},
		PlusSet:  []uint8{constant.S_EQUAL_EQUAL},
	}
	g.RuleMap[68] = ProductionRule{
		Name:     "relop",
		NameID:   constant.R_RELOP,
		RuleList: []uint8{constant.S_NOT_EQUAL},
		PlusSet:  []uint8{constant.S_NOT_EQUAL},
	}
	g.RuleMap[69] = ProductionRule{
		Name:     "arithmetic_expression",
		NameID:   constant.R_ARITHMETIC_EXPRESSION,
		RuleList: []uint8{constant.R_TERM, constant.R_ARITHMETIC_EXPRESSION_AUX},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS, constant.H_NUMBER, constant.H_WORD},
	}
	g.RuleMap[70] = ProductionRule{
		Name:     "arithmetic_expression_aux",
		NameID:   constant.R_ARITHMETIC_EXPRESSION_AUX,
		RuleList: []uint8{constant.R_ADDOP, constant.R_TERM, constant.R_ARITHMETIC_EXPRESSION_AUX},
		PlusSet:  []uint8{constant.S_SUM, constant.S_SUBTRACT},
	}
	g.RuleMap[71] = ProductionRule{
		Name:     "arithmetic_expression_aux",
		NameID:   constant.R_ARITHMETIC_EXPRESSION_AUX,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_CLOSE_SQR_BRACKET, constant.S_LESS_EQUAL, constant.S_LESS, constant.S_MORE, constant.S_MORE_EQUAL, constant.S_EQUAL_EQUAL, constant.S_NOT_EQUAL, constant.S_CLOSE_PARENTHESIS, constant.S_COMMA},
	}
	g.RuleMap[72] = ProductionRule{
		Name:     "addop",
		NameID:   constant.R_ADDOP,
		RuleList: []uint8{constant.S_SUM},
		PlusSet:  []uint8{constant.S_SUM},
	}
	g.RuleMap[73] = ProductionRule{
		Name:     "addop",
		NameID:   constant.R_ADDOP,
		RuleList: []uint8{constant.S_SUBTRACT},
		PlusSet:  []uint8{constant.S_SUBTRACT},
	}
	g.RuleMap[74] = ProductionRule{
		Name:     "term",
		NameID:   constant.R_TERM,
		RuleList: []uint8{constant.R_FACTOR, constant.R_TERM_AUX},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS, constant.H_NUMBER, constant.H_WORD},
	}
	g.RuleMap[75] = ProductionRule{
		Name:     "term_aux",
		NameID:   constant.R_TERM_AUX,
		RuleList: []uint8{constant.R_MULOP, constant.R_FACTOR, constant.R_TERM_AUX},
		PlusSet:  []uint8{constant.S_ASTERISK, constant.S_FORWARD_SLASH},
	}
	g.RuleMap[76] = ProductionRule{
		Name:     "term_aux",
		NameID:   constant.R_TERM_AUX,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_SUM, constant.S_SUBTRACT, constant.S_CLOSE_SQR_BRACKET, constant.S_LESS_EQUAL, constant.S_LESS, constant.S_MORE, constant.S_MORE_EQUAL, constant.S_EQUAL_EQUAL, constant.S_NOT_EQUAL, constant.S_CLOSE_PARENTHESIS, constant.S_COMMA},
	}
	g.RuleMap[77] = ProductionRule{
		Name:     "mulop",
		NameID:   constant.R_MULOP,
		RuleList: []uint8{constant.S_ASTERISK},
		PlusSet:  []uint8{constant.S_ASTERISK},
	}
	g.RuleMap[78] = ProductionRule{
		Name:     "mulop",
		NameID:   constant.R_MULOP,
		RuleList: []uint8{constant.S_FORWARD_SLASH},
		PlusSet:  []uint8{constant.S_FORWARD_SLASH},
	}
	g.RuleMap[79] = ProductionRule{
		Name:     "factor",
		NameID:   constant.R_FACTOR,
		RuleList: []uint8{constant.S_OPEN_PARENTHESIS, constant.R_ARITHMETIC_EXPRESSION, constant.S_CLOSE_PARENTHESIS},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS},
	}
	g.RuleMap[80] = ProductionRule{
		Name:     "factor",
		NameID:   constant.R_FACTOR,
		RuleList: []uint8{constant.R_FACTOR_ID},
		PlusSet:  []uint8{constant.H_WORD},
	}
	g.RuleMap[81] = ProductionRule{
		Name:     "factor",
		NameID:   constant.R_FACTOR,
		RuleList: []uint8{constant.H_NUMBER},
		PlusSet:  []uint8{constant.H_NUMBER},
	}
	g.RuleMap[82] = ProductionRule{
		Name:     "factor_ID",
		NameID:   constant.R_FACTOR_ID,
		RuleList: []uint8{constant.H_WORD, constant.R_FACTOR_ID_DECORATION},
		PlusSet:  []uint8{constant.H_WORD},
	}
	g.RuleMap[83] = ProductionRule{
		Name:     "factor_ID_decoration",
		NameID:   constant.R_FACTOR_ID_DECORATION,
		RuleList: []uint8{constant.R_VAR_DECORATION},
		PlusSet:  []uint8{constant.S_OPEN_SQR_BRACKET, constant.S_ASTERISK, constant.S_FORWARD_SLASH, constant.S_SUM, constant.S_SUBTRACT, constant.S_SUM, constant.S_SUBTRACT, constant.S_CLOSE_SQR_BRACKET, constant.S_LESS_EQUAL, constant.S_LESS, constant.S_MORE, constant.S_MORE_EQUAL, constant.S_EQUAL_EQUAL, constant.S_NOT_EQUAL, constant.S_CLOSE_PARENTHESIS, constant.S_COMMA},
	}
	g.RuleMap[84] = ProductionRule{
		Name:     "factor_ID_decoration",
		NameID:   constant.R_FACTOR_ID_DECORATION,
		RuleList: []uint8{constant.S_OPEN_PARENTHESIS, constant.R_ARGS, constant.S_CLOSE_PARENTHESIS},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS},
	}
	g.RuleMap[85] = ProductionRule{
		Name:     "args",
		NameID:   constant.R_ARGS,
		RuleList: []uint8{constant.R_ARGS_LIST},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS, constant.H_NUMBER, constant.H_WORD},
	}
	g.RuleMap[86] = ProductionRule{
		Name:     "args",
		NameID:   constant.R_ARGS,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_CLOSE_PARENTHESIS},
	}
	g.RuleMap[87] = ProductionRule{
		Name:     "args_list",
		NameID:   constant.R_ARGS_LIST,
		RuleList: []uint8{constant.R_ARITHMETIC_EXPRESSION, constant.R_ARGS_LIST_AUX},
		PlusSet:  []uint8{constant.S_OPEN_PARENTHESIS, constant.H_NUMBER, constant.H_WORD},
	}
	g.RuleMap[88] = ProductionRule{
		Name:     "args_list_aux",
		NameID:   constant.R_ARGS_LIST_AUX,
		RuleList: []uint8{constant.S_COMMA, constant.R_ARITHMETIC_EXPRESSION, constant.R_ARGS_LIST_AUX},
		PlusSet:  []uint8{constant.S_COMMA},
	}
	g.RuleMap[89] = ProductionRule{
		Name:     "args_list_aux",
		NameID:   constant.R_ARGS_LIST_AUX,
		RuleList: []uint8{},
		PlusSet:  []uint8{constant.S_CLOSE_PARENTHESIS},
	}
}
