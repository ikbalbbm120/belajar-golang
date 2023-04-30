package main

import "fmt"


func getFUllName() (firstName, middleName, lastName string) {
	firstName = "ikbal"
	middleName = "jmbd"
	lastName = "kontol"
	return
}

func main() {
	firstName, middleName, lastName := getFUllName()
	fmt.Println(firstName,middleName,lastName)
}