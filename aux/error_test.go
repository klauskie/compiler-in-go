package aux

import "testing"

func TestNewFoul(t *testing.T) {
	f := NewFoul(UNKNOWN_TOKEN, 3, "|")
	if f.Type != UNKNOWN_TOKEN {
		t.Errorf("NewFoul; Expected Foul of type UNKNOWN (1); got %d", f.Type)
	}

	f = NewFoul(uint8(9), 3, "|")
	if f.Message != "error" {
		t.Errorf("NewFoul; Expected default Foul with message = error; got %s", f.Message)
	}
}
