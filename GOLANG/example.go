package main

import "fmt"

func main() {
	fmt.Println("hello world")

	const a int = 10
	x := 5
	y := 6
	fmt.Println("sum =", x+y-a)
}
