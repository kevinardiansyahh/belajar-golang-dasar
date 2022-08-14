package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Kevin", "Ardiasnyah"
}

func getFName() (firstName, lastName string) {
	firstName = "Kevin"
	lastName = "Ardi"
	return
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func getGoodbye(name string) string {
	return "Bye" + name
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're Blocked")
	} else {
		fmt.Println("You're Registered")
	}
}

//func blackLiistAdmin(name string) bool {
//	return name == "admin"
//}
//
//func blackLiistRoot(name string) bool {
//	return name == "root"
//}

func factorialoop(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func factoriaRecursive(num int) int {
	if num == 1 {
		return 1
	}
	return num * factoriaRecursive(num-1)
}

func main() {
	//function
	sayHello()
	//parameterize function
	sayHelloTo("Kevin", "Hidayat")
	//returning value function
	result := getHello("Kevin")
	fmt.Println(result)
	//returning multiple value
	firstName, lastName := getFullName()
	//ignoring some return value
	//firstName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	//function retrurn namedValue
	fName, lName := getFName()
	fmt.Println(fName)
	fmt.Println(lName)

	//function variadic
	fmt.Println(sumAll(10, 10, 10))
	slice := []int{123, 123, 3, 2}
	total := sumAll(slice...)
	fmt.Println(total)

	//function as value
	//digunakan jika nantinya membutuhkan sebuah function yg parameternya function lain
	sayGoodBye := getGoodbye
	hasil := sayGoodBye("Kevin")
	fmt.Println(hasil)

	//function as parameter
	sayHelloWithFilter("Kevin", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//function anonymous -> membuat function langsung tanpa namanya
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("Kevin", blacklist)
	registerUser("admin", blacklist)

	//function recursive
	fmt.Println(factorialoop(5))
	fmt.Println(factoriaRecursive(5))
}
