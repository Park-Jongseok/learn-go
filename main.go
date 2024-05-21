package main

import "fmt"

func main() {
	names := []string{"John", "Paul", "George", "Ringo", "Stuart"}
	names = append(names, "Ricky")
	fmt.Println(names)
}
