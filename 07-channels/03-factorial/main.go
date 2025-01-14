package main

import "fmt"

func main() {
	c := generate(4, 10, 8)
	f := factorial(c)
	for n := range f {
		fmt.Println("The factorial", n)

	}

}

func generate(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {

		for v := range in {
			total := 1
			for i := v; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
