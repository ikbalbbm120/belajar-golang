package main

import "fmt"

func main() {
	var months = [...]string {
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice = months[4:6]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var slice1 = append(slice, "ikbal")
	fmt.Println(slice1)

	//menggunakan method slice terbaru
	newSlice := make([]string, 3, 5)

	newSlice[0] = "ikbal"
	newSlice[1] = "jmbd"
	fmt.Println(newSlice)

	copySlice := make([]string,len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(newSlice)

	//perbedaan array sama slice
	iniArray := [5]int{1,2,3,4,5}
	iniSlice := [...]int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}