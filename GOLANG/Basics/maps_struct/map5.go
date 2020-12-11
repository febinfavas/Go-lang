package main

import "fmt"

func main() {

	Employee := map[int]string{
		100: "febin",
		101: "favas",
		102: "murshid",
		103: "shibil",
	}

	Ep := Employee
	delete(Ep, 103)
	fmt.Println(Ep)
	fmt.Println(Employee)
}
