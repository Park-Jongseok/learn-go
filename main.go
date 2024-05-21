package main

import "fmt"

func main() {
	stringMap := map[string]string{"name": "foo", "bar": "bar"}
	for key, value := range stringMap {
		fmt.Println(key, value)
	}
}
