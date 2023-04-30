package main

import "fmt"

type address struct {
	city, province, country string
}

//pointer function
func changecountrytoindonesia(address *address) {
	address.country = "indonesia"
}

func main() {
	address1 address := address{"jakarta", "dki jakarta", "indonesia"}
	address2 *address = &address1

	address2.city = "cilacap"

	address2 = &address{"bandung", "jawa barat", "indonesia"}
	//menggunakan bintang
	*address2 = address{"bandung", "jawa barat", "indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	//menggunakan func new
	alamat1 := new(address)
	alamat2 := alamat1

	alamat2.country = "indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

	var alamat = address {
		city: "kontolodon",
		province: "ngentot",
		country: ,
	}
	changechangecountrytoindonesia(&alamat)
	fmt.Println(alamat)
}