package main

import "fmt"

func getFullName() (string, string) {
	return "ikbal", "bahrudin"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}