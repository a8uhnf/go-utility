package stack

import (
	"sync"
)

type Stack struct {
	top  *stackElement
	size int
	lock sync.Mutex
}

type stackElement struct {
	value interface{}
	prev  *stackElement
}

// Len returns the length of stack
func (s Stack) Len() int {
	return s.size
}

// Top returns the top element of stack
func (s Stack) Top() interface{} {
	if s.Len() > 0 {
		return s.top.value
	}
	return nil
}

// Pop remove top value of stack and returns it
func (s *Stack) Pop() interface{} {
	if s.Len() > 0 {
		s.lock.Lock()
		defer s.lock.Unlock()
		ret := s.top.value
		s.top = s.top.prev
		s.size--
		return ret
	}
	return nil
}

// Push add element to stack
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top = &stackElement{
		value: v,
		prev:  s.top,
	}
	s.size++
}
