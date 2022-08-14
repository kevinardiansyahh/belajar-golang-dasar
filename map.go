package main

import "fmt"

func main() {
	var profile = map[string]string{
		"nama":   "Kevin",
		"alamat": "Jakarta",
	}
	fmt.Println(profile)         // map[alamat:Jakarta nama:Kevin]
	fmt.Println(profile["nama"]) // Kevin
	profile["nama"] = "Ardiansyah"
	fmt.Println(profile) // map[alamat:Jakarta nama:Ardiansyah]

	//tambah baru
	profile["job"] = "Developer"
	fmt.Println(profile) // map[alamat:Jakarta job:Developer nama:Ardiansyah]

	//map function (len, make, delete, map[key])
	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Kevin"
	book["ups"] = "salah"
	fmt.Println(book)      // map[author:Kevin title:Belajar Golang ups:salah]
	fmt.Println(len(book)) // 3

	delete(book, "ups")

	fmt.Println(book)      // map[author:Kevin title:Belajar Golang]
	fmt.Println(len(book)) // 2
}
