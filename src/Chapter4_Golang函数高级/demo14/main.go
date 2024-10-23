package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
	return "Wang!"
}

type Cat struct {}

func (c Cat) Speak() string {
	return "Miao!"
}

func main() {
	var s Speaker
	s = Dog{}
	fmt.Println(s.Speak()) // 输出：Wang!

    s = Cat{}
    fmt.Println(s.Speak()) // 输出：Miao!
}