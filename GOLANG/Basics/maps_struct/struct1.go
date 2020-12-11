package main

import "fmt"

type Rectangle struct {
	height float64
	width  float64
}

func main() {
	rect := Rectangle{height: 10, width: 5}

	fmt.Println("Height = ", rect.height)
	fmt.Println("Width = ", rect.width)
	fmt.Println("Area of Rectagle  = ", rect.area())
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
