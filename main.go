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


	bs := new(stack.BlockingStack)
	bs.SetSize(10)
	var chk int
	go func() {
		chk = bs.Pop().(int)
		fmt.Println(chk)
	}()
	bs.Push(2)
	fmt.Println(chk)

	fmt.Scanf("%d",&chk)
}
