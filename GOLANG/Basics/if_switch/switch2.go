package main

import "fmt"

func main() {

	var i interface{} = "febin"

	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is an string")
	default:
		fmt.Println("i is an other type")

	}
}
