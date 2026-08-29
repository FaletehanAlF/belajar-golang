package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Class string
}

func main() {
	var students []Student

	for {
		var choice int

		fmt.Println("\n=== STUDENT DATA MANAGER ===")
		fmt.Println("1. Tambah Siswa")
		fmt.Println("2. Lihat Data Siswa")
		fmt.Println("3. Hapus Data Siswa")
		fmt.Println("4. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			var student Student

			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&student.Name)

			fmt.Print("Masukkan umur: ")
			fmt.Scanln(&student.Age)

			fmt.Print("Masukkan kelas: ")
			fmt.Scanln(&student.Class)

			students = append(students, student)
			fmt.Println("Data siswa berhasil ditambahkan!")

		} else if choice == 2 {
			fmt.Println("\n=== DAFTAR SISWA ===")

			if len(students) == 0 {
				fmt.Println("Belum ada data siswa.")
			} else {
				for i, student := range students {
					fmt.Printf("%d. %s | Umur: %d | Kelas: %s\n",
						i+1,
						student.Name,
						student.Age,
						student.Class,
					)
				}
			}

		} else if choice == 3 {
			if len(students) == 0 {
				fmt.Println("Tidak ada data siswa yang bisa dihapus.")
			} else {
				fmt.Println("\n=== HAPUS DATA SISWA ===")

				for i, student := range students {
					fmt.Printf("%d. %s | Umur: %d | Kelas: %s\n",
						i+1,
						student.Name,
						student.Age,
						student.Class,
					)
				}

				var studentNumber int
				fmt.Print("Masukkan nomor siswa yang ingin dihapus: ")
				fmt.Scanln(&studentNumber)

				if studentNumber >= 1 && studentNumber <= len(students) {
					students = append(students[:studentNumber-1], students[studentNumber:]...)
					fmt.Println("Data siswa berhasil dihapus!")
				} else {
					fmt.Println("Nomor siswa tidak valid.")
				}
			}

		} else if choice == 4 {
			fmt.Println("Terima kasih!")
			break

		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}