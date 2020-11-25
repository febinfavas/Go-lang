package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	fmt.Printf("Value = %v and Type is %T \n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("Value = %v and Type is %T \n", j, j)
}
