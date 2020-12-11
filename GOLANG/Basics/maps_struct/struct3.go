package main

import "fmt"

func main() {
	rect := struct{ height, width int }{10, 5}

	fmt.Println(rect)
	arect := rect
	arect.height = 15
	arect.width = 20
	fmt.Println(arect)
}
