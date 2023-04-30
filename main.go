package main

import "fmt"

func main() {
	person := map[string]string {
		"name": "ikbal",
		"address": "jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//function map
	book := make(map[string]string)
	book["tittle"] = "buku go-lang"
	book["author"] = "ikbal"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)
}