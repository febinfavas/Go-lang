package main

import (
	"fmt"
)

func main() {
	student := [5]string{"febin", "favas", "rizwan", "shibil", "ashiq"}
	marks := [5]int{97, 98, 99, 96, 95}
	fmt.Println("Students :", student)
	fmt.Println("Marks :", marks)

	slice1 := student[0:0]
	slice2 := student[0:]
	slice3 := student[:0]
	slice4 := student[:len(student)]
	slice5 := student[0:5]
	slice6 := marks[2:5]

	fmt.Println(slice1, "\n", slice2, "\n", slice3, "\n", slice4, "\n", slice5, "\n", slice6)

}
