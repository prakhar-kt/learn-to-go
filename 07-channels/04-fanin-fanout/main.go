package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := generate()
	c1 := factorial(nums)
	c2 := factorial(nums)
	c3 := factorial(nums)
	c5 := factorial(nums)
	c4 := factorial(nums)
	c6 := factorial(nums)
	c7 := factorial(nums)
	c8 := factorial(nums)
	c9 := factorial(nums)
	c10 := factorial(nums)

	for n := range merge(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
		fmt.Println(n)
	}

}

func generate() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 1; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
