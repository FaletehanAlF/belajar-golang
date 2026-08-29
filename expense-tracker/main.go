package main

import "fmt"

type Expense struct {
	Name   string
	Amount int
}

func showMenu() {
	fmt.Println("\n=== SIMPLE EXPENSE TRACKER ===")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Lihat Pengeluaran")
	fmt.Println("3. Lihat Total Pengeluaran")
	fmt.Println("4. Keluar")
}

func addExpense(expenses []Expense) []Expense {
	var expense Expense

	fmt.Print("Nama pengeluaran: ")
	fmt.Scanln(&expense.Name)

	fmt.Print("Jumlah pengeluaran: ")
	fmt.Scanln(&expense.Amount)

	expenses = append(expenses, expense)
	fmt.Println("Pengeluaran berhasil ditambahkan!")

	return expenses
}

func viewExpenses(expenses []Expense) {
	fmt.Println("\n=== DAFTAR PENGELUARAN ===")

	if len(expenses) == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		for i, expense := range expenses {
			fmt.Printf("%d. %s - Rp%d\n", i+1, expense.Name, expense.Amount)
		}
	}
}

func showTotal(expenses []Expense) {
	var total int

	for _, expense := range expenses {
		total += expense.Amount
	}

	fmt.Println("\n=== TOTAL PENGELUARAN ===")
	fmt.Printf("Total: Rp%d\n", total)
}

func main() {
	var expenses []Expense

	for {
		var choice int

		showMenu()

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			expenses = addExpense(expenses)
		} else if choice == 2 {
			viewExpenses(expenses)
		} else if choice == 3 {
			showTotal(expenses)
		} else if choice == 4 {
			fmt.Println("Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}