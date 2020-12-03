package main

import (
	"fmt"
)

func main() {
	student := [5]string{"febin", "favas", "rizwan", "shibil", "ashiq"}
	marks := [5]string{"97", "98", "99", "96", "95"}

	var table [5][2]string

	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			table[i][0] = student[i]
			table[i][1] = marks[i]
		}
	}

	fmt.Println("Table :", table)

}
