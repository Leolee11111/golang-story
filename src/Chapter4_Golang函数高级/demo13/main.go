package main

type Animal struct {
	Name string 
}

func (a Animal) Speak() string {
	return "..."
}

type Dog struct {
	Animal
	Age int
}

func (d Dog) Speak() string {
	return "Wang!"
}

func main() {
	dog := Dog{Animal: Animal{Name: "little pig"}, Age: 3}
	println(dog.Name)
	println(dog.Speak())
}