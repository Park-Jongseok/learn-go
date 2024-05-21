package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch age {
	case 16:
		return false
	default:
		return true
	}
}
func main() {
	fmt.Println(canIDrink(16))
}
