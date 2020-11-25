package main

import (
	"fmt"
)

func main() {
	var i float32 = 10.75
	fmt.Printf("Value = %v and Type is %T \n", i, i)

	var j int
	j = int(i)
	fmt.Printf("Value = %v and Type is %T \n", j, j)
}
