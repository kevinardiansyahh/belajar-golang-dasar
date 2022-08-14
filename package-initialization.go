package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

// import _"belajar-golang-dasar/database" -> use blank identifier to only run init function in package

func main() {
	fmt.Println(database.GetDatabase())
}
