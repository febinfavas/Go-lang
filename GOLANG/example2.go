package main

import "fmt"

func main() {

	x := 10

	fmt.Println(x)
	fmt.Println(&x)

	changevalue(&x)
	fmt.Println(x)
}

func changevalue(x *int) {
	*x = 7
}
