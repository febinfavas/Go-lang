package main

import "fmt"

func main() {

	var array1 [5]int
	fmt.Println("First Array    :", array1)

	array1[4] = 100
	fmt.Println("Setting Element:", array1)
	fmt.Println("Getting Elemrnt:", array1[4])

	array1[2] = 50
	fmt.Println("Setting Element:", array1)
	fmt.Println("Getting Elemrnt:", array1[2])

	fmt.Println("Length of Array:", len(array1))

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Second Array   :", array2)

}
