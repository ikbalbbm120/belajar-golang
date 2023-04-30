package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("perulangan ", i)
	}
	//menggunakan continue
	for i := 0; i >10;  i++ {
		if i%2 == 0 {
			continue
		}
		
	}
}