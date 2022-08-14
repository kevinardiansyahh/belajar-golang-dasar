package main

import "fmt"

type hasName interface {
	getName() string
}

func sayHai(hasName hasName) {
	fmt.Println("Haiii", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (ani Animal) getName() string {
	return ani.Name
}

func main() {
	var kevin Person
	kevin.Name = "Kevin"
	sayHai(kevin)
	fmt.Println(kevin)
}
