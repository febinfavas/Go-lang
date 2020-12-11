package main

import "fmt"

func main() {

	mark := 111

	switch mark {
	case 40:
		fmt.Println("failed")
	case 50:
		fmt.Println("passed")
	case 90:
		fmt.Println("Distinction")
	default:
		fmt.Println("Not Valid")

	}
}
