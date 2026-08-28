package main

import "fmt"

func main() {
	var username string
	var password string

	fmt.Println("=== Simple Login ===")

	fmt.Print("Masukkan username: ")
	fmt.Scanln(&username)

	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	fmt.Println("Username yang dimasukkan:", username)
	fmt.Println("Password yang dimasukkan:", password)
}