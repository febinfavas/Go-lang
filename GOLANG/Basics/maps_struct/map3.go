package main

import "fmt"

func main() {

	Employee := map[int]string{
		100: "febin",
		101: "favas",
		102: "murshid",
		103: "shibil",
	}
	fmt.Println(Employee)
	Employee[104] = "rizwan"
	Employee[105] = "anjah"
	fmt.Println(Employee)
	delete(Employee, 104)
	fmt.Println(Employee)
}
