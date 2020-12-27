package main

import "fmt"

func main() {

	var name string = "febin"
	const a float64 = 5.1258
	pass := true
	x := 10

	fmt.Println(len(name))
	fmt.Println(name + "favas")

	fmt.Printf("%f \n", a)
	fmt.Printf("%.2f \n", a)
	fmt.Printf("%T \n", name)
	fmt.Printf("%t \n", name)
	fmt.Printf("%t \n", pass)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 8)
	fmt.Printf("%c \n", 67)
	fmt.Printf("%x \n", 12)
	fmt.Printf("%e \n", 5.2)
}
