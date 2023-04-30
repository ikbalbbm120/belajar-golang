package main

import (
	"Errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, Errors.new("pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}


func main() {
	hasil, err := pembagi(100, 10)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}