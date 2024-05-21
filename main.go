package main

import "fmt"

func main() {
	names := [5]string{"John", "Paul", "George", "Ringo", "Stuart"}
	// Error
	// names[5] = "Stuart"
	fmt.Println(names)
}
