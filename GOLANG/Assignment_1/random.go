package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := make(chan int)
	go RandNumber(number)
	n := <-number
	fmt.Println("Random Number is :: ", n)
}

func RandNumber(number chan<- int) {
	number <- rand.Intn(50)
}
