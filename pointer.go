package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	//pembuktian bahwa golang itu secara default pass by value, bukan pass by reference
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)

	//pass by reference menggunakan pointer
	address3 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address4 := &address3 //address4 akan merefer ke memory yang sama dengan address3

	address4 = &Address{"Malang", "Jawa Tengah", "Indonesia"} //address4 akan membuat memory
	//baru & memisahkan diri dengan memory di address3

	//memindahkan seluruh variable yg merefer ke memory Address sebelumnya ke memory address yg baru
	//address3 juga akan mengarah ke memory address yang baru
	*address4 = Address{"Malang", "Jawa Tengah", "Indonesia"}

	fmt.Println(address3)
	fmt.Println(address4)

	//jika ingin membuat pointer baru tanpa instansiasi data awal (data dikosongkan)
	var alamat1 *Address = new(Address)
	fmt.Println(alamat1)

	//pointer diparameter function
	//jika data yg dikirimkan ke function ingin ikut berubah
	//maka bisa dibuat variable parameternya as pointer
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
