package main

import "fmt"

func main() {

	defer firstrun()
	secondrun()
}

func firstrun()  { fmt.Println("I executed first") }
func secondrun() { fmt.Println("I executed second") }
