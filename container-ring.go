package main

import (
	"fmt"
	"strconv"
)

func main() {

	var data *ring.Ring = ring.
	data := ring.New(5)
	for i := 0; i < data.len(); i++ {
		data.value = "value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
	fmt.Println(value)
	})
}