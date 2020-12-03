package main

import (
	"fmt"
)

func main() {
	ID := []int{11, 22, 33, 44, 55}

	for i, value := range ID {
		fmt.Printf("%v --> %v \n", i, value)
	}

	fmt.Println(ID)
	slice1 := ID[0:2]
	fmt.Printf("First Slice is %v \n", slice1)
	slice2 := ID[3:5]
	fmt.Printf("Second Slice is %v \n", slice2)

	copy(slice1, slice2)
	fmt.Printf("Slice1 after coping is %v \n", slice1)

	slice3 := append(ID, 6, 7, 8, 9, 10)
	fmt.Printf("Third Slice is %v \n", slice3)
}
