package main

import "fmt"

func main() {
	// 1.定义匿名函数并立即执行
	result1 := func (a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println("result1: ", result1)

	// 2.将匿名函数赋值给变量
	addFunc := func (a int, b int) int {
		return a + b
	}
	result2 := addFunc(1, 2)
	fmt.Println("result2: ", result2)

	// 3.匿名函数作为参数传递
	result3 := calc(1, 2, func(a int, b int) int {
		return a + b
	})
	fmt.Println("result3: ", result3)
}

func calc(a int, b int, op func(int, int) int) int {
	return op(a, b)
}