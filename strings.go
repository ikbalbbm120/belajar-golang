package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("ikbal", "ikbal"))
	fmt.Println(strings.Split("ikbal", " "))

	fmt.Println(strings.ToLower("ikbal"))
	fmt.Println(strings.ToUpper("ikbal"))
	fmt.Println(strings.ToTitle("ikbal"))

	fmt.Println(strings.Trim("  ikbal  ", " "))
	fmt.Println(strings.ReplaceAll("ikbal ikbal ikbal", "ikbal", "jmbd"))
}