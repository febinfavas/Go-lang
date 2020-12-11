package main

import "fmt"

func main() {
	x := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13}
	fmt.Println("My Map is :: ", x)
	fmt.Println("Slice from 1 to 4 :: ", x["A"])
}
