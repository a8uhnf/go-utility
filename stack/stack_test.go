package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)
	if s.Len() != 0 {
		t.Fail()
	}
	s.Push(2)
	if s.Top().(int) != 2 {
		t.Fail()
	}
	s.Push(3.2)
	if int(s.Top().(float64)) != 3 {
		t.Fail()
	}
}
