package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 	int
}

func (customer Customer)sayhello(name string ) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main() {
	var ikbal Customer
	ikbal.Name = "ikbal"
	ikbal.Address = "indonesia"
	ikbal.Age = 20


}