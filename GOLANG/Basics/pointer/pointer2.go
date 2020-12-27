package main

import "fmt"

func main() {

	var x int = 10
	var y *int = &x

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(y)
	fmt.Println(*y)

	x = 25

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(y)
	fmt.Println(*y)
}
