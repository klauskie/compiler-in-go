package parser

import (
	"../aux/constant"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func FillParsingTable(file string) (map[string][]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	tableMap := make(map[string][]int)
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return tableMap, err
		}

		var tempList []int

		for _, num := range row[1:] {
			if numInt, err := strconv.Atoi(num); err == nil {
				tempList = append(tempList, numInt)
			}
			if len(num) == 0 {
				tempList = append(tempList, 0)
			}
		}

		tableMap[row[0]] = tempList
	}
}

func FillParsingTable2(file string) (map[uint8][]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	tableMap := make(map[uint8][]int)
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return tableMap, err
		}

		var tempList []int

		for _, num := range row[1:] {
			if numInt, err := strconv.Atoi(num); err == nil {
				tempList = append(tempList, numInt)
			}
			if len(num) == 0 {
				tempList = append(tempList, 0)
			}
		}

		if numInt, err := strconv.Atoi(row[0]); err == nil {
			tableMap[uint8(numInt)] = tempList
		}
	}
}

// ID	NUM	IF	ELSE	INT	VOID	RETURN	WHILE	INPUT	OUTPUT	+	-	*	/	<	<=	>	>=	==	!=	=	;	,	(	)	[	]	{	}	$

func GetTokenColumn(tokenID uint8) int {
	switch tokenID {
	case constant.H_WORD:
		return 0
	case constant.H_NUMBER:
		return 1
	case constant.K_IF_ID:
		return 2
	case constant.K_ELSE_ID:
		return 3
	case constant.K_INT_ID:
		return 4
	case constant.K_VOID_ID:
		return 5
	case constant.K_RETURN_ID:
		return 6
	case constant.K_WHILE_ID:
		return 7
	case constant.K_INPUT_ID:
		return 8
	case constant.K_OUTPUT_ID:
		return 9
	case constant.S_SUM:
		return 10
	case constant.S_SUBTRACT:
		return 11
	case constant.S_ASTERISK:
		return 12
	case constant.S_FORWARD_SLASH:
		return 13
	case constant.S_LESS:
		return 14
	case constant.S_LESS_EQUAL:
		return 15
	case constant.S_MORE:
		return 16
	case constant.S_MORE_EQUAL:
		return 17
	case constant.S_EQUAL_EQUAL:
		return 18
	case constant.S_NOT_EQUAL:
		return 19
	case constant.S_EQUAL:
		return 20
	case constant.S_SEMICOLON:
		return 21
	case constant.S_COMMA:
		return 22
	case constant.S_OPEN_PARENTHESIS:
		return 23
	case constant.S_CLOSE_PARENTHESIS:
		return 24
	case constant.S_OPEN_SQR_BRACKET:
		return 25
	case constant.S_CLOSE_SQR_BRACKET:
		return 26
	case constant.S_OPEN_CURLY_BRACKET:
		return 27
	case constant.S_CLOSE_CURLY_BRACKET:
		return 28
	case constant.S_ENDLINE:
		return 29
	default:
		return 100
	}
}