package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {

	var ms *myStruct
	ms = new(myStruct)
	fmt.Println((*ms))
	ms.foo = 42
	fmt.Println((*ms))
}
