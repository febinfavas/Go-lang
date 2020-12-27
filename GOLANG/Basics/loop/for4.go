package main

import "fmt"

func main() {

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, value := range a {
		fmt.Println(value)
	}

	b := "febin favas"
	for _, value := range b {
		fmt.Println(string(value))
	}
}
