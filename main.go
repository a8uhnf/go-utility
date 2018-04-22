package main

import (
	"fmt"

	"github.com/a8uhnf/go-utility/queue"
	"github.com/a8uhnf/go-utility/stack"
)

func main() {
	fmt.Println("Hello World!!!")
	s := new(stack.ConcurrentStack)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Top())

	fmt.Println("----------------------")

	q := new(queue.ConcurrentQueue)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Top())
	fmt.Println("----------------------")
	var c chan int
	c = make(chan int, 1)
	c <- 1
	<-c
}
