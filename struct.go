package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Haii ", name, "My Name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huuu from", a.Name)
}

func main() {
	var kevin Customer
	kevin.Name = "Kevin Ardiansyah Hidayat"
	kevin.Address = "Cianjur"
	kevin.Age = 26

	fmt.Println(kevin)
	fmt.Println(kevin.Name)

	joko := Customer{
		Name:    "Joko",
		Address: "Jawa",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Bandung", 28}

	fmt.Println(budi)

	//struct function
	//function yang seakan-akan dikhususkan untuk struct tertentu
	budi.sayHi("Ardi")
	kevin.sayHuu()
}
