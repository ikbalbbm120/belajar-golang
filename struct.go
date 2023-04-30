package main

import "fmt"

type costumer struct {
	Name, Address string
	Age    int
}

func main() {
	var ikbal costumer
	ikbal.Name = "ikbal"
	ikbal.Address = "indonesia"
	ikbal.Age = 20
	fmt.Println(ikbal)

	//menggunakan struct literal
	ikbal := costumer {
		Name: "ikbal",
		Address: "indonesia",
		Age: 21,
	}
	fmt.Println(ikbal)
}