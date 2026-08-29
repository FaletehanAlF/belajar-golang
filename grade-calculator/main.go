package main

import "fmt"

func main() {
	var grade float64

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&grade)

	if grade < 0 || grade > 100 {
		fmt.Println("Nilai tidak valid. Masukkan nilai antara 0 sampai 100.")
	} else if grade >= 90 {
		fmt.Println("Grade: A")
	} else if grade >= 80 {
		fmt.Println("Grade: B")
	} else if grade >= 70 {
		fmt.Println("Grade: C")
	} else if grade >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: E")
	}
}