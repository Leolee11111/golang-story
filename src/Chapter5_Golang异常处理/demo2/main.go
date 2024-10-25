package main

import (
	"errors"
	"fmt"
)

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	res, err := devide(1, 0)
	if err != nil {
		fmt.Printf("devide(1, 0) error: %s\n", err)
	} else {
		fmt.Printf("devide(1, 0) = %d\n", res)
	}
}