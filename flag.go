package main

import (
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your database host")
	var user *string = flag.String("user", "root", "put your database user")
	var password *string = flag.String("password", "", "put your database password")

	flag.parse()

	fmt.Println("host : ", host)
	fmt.Println("user : ", user)
	fmt.Println("password : ", password)
}