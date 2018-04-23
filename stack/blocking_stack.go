package stack

import (
	"fmt"
	"sync"
)

const (
	maxedoutBlockingStackError = "BlockingStack maxed out. Unable to push"
	emptyBlockingStackError    = "BlockingStack empty. Unable to pop."
)

// BlockingStack struct keep tracks of top element, Blockingstack size and lock for concurrent safe
type BlockingStack struct {
	top            *BlockingstackElement
	size           int
	maxSize        int
	lock           sync.Mutex
	popLock        sync.Mutex
	pushLock       sync.Mutex
	popBlock       chan int
	pushBlock      chan int
	pushBlockState bool
	popBlockState  bool
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

// Len returns the length of Blockingstack.
func (s *BlockingStack) Len() int {
	return s.size
}

// Top returns the top element of Blockingstack.
func (s *BlockingStack) Top() interface{} {
	if s.Len() <= 0 {
		return nil
	}
	return s.top.value
}

// Pop remove top element of Blockingstack and returns it.
func (s *BlockingStack) Pop() interface{} {
	s.popLock.Lock()
	fmt.Println("Hello World!!!----", s.Top())
	fmt.Println("Hello World!!!----", s.Len())
	defer s.popLock.Unlock()
	if s.Len() <= 0 {
		s.popBlockState = true
		<-s.popBlock
		s.popBlockState = false
	}
	ret := s.top.value
	s.top = s.top.prev
	s.size--
	if s.pushBlockState {
		fmt.Println("exting push-block-state")
		s.pushBlock <- 1
		fmt.Println("Push block state...")
	}
	return ret
}

// Push add element to Blockingstack
func (s *BlockingStack) Push(v interface{}) {
	s.pushLock.Lock()
	defer s.pushLock.Unlock()
	if s.Len() >= s.maxSize {
		s.pushBlockState = true
		fmt.Println("entering push-block-state")
		<-s.pushBlock
		fmt.Println("entering push-block-state")
		s.pushBlockState = false
	}
	s.top = &BlockingstackElement{
		value: v,
		prev:  s.top,
	}
	s.size++
	if s.popBlockState {
		s.popBlock <- 1
	}
}
