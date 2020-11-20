package main

import "fmt"

func main() {
	ID := []int{1, 2, 3, 4, 5}

	for i, value := range ID {
		fmt.Println(value, i)
	}

	fmt.Println(ID)
	slice1 := ID[0:2]
	fmt.Println(slice1)
	slice2 := ID[3:5]
	fmt.Println(slice2)

	copy(slice1, slice2)
	fmt.Println(slice1)

	slice3 := append(ID, 6, 7, 8, 9, 10)
	fmt.Println(slice3)
}
