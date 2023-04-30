package main

import "fmt"

func main() {
	name := "ikbal"
	counter := 0

	increment := func() {
		name = "jmbd"
		fmt.Println("incremenet")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}