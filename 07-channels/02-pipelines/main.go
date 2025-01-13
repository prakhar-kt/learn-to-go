package main

import "fmt"

func main() {
	c := gen(6, 5, 4)
	b := square(c)
	for n := range b {
		fmt.Println("The output chan", n)
	}

}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)

	}()
	return out
}

func square(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}
