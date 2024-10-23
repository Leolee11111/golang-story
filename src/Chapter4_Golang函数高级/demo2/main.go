package main

import "fmt"

// 定义一个函数类型
type opFunc func(int, int) int

// 定义加法函数
func add(a int, b int) int {
    return a + b
}

// 定义减法函数
func sub(a int, b int) int {
    return a - b
}

// 定义乘法函数
func mul(a int, b int) int {
    return a * b
}

// 定义除法函数
func div(a int, b int) int {
    if b == 0 {
        fmt.Println("Error: Division not be zero")
        return 0
    }
    return a / b
}

// 定义一个高阶函数，接受函数作为参数
func operate(a int, b int, op opFunc) int {
    return op(a, b)
}

func main() {
    a, b := 10, 5

    fmt.Println("Add:", operate(a, b, add))         // 输出：Add: 15
    fmt.Println("Subtract:", operate(a, b, sub)) // 输出：Subtract: 5
    fmt.Println("Multiply:", operate(a, b, mul)) // 输出：Multiply: 50
    fmt.Println("Divide:", operate(a, b, div))     // 输出：Divide: 2
}