package main

import "fmt"

func main() {

	num := 5.0
	if num < 0 {
		fmt.Println("negative number")
	} else if num > 0 {
		fmt.Println("positive number")
	} else {
		fmt.Println("zero")
	}

}
