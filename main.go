package main

import "fmt"

type person struct {
	name         string
	age          int
	favoriteFood []string
}

func main() {
	favoriteFood := []string{"kimchi", "ramen"}
	p := person{name: "james", age: 32, favoriteFood: favoriteFood}
	fmt.Println(p)
}
