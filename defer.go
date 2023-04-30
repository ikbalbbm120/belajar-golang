package main

import "fmt"

func logging() {
	fmt.Println("jancok")
}

func runaplication(value int) {
	defer logging()
	fmt.Println("jancok")
	result := 20 / value
	fmt.Println("result", result)
}

func main() {
	runaplication(0)
}