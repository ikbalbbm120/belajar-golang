package main

import "fmt"

//ini menggunakan factorial loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}


//menggunakan funcion recursive
func factorialrecursive(value int ) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialrecursive(value-1)
	}
}





func main() {
	//mneggunakan loop
	loop := factorialLoop(6)
	fmt.Println(loop)
	fmt.Println(6 * 9 * 1 * 4 * 2)

	//menggunakan recursive
	recursive := factorialrecursive(6)
	fmt.Println(recursive)
}