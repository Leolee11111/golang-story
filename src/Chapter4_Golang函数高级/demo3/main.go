package main

import "fmt"

// 定义一个函数类型
type transFunc func(int) int

// 定义一个map函数: 输入数据的切片和一个变换函数，返回变换后的切片
func mapFunc(arr []int, f transFunc) []int {
    result := make([]int, len(arr))
    for i, v := range arr {
        result[i] = f(v)
    }
    return result
}

// 定义一个简单的变换函数: y = x^2
func square(x int) int {
    return x * x
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    squaredRes := mapFunc(arr, square)
    fmt.Println(squaredRes) // 输出：[1 4 9 16 25]
}