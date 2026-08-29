package main

import "fmt"

type Expense struct {
	Name   string
	Amount int
}

func main() {
	var expenses []Expense

	for {
		var choice int

		fmt.Println("\n=== SIMPLE EXPENSE TRACKER ===")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Lihat Pengeluaran")
		fmt.Println("3. Lihat Total Pengeluaran")
		fmt.Println("4. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			var expense Expense

			fmt.Print("Nama pengeluaran: ")
			fmt.Scanln(&expense.Name)

			fmt.Print("Jumlah pengeluaran: ")
			fmt.Scanln(&expense.Amount)

			expenses = append(expenses, expense)
			fmt.Println("Pengeluaran berhasil ditambahkan!")

		} else if choice == 2 {
			fmt.Println("\n=== DAFTAR PENGELUARAN ===")

			if len(expenses) == 0 {
				fmt.Println("Belum ada data pengeluaran.")
			} else {
				for i, expense := range expenses {
					fmt.Printf("%d. %s - Rp%d\n", i+1, expense.Name, expense.Amount)
				}
			}

		} else if choice == 3 {
			var total int

			for _, expense := range expenses {
				total += expense.Amount
			}

			fmt.Println("\n=== TOTAL PENGELUARAN ===")
			fmt.Printf("Total: Rp%d\n", total)

		} else if choice == 4 {
			fmt.Println("Terima kasih!")
			break

		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}