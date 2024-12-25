package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{41, 7, 18, 9, 10, 11, 8, 7}
	xs := []string{"James", "Jenny", "Amy", "Sheila", "Matt", "aby"}
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("-----------")
	sort.Strings(xs)
	fmt.Println(xs)

}
