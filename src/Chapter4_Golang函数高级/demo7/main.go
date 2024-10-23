package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// 调用外部函数，返回闭包函数
    pos := adder()

    // 多次调用闭包函数
    fmt.Println(pos(1)) // 输出：1
    fmt.Println(pos(2)) // 输出：3
    fmt.Println(pos(3)) // 输出：6
}