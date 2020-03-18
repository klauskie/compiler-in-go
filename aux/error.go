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
	UNKNOWN_TOKEN = uint8(1)
)

func NewFoul(t uint8, params ...interface{}) Foul {
	var message string
	switch t {
	case UNKNOWN_TOKEN:
		message = fmt.Sprintf("unknown token at line:  %d", params[0].(int))
		break
	default:
		message = "error"
	}

	return Foul{t, message}
}