package main

import (
	"fmt"
)

func main() {
	arr := []int{89, 12, 56, 74, 32}
	fmt.Println("Unsorted Array ", arr)
	quicksort(arr)
	fmt.Println("Merge sort ", arr)
}

func quicksort(arr []int) []int {
	if len(arr) > 1 {
		pivot := len(arr) / 2
		var left_side = []int{}
		var right_side = []int{}
		for i := range arr {
			val := arr[i]
			if i != pivot {
				if val < arr[pivot] {
					left_side = append(left_side, val)
				} else {
					right_side = append(right_side, val)
				}
			}
		}
		quicksort(left_side)
		quicksort(right_side)
		var merge []int = append(append(append([]int{}, left_side...), []int{arr[pivot]}...), right_side...)
		for j := 0; j < len(arr); j++ {
			arr[j] = merge[j]
		}
	}
	return arr
}
