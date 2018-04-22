package queue

import (
	"log"
	"sync"
)

// Queue struct keeps track of top element, size, last element of queue
type Queue struct {
	top     *queueElement
	size    int
	maxSize int
	lock    sync.Mutex
	last    *queueElement
}
type queueElement struct {
	value interface{}
	next  *queueElement
	prev  *queueElement
}

// Len returns the length of queue
func (q *Queue) Len() int {
	return q.size
}

// SetSize sets the maximum size of Queue
func (q *Queue) SetSize(size int) {
	q.maxSize = size
}

// Top function returns top elemetn of queue
func (q *Queue) Top() interface{} {
	if q.Len() > 0 {
		return q.top.value
	}
	return nil
}

// Push function insert element at the last of queue
func (q *Queue) Push(v interface{}) {
	if q.Len() >= q.maxSize {
		log.Println("Queue maxed out. Unable to push")
		return
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.Len() == 0 {
		q.top = &queueElement{
			value: v,
			next:  nil,
		}
		q.last = q.top
		q.size++
		return
	}
	q.size++
	tmp := q.last
	q.last = &queueElement{
		next:  q.last,
		value: v,
	}
	tmp.prev = q.last
}

// Pop function remove top element from the queue
func (q *Queue) Pop() interface{} {
	if q.Len() <= 0 {
		log.Println("Queue empty. Unable to pop.")
		return nil
	}
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
