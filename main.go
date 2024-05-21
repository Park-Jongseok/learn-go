package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(&a, b)
	fmt.Println(a, *b)

	*b = 3
	fmt.Println(a)
}
