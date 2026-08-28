package main

import "fmt"

func main() {
	var grade float64

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&grade)

	fmt.Println("Nilai yang kamu masukkan:", grade)
}