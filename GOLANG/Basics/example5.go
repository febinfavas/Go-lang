package main

import "fmt"

func main() {

	age := 10

	if age >= 18 {
		fmt.Println("You Can Vote!!")
	} else {
		fmt.Println("You Can't Vote!!")
	}

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
