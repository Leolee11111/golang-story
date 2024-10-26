package main

import "fmt"


func main() {
	var i int = 1
	defer func() {
		fmt.Println("defer i = ", i)
	}()
	i++
}