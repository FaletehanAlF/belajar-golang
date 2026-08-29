package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
}

func showMenu() {
	fmt.Println("\n=== CONTACT MANAGER ===")
	fmt.Println("1. Tambah Kontak")
	fmt.Println("2. Lihat Semua Kontak")
	fmt.Println("3. Cari Kontak")
	fmt.Println("4. Hapus Kontak")
	fmt.Println("5. Keluar")
}

func addContact(contacts []Contact) []Contact {
	var contact Contact

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&contact.Name)

	fmt.Print("Masukkan nomor telepon: ")
	fmt.Scanln(&contact.Phone)

	contacts = append(contacts, contact)
	fmt.Println("Kontak berhasil ditambahkan!")

	return contacts
}

func viewContacts(contacts []Contact) {
	fmt.Println("\n=== DAFTAR KONTAK ===")

	if len(contacts) == 0 {
		fmt.Println("Belum ada kontak.")
	} else {
		for i, contact := range contacts {
			fmt.Printf("%d. %s - %s\n", i+1, contact.Name, contact.Phone)
		}
	}
}

func searchContact(contacts []Contact) {
	var keyword string
	found := false

	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scanln(&keyword)

	fmt.Println("\n=== HASIL PENCARIAN ===")

	for _, contact := range contacts {
		if strings.EqualFold(contact.Name, keyword) {
			fmt.Printf("%s - %s\n", contact.Name, contact.Phone)
			found = true
		}
	}

	if !found {
		fmt.Println("Kontak tidak ditemukan.")
	}
}

func deleteContact(contacts []Contact) []Contact {
	if len(contacts) == 0 {
		fmt.Println("Tidak ada kontak yang bisa dihapus.")
		return contacts
	}

	viewContacts(contacts)

	var contactNumber int
	fmt.Print("Masukkan nomor kontak yang ingin dihapus: ")
	fmt.Scanln(&contactNumber)

	if contactNumber >= 1 && contactNumber <= len(contacts) {
		contacts = append(contacts[:contactNumber-1], contacts[contactNumber:]...)
		fmt.Println("Kontak berhasil dihapus!")
	} else {
		fmt.Println("Nomor kontak tidak valid.")
	}

	return contacts
}

func main() {
	var contacts []Contact

	for {
		var choice int

		showMenu()

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			contacts = addContact(contacts)
		} else if choice == 2 {
			viewContacts(contacts)
		} else if choice == 3 {
			searchContact(contacts)
		} else if choice == 4 {
			contacts = deleteContact(contacts)
		} else if choice == 5 {
			fmt.Println("Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}