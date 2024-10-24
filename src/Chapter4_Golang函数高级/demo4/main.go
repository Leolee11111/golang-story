package main

import "fmt"

// 定义一个函数类型
type predicateFunc func(int) bool

// 定义一个filter函数: 输入数据的切片和一个谓词函数，返回满足条件的切片
func filter(arr []int, f predicateFunc) []int {
    result := make([]int, 0)
    for _, v := range arr {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

// 定义一个简单的谓词函数: 判断是否为偶数
func isEven(x int) bool {
    return x % 2 == 0
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6}
    evenArr := filter(arr, isEven)
    fmt.Println(evenArr) // 输出：[2, 4, 6]
}