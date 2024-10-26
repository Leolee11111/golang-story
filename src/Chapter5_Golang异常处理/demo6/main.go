package main

import "fmt"

// 定义一个函数，使用 panic 引发异常
func devide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	res := devide(10, 0)
	fmt.Printf("devide(10, 0) = %d\n", res)
}