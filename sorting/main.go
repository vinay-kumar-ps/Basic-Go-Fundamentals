package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 88, 65, 64}
	xs := []string{"kunu", "vava", "two", "a", "b"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
