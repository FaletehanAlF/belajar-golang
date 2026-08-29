package main

import (
	"bufio"
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("\n=== SIMPLE NOTES APP ===")
	fmt.Println("1. Tambah Catatan")
	fmt.Println("2. Lihat Semua Catatan")
	fmt.Println("3. Keluar")
}

func addNote(reader *bufio.Reader) {
	fmt.Print("Masukkan catatan: ")
	note, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input.")
		return
	}

	file, err := os.OpenFile(
		"notes.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Println("Terjadi kesalahan saat membuka file.")
		return
	}

	defer file.Close()

	_, err = file.WriteString(note)

	if err != nil {
		fmt.Println("Terjadi kesalahan saat menyimpan catatan.")
		return
	}

	fmt.Println("Catatan berhasil disimpan!")
}

func viewNotes() {
	data, err := os.ReadFile("notes.txt")

	if err != nil {
		fmt.Println("Belum ada catatan.")
		return
	}

	fmt.Println("\n=== SEMUA CATATAN ===")
	fmt.Println(string(data))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var choice int

		showMenu()

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			reader.ReadString('\n')
			addNote(reader)
		} else if choice == 2 {
			viewNotes()
		} else if choice == 3 {
			fmt.Println("Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}