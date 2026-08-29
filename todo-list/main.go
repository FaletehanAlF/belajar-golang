package main

import "fmt"

func showMenu() {
	fmt.Println("\n=== TO-DO LIST ===")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Tugas")
	fmt.Println("3. Hapus Tugas")
	fmt.Println("4. Keluar")
}

func addTask(tasks []string) []string {
	var task string

	fmt.Print("Masukkan tugas: ")
	fmt.Scanln(&task)

	tasks = append(tasks, task)
	fmt.Println("Tugas berhasil ditambahkan!")

	return tasks
}

func viewTasks(tasks []string) {
	fmt.Println("\n=== DAFTAR TUGAS ===")

	if len(tasks) == 0 {
		fmt.Println("Belum ada tugas.")
	} else {
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
}

func deleteTask(tasks []string) []string {
	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas yang bisa dihapus.")
		return tasks
	}

	viewTasks(tasks)

	var taskNumber int
	fmt.Print("Masukkan nomor tugas yang ingin dihapus: ")
	fmt.Scanln(&taskNumber)

	if taskNumber >= 1 && taskNumber <= len(tasks) {
		tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
		fmt.Println("Tugas berhasil dihapus!")
	} else {
		fmt.Println("Nomor tugas tidak valid.")
	}

	return tasks
}

func main() {
	var tasks []string

	for {
		var choice int

		showMenu()

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			tasks = addTask(tasks)
		} else if choice == 2 {
			viewTasks(tasks)
		} else if choice == 3 {
			tasks = deleteTask(tasks)
		} else if choice == 4 {
			fmt.Println("Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}