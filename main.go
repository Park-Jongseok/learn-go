package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	// Array
	fmt.Println(numbers)
	return 0
}
func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
