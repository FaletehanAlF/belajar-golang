package main

import "fmt"

func getGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	}

	return "E"
}

func getDescription(grade string) string {
	if grade == "A" {
		return "Sangat Baik"
	} else if grade == "B" {
		return "Baik"
	} else if grade == "C" {
		return "Cukup"
	} else if grade == "D" {
		return "Kurang"
	}

	return "Sangat Kurang"
}

func getValidScore() int {
	for {
		var score int

		fmt.Print("\nMasukkan nilai (0-100): ")
		_, err := fmt.Scanln(&score)

		if err != nil {
			fmt.Println("Input tidak valid. Masukkan angka.")
			continue
		}

		if score < 0 || score > 100 {
			fmt.Println("Nilai harus antara 0 sampai 100.")
			continue
		}

		return score
	}
}

func main() {
	fmt.Println("=== GRADE CALCULATOR ===")

	for {
		score := getValidScore()

		grade := getGrade(score)
		description := getDescription(grade)

		fmt.Println("\n=== HASIL ===")
		fmt.Println("Nilai:", score)
		fmt.Println("Grade:", grade)
		fmt.Println("Keterangan:", description)

		var answer string
		fmt.Print("\nHitung nilai lagi? (y/n): ")
		fmt.Scanln(&answer)

		if answer != "y" && answer != "Y" {
			fmt.Println("\nTerima kasih telah menggunakan Grade Calculator!")
			break
		}
	}
}