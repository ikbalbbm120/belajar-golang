package main

import "fmt"

func random() interface{} {
	return "ikbal"
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	}
}