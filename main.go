package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	// Array
	fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 0
}
func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
