package googlesheett

import "testing"

func TestParseRange(t *testing.T) {
	firstNum, secondNum, err := ParseRange("3827:3899")
	if firstNum != 3827 {
		t.Error("firstNum err")
	}
	if secondNum != 3899 {
		t.Error("secondNum err")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestNewRange(t *testing.T) {
	str := NewRange(1, 2, "A")
	if str != "A1:A2" {
		t.Errorf("%+v not match A1:A2", str)
	}
}
