package main

import "fmt"

func main() {

	var name1 = "ikbal"
	var name2 = "ikbal"

	var result bool = name1 == name2
	fmt.Println(result)

	//ini menggunakan angka
	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}