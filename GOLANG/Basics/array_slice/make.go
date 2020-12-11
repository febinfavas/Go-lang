package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)
	fmt.Println(slice)
	fmt.Println("Length ", len(slice))
	fmt.Println("Capacity ", cap(slice))

	b := []int{}
	fmt.Println(b)
	fmt.Println("Length ", len(b))
	fmt.Println("Capacity ", cap(b))

	b = append(b, 1)
	fmt.Println(b)
	fmt.Println("Length ", len(b))
	fmt.Println("Capacity ", cap(b))

	b = append(b, 2, 3, 4, 5)
	fmt.Println(b)
	fmt.Println("Length ", len(b))
	fmt.Println("Capacity ", cap(b))

}
