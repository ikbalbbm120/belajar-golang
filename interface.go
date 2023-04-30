package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type animal struct {
	name string
}

func (animal animal) GetName() string {
	return animal.name
}

func main() {
	var ikbal Person
	ikbal.Name = "ikbal"

	sayHello(ikbal)

	cat := animal { 
		name: "jmbd",
	}
	sayhHello(cat)
}