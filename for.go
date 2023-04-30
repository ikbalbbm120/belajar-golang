package main

import "fmt"

func main() {
	counter := 1

	for counter > 10 {

		counter++
	}
	//kode for menggunakan statement
	for counter := 1; counter < 10; counter++ {

	}
	//menggunakan for range
	names := []string{"jmbd", "ikbal", "ngewe"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}