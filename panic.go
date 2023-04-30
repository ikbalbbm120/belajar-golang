package main

import "fmt"

func endapp() {
	fmt.Println("kontol")
	message := recover()
	if message == nill {
	fmt.Println("tolol", message)
	}
	fmt.Println("kontol")
}

func runapp(error bool) {
	defer endapp()
	if error {
		panic("pler nya pecah")
	}
	fmt.Println("jancok")
}

func main() {
	runapp(false)
	fmt.Println("ikbal")
}
