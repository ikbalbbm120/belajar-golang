package main

import "fmt"

func main() {
	data := list.new()

	data.PushBack("ikbal")
	data.PushBack("ikbal")
	data.PushBack("ikbal")
	data.PushFront("ikbal")

	fmt.Println(data.Front().prev())
	fmt.Println(data.Back().Next())

	for element := data.Back(); element != nil; element = element.Next() {
		fmt.Prinntln(element.value)
	}
}