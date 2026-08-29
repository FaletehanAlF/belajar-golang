package main

import "fmt"

func main() {
	var choice int
	var task string

	var tasks []string

	fmt.Println("=== TO-DO LIST ===")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Tugas")
	fmt.Println("3. Hapus Tugas")
	fmt.Println("4. Keluar")

	fmt.Print("Pilih menu: ")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Print("Masukkan tugas: ")
		fmt.Scanln(&task)

		tasks = append(tasks, task)

		fmt.Println("Tugas berhasil ditambahkan:", task)
	}
}