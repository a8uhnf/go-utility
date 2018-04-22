package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := new(Queue)
	q.Push(2)
	if q.Top().(int) != 2 {
		t.Fail()
	}
	q.Push(3)
	if q.Top().(int) != 2 {
		t.Fail()
	}
	top := q.Pop()
	if top != 2 {
		t.Fail()
	}
	if q.Top() != 3 {
		t.Fail()
	}
	top = q.Pop()
	if top != 3 {
		t.Fail()
	}
	if q.Len() != 0 {
		t.Fail()
	}
}
