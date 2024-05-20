package main

import (
	"fmt"
	"hello"
)

func main() {
	// Error
	// const constant string = "Hello, World!"
	// constant = "World! Hello"

	var message string = hello.Hello()
	message = "World! Hello"
	fmt.Println(message)
}
