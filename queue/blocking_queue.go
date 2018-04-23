package queue

import (
	"sync"
)

// BlockingQueue struct keeps track of top element, size, last element of BlockingQueue
type BlockingQueue struct {
	top            *BlockingQueueElement
	size           int
	maxSize        int
	lock           sync.Mutex
	popLock        sync.Mutex
	pushLock       sync.Mutex
	popBlock       chan int
	pushBlock      chan int
	pushBlockState bool
	popBlockState  bool
	last           *BlockingQueueElement
}

// BlockingQueueElement contains queue's element value, next, previous element
type BlockingQueueElement struct {
	value interface{}
	next  *BlockingQueueElement
	prev  *BlockingQueueElement
}

// NewBlockingQueue initialize blocking queue.
func NewBlockingQueue() *BlockingQueue {
	ret := &BlockingQueue{
		popBlock:  make(chan int),
		pushBlock: make(chan int),
	}
	return ret
}

// Len returns the length of BlockingQueue
func (q *BlockingQueue) Len() int {
	return q.size
}

// Top function returns top elemetn of BlockingQueue
func (q *BlockingQueue) Top() interface{} {
	if q.Len() > 0 {
		return q.top.value
	}
	return nil
}

// Push function insert element at the last of BlockingQueue
func (q *BlockingQueue) Push(v interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.Len() == 0 {
		q.top = &BlockingQueueElement{
			value: v,
			next:  nil,
		}
		q.last = q.top
		q.size++
		return
	}
	q.size++
	tmp := q.last
	q.last = &BlockingQueueElement{
		next:  q.last,
		value: v,
	}
	tmp.prev = q.last
}

// Pop function remove top element from the BlockingQueue
func (q *BlockingQueue) Pop() interface{} {
	if q.Len() > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()
		ret := q.top.value
		q.top.next = nil
		q.top = q.top.prev
		q.size--
		return ret
	}
	return nil
}
