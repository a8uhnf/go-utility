package stack

import (
	"log"
	"sync"
)

const (
	MaxedoutStackError = "Stack maxed out. Unable to push"
	EmptyStackError    = "Stack empty. Unable to pop."
)

// Stack struct keep tracks of top element, stack size and lock for concurrent safe
type Stack struct {
	top     *stackElement
	size    int
	maxSize int
	lock    sync.Mutex
}

type stackElement struct {
	value interface{}
	prev  *stackElement
}

// Len returns the length of stack.
func (s *Stack) Len() int {
	return s.size
}

// SetSize sets the stack size.
func (s *Stack) SetSize(size int) {
	s.maxSize = size
}

// Top returns the top element of stack.
func (s *Stack) Top() interface{} {
	if s.Len() > 0 {
		return s.top.value
	}
	return nil
}

// Pop remove top element of stack and returns it.
func (s *Stack) Pop() interface{} {
	if s.Len() <= 0 {
		log.Println(EmptyStackError)
		return nil
	}
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
	if s.Len() >= s.maxSize {
		log.Println(MaxedoutStackError)
		return
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top = &stackElement{
		value: v,
		prev:  s.top,
	}
	s.size++
}
