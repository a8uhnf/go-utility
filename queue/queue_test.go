package queue

import (
	"fmt"
	"testing"
)

func testQueuePush(t *testing.T) {
	q := new(Queue)
	q.Push(2)
	fmt.Println("------------")
	if q.Top().(int) != 2 {
		t.Fail()
	}
}
