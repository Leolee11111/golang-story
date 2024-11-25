package main

import "fmt"

func main() {
	name := "Alice"
	age := 30
	message := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(message)

	number := 42
	fmt.Printf(fmt.Sprintf("|%05d|\n", number))
}
