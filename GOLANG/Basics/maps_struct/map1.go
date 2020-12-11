package main

import "fmt"

func main() {

	StudentRoll := make(map[int]string)
	StudentRoll[10] = "Febin"
	StudentRoll[5] = "Favas"
	StudentRoll[12] = "Murshid"
	StudentRoll[15] = "Rizwan"
	StudentRoll[20] = "Shibil"

	fmt.Println(StudentRoll[12])
	fmt.Println(StudentRoll)

}
