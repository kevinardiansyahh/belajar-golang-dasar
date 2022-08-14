package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr." + man.Name
}

func main() {
	kevin := Man{"Kevin"}
	kevin.married()
	fmt.Println(kevin.Name)
}
