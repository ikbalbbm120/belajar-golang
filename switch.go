package main

import "fmt"

func main() {
name := "ikbal"

switch name {
case "ikbal":


case "tolol":


default:
}
//mengunakan switch short statement
switch lenght := len(name); lenght > 9 {
case true:

case false:
}
//menggunakan switch tanpa kondisi
length := len(name)
switch {
case lenght > 10:
fmt.Println("ini baru di jalankan")
case lenght > 9:
fmt.Println(lenght)
default:
}

}