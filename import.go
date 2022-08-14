package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Kevin")
	// helper.sayGoodbye("Kevin") -> error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) -> error
}
