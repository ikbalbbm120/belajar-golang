package main

import "fmt"

func main() {
	var name string 

	name = "ikbal"
	fmt.Println(name)

	//tanpa perlu menggunakan var
	country := "singapore"
	fmt.Println(country)

	//menggabungkan var dalam 1 kurung
	var (
		firstName = "ikbal"
		lastName = "bahrudin"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}