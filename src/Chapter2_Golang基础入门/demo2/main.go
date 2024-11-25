package main

import "fmt"

// 全局变量的声明
var globalVar = "I am a global variable"

func main() {

	// 局部变量的声明
	localVar := "I am a local variable"

	fmt.Println(globalVar)
	fmt.Println(localVar)
}
