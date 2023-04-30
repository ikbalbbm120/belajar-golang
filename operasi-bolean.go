package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusnilaiAkhir = nilaiAkhir > 90
	var lulusabsensi = absensi > 80

	var lulus = lulusabsensi && lulusnilaiAkhir
	fmt.Println(lulus)
}