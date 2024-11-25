package main

import "fmt"

func main() {
	// 显式声明
	var age int
	fmt.Println(age)

	// 隐式声明
	name := "hello world"
	fmt.Println(name)

	// 批量声明
	var (
		price int    = 25
		thing string = "box"
	)
	fmt.Println(price)
	fmt.Println(thing)
}
