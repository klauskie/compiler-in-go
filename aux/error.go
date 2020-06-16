package aux

import "fmt"

type FoulError interface {
	ToString() string
}

type Foul struct {
	Type     	uint8
	Message		string
}

func (f Foul) ToString() string {
	return fmt.Sprintf("Error status %d: %s", f.Type, f.Message)
}

const (
	UNKNOWN_TOKEN uint8 = iota + 1
	UNEXPECTED_EOF
	UNEXPECTED_TOKEN
	INVALID_EXPRESSION
	GENERIC_ERROR
)

func NewFoul(t uint8, params ...interface{}) Foul {
	var message string
	switch t {
	case UNKNOWN_TOKEN:
		message = fmt.Sprintf("error: unknown token [ %s ] at line:  %d", params[1].(string), params[0].(int))
		break
	case UNEXPECTED_EOF:
		message = fmt.Sprintf("error: unexpected end of file:  %s", params[0].(string))
		break
	case UNEXPECTED_TOKEN:
		message = fmt.Sprintf("error: expected '%s' before '%s' token at line %d", params[0].(string), params[1].(string), params[2].(int))
		break
	case INVALID_EXPRESSION:
		message = fmt.Sprintf("error: invalid expression before '%s' token at line %d", params[0].(string), params[1].(int))
		break
	case GENERIC_ERROR:
		message = fmt.Sprintf("error: syntax error at line %d", params[0].(int))
		break
	default:
		message = "error"
	}

	return Foul{t, message}
}