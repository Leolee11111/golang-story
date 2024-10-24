package main

import "fmt"

// 定义一个结构体
type Person struct {
    Name string
    Age  int
}

// 定义一个方法
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    // 创建一个结构体实例
    p := Person{Name: "Alice", Age: 30}
    p.Greet() // 输出：Hello, my name is Alice and I am 30 years old.
}