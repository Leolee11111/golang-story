package main

import "fmt"

func devide(a, b int) int {
	return a / b
}

func main () {
	res := devide(1, 0)
	fmt.Printf("devide(1, 0) = %d\n", res)
}