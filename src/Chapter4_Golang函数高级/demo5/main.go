package main

import "fmt"

// 定义一个函数类型
type reduceFunc func(int, int) int

// 定义一个reduce函数: 输入数据的切片、一个累积函数和一个初始值，返回累积的结果
func reduce(arr []int, f reduceFunc, initial int) int {
    result := initial
    for _, v := range arr {
        result = f(result, v)
    }
    return result
}

// 定义一个简单的累积函数
func sum(x, y int) int {
    return x + y
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    total := reduce(arr, sum, 0)
    fmt.Println(total) // 输出：15
}