package main

import "fmt"

func main() {

	defer firstrun()
	secondrun()

	fmt.Println(div(3, 0))
	fmt.Println(div(6, 3))
	dempanic()
}

func firstrun()  { fmt.Println("I executed first") }
func secondrun() { fmt.Println("I executed second") }

func div(n1, n2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	return (n1 / n2)
}

func dempanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("panic")
}
