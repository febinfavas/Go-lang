package main

import "fmt"

func main() {

	var twoD_array [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD_array[i][j] = i + j + 2
		}
	}
	fmt.Println("2d: ", twoD_array)
}
