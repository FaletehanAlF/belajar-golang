package main

import "fmt"

func main() {
	var number int

	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scanln(&number)

	fmt.Println("Angka yang kamu masukkan:", number)
}