package main

import (
	"golang1/database"
	"fmt"
)

func main() {
	result := database.GetDatabases()
	fmt.Println(result)
}