package database

import "fmt"

var connection string

func init() {
	fmt.Println("init dipanggil")
	connection = "PostgreSQL"
}

func GetDatabase() string {
	return connection
}
