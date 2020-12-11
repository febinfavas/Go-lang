package main

import "fmt"

func main() {

	Employee := map[int]string{
		100: "febin",
		101: "favas",
		102: "murshid",
		103: "shibil",
	}

	_, ok := Employee[102]
	fmt.Println(ok)

	pop, ok := Employee[103]
	fmt.Println(pop, ok)

	pop1, ok := Employee[104]
	fmt.Println(pop1, ok)

	fmt.Println(len(Employee))
}
