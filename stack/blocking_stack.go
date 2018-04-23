package stack

import (
	"sync"
	"fmt"
)

const (
	maxedoutBlockingStackError = "BlockingStack maxed out. Unable to push"
	emptyBlockingStackError    = "BlockingStack empty. Unable to pop."
)

// BlockingStack struct keep tracks of top element, Blockingstack size and lock for concurrent safe
type BlockingStack struct {
	top        *BlockingstackElement
	size       int
	maxSize    int
	lock       sync.Mutex
	block      chan int
	blockState bool
}

// BlockingstackElement represent stack element. Contains value and previous element.
type BlockingstackElement struct {
	value interface{}
	prev  *BlockingstackElement
}

// SetSize sets the size of stack.
func (s *BlockingStack) SetSize(sz int) {
	s.maxSize = sz
}

// BlockStack blocks the stack, when its empty or full.
func (s *BlockingStack) BlockStack() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.blockState = true
	<-s.block
	s.blockState = false
}

// AddStack unblock the block state of blocking stack.
func (s *BlockingStack) AddStack() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.block <- 1
}

// Len returns the length of Blockingstack.
func (s *BlockingStack) Len() int {
	return s.size
}

// Top returns the top element of Blockingstack.
func (s *BlockingStack) Top() interface{} {
	if s.Len() <= 0 {
		s.BlockStack()
	}
	return s.top.value
}

// Pop remove top element of Blockingstack and returns it.
func (s *BlockingStack) Pop() interface{} {
	if s.Len() <= 0 {
		s.BlockStack()
	}
	if s.Len() == s.maxSize && s.blockState {
		s.AddStack()
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	ret := s.top.value
	s.top = s.top.prev
	s.size--
	return ret
}

// Push add element to Blockingstack
func (s *BlockingStack) Push(v interface{}) {
	if s.Len() == 0 && s.blockState {
		s.AddStack()
	}
	if s.Len() >= s.maxSize {
		fmt.Println("YYYYYYYYYYYYY")
		s.BlockStack()
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top = &BlockingstackElement{
		value: v,
		prev:  s.top,
	}
	s.size++
}
