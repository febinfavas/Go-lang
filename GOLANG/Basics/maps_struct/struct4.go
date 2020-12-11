//Embedded Struct
package main

import "fmt"

type Student struct {
	name   string
	rollno int
}

type School struct {
	Student
	school string
	place  string
}

func main() {
	abc := School{}
	abc.name = "febin"
	abc.rollno = 10
	abc.school = "guidance"
	abc.place = "edakkara"

	fmt.Println(abc)
	fmt.Println(abc.name)
}
