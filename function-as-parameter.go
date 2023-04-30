package main

import "fmt"

type Filter func(string) string

func sayHelloWithFillter(name string, Filter Filter) {
	nameFiltered := Filter(name)
	fmt.Println("ikbal", nameFiltered)
}

func spamFillter(name string) string {
	if name == "babi" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFillter("ikbal", spamFillter)
	sayHelloWithFillter("kontol", spamFillter)
}