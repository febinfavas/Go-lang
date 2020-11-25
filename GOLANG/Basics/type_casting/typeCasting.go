package main

import (
	"fmt"
)

func main() {
	var x, y int = 10, 4
	var z float64 = float64(x) / float64(y)
	fmt.Printf("x = %v -> Type is %T \ny = %v  -> Type is %T \nz = %v  -> Type is %T \n", x, x, y, y, z, z)

	var i, j int = 10, 4
	var k float64 = float64(x / y)
	fmt.Printf("\nx = %v -> Type is %T \ny = %v  -> Type is %T \nz = %v  -> Type is %T \n", i, i, j, j, k, k)
}
