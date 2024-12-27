package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int32
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			fmt.Println(atomic.LoadInt32(&counter))
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}
