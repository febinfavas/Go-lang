package main

import "fmt"

func main() {

	x, y := 20, 6
	fmt.Println(add(x, y))

	fmt.Println(fact(x))
}

func add(n1, n2 int) int {
	return (n1 + n2)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
