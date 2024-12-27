package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			m.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}
