package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func saveContacts(contacts []Contact) {
	data, err := json.MarshalIndent(contacts, "", "  ")

	if err != nil {
		fmt.Println("Gagal mengubah data menjadi JSON:", err)
		return
	}

	err = os.WriteFile("contacts.json", data, 0644)

	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
	}
}

func loadContacts() []Contact {
	data, err := os.ReadFile("contacts.json")

	if err != nil {
		return []Contact{}
	}

	var contacts []Contact

	err = json.Unmarshal(data, &contacts)

	if err != nil {
		fmt.Println("Gagal membaca data JSON:", err)
		return []Contact{}
	}

	return contacts
}

func showMenu() {
	fmt.Println("\n=== CONTACT MANAGER JSON ===")
	fmt.Println("1. Tambah Kontak")
	fmt.Println("2. Lihat Kontak")
	fmt.Println("3. Hapus Kontak")
	fmt.Println("4. Keluar")
}

func addContact(contacts []Contact) []Contact {
	var contact Contact

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&contact.Name)

	fmt.Print("Masukkan nomor telepon: ")
	fmt.Scanln(&contact.Phone)

	contacts = append(contacts, contact)
	saveContacts(contacts)

	fmt.Println("Kontak berhasil ditambahkan dan disimpan!")
	return contacts
}

func viewContacts(contacts []Contact) {
	fmt.Println("\n=== DAFTAR KONTAK ===")

	if len(contacts) == 0 {
		fmt.Println("Belum ada kontak.")
		return
	}

	for i, contact := range contacts {
		fmt.Printf("%d. %s - %s\n", i+1, contact.Name, contact.Phone)
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

	if contactNumber < 1 || contactNumber > len(contacts) {
		fmt.Println("Nomor kontak tidak valid.")
		return contacts
	}

	contacts = append(contacts[:contactNumber-1], contacts[contactNumber:]...)
	saveContacts(contacts)

	fmt.Println("Kontak berhasil dihapus dan data diperbarui!")
	return contacts
}

func main() {
	contacts := loadContacts()

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
			contacts = deleteContact(contacts)
		} else if choice == 4 {
			fmt.Println("Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}