package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case koreanAge:
		return false
	default:
		return true
	}
}
func main() {
	fmt.Println(canIDrink(16))
}
