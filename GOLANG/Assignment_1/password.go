package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[A-Z]{1}[a-z]{4}[0-9]{2}")
	fmt.Println(re.FindString("kloudOne20"))
	fmt.Println(re.FindString("Kloud20"))
	fmt.Println(re.FindString("klouDone20"))
	fmt.Println(re.FindString("Alokk20"))
}
