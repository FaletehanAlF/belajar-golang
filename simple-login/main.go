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

	if username == "Falen" && password == "123" {
		fmt.Println("Login berhasil!")
	} else {
		fmt.Println("Username atau password salah!")
	}
}