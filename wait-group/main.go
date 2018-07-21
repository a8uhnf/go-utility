package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello World!!!")
	var wg sync.WaitGroup
	c := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(a int) {
			defer wg.Done()
			<-c
			fmt.Println(a)
		}(i)
	}

	time.Sleep(time.Second * 3)
	close(c)
	wg.Wait()
}
