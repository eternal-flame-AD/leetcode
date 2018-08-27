package lispexp

import "testing"

func TestStringSplit(t *testing.T) {
	res := splitExprstrParts("add a (add a 1) 2 3 f")
	if len(res) != 6 {
		t.Fail()
	}
}

func TestExprEval(t *testing.T) {
	res := ConsumeLispExpString("(let x 2 (mult x 5))", nil).eval()
	if res != 10 {
		t.Fail()
	}
}
