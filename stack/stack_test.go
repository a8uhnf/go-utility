package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := new(Stack)
	if s.Len() != 0 {
		t.Fail()
	}
	s.Push(2)
	if s.Top().(int) != 2 {
		t.Fail()
	}
}
