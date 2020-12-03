package main

import (
	"errors"
	"fmt"
)

func main() {
	if result, err := divide(35, 0); err != nil {
		fmt.Println("Error Occured :: ", err)
	} else {
		fmt.Println("Result is :: ", result)
	}
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("denumerator should not be zero ... ")
	}
	return a / b, nil
}
