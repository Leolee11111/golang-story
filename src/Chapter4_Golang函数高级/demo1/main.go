package main

import "fmt"

// 定义一个简单的函数: 函数名add 参数列表(a, b int) 返回值类型int
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Println("Result:", result) // 输出：Result: 7
}
