package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}
func main() {
	totalLength, upperName := lenAndUpper("mike")
	fmt.Println(totalLength, upperName)
}
