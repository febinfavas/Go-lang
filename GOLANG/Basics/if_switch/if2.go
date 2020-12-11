package main

import "fmt"

func main() {

	age := 20

	if age >= 18 && age <= 30 {
		fmt.Println("You are youth!!")
	} else {
		fmt.Println("You are midle aged ")
	}

	if true {
		fmt.Println("The value printed")
	}

	if false {
		fmt.Println("The value printed")
	}

}
