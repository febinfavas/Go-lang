package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i string = "10"
	fmt.Printf("Value = %v and Type is %T \n", i, i)

	j, _ := strconv.Atoi(i)
	fmt.Printf("Value = %v and Type is %T \n", j, j)
}
