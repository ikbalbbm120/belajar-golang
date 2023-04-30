package main

import "fmt"

func sumAll(numbers ... int) int {
	total := 0
	
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll()
	fmt.Println(total)
}