package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}
	fmt.Println("My Array is :: ", arr)
	fmt.Println(" ")
	fmt.Println("Zig Zag Matrix is :: ")
	zigzagMatrix(arr)
}

func min2(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func min3(a int, b int, c int) int {
	return min2(min2(a, b), c)
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func zigzagMatrix(matrix [][]int) {
	for line := 1; line <= 8; line++ {
		start_col := max(0, line-5)
		count := min3(line, (4 - start_col), 5)
		for j := 0; j < count; j++ {
			fmt.Print(matrix[min2(5, line)-j-1][start_col+j], " ")
		}
		fmt.Println(" ")
	}
}
