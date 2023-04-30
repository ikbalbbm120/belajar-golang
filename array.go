package main

import "fmt"

func main() {
	var names [3] string
	
	names[0] = "ikbal"
	names[1] = "bahrudin"
	names[2] = "jmbd"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//mengunakan array langsung
	var values = [4]int{
		80,
		70,
		90,
	}

	fmt.Println(values)

}