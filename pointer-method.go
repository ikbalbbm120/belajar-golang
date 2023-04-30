package main

import "fmt"

type man struct {
	new string
}

func (man *Man) Married() {
	man.Name = "mr. " + man.Name
}

func main() {
	ikbal := Man("ikbal")
	ikbal.married()
}