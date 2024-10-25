package main

import "fmt"

func main() {
    fmt.Println("Start of main function")

    defer fmt.Println("Deferred call 1")
    defer fmt.Println("Deferred call 2")
    defer fmt.Println("Deferred call 3")

    fmt.Println("End of main function")
}