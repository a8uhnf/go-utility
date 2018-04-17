package queue

import (
	"testing"
)

func TestQueuePush(t *testing.T) {
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
}
