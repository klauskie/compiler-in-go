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
)

func NewFoul(t uint8, params ...interface{}) Foul {
	var message string
	switch t {
	case UNKNOWN_TOKEN:
		message = fmt.Sprintf("unknown token [ %s ] at line:  %d", params[1].(string), params[0].(int))
		break
	case UNEXPECTED_EOF:
		message = fmt.Sprintf("unexpected end of file:  %s", params[0].(string))
		break
	default:
		message = "error"
	}

	return Foul{t, message}
}