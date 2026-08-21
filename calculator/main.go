package main

import "fmt"

func tambah(a float64, b float64) float64 {
    return a+b
}

func kurang(a float64, b float64) float64 {
    return a-b
}

func kali(a float64, b float64) float64 {
    return a * b
}

func bagi(a float64, b float64) float64 {
    return a / b
}

func main() {
    var angka1 float64
    var angka2 float64
    var operator string

    fmt.Println("=== CALCULATOR SEDERHANA ===")

    fmt.Println("Masukan Angka Pertama: ")
    fmt.Scan(&angka1)

    fmt.Println("masukan operator (+, -, *, /): ")
    fmt.Scan(&operator)

    fmt.Println("Masukan Angka kedua: ")
    fmt.Scan(&angka2)

    switch operator {
    case "+" :
        fmt.Println("Hasil: ", tambah(angka1, angka2))

    case "-" :
        fmt.Println("Hasil: ", kurang(angka1, angka2))

    case "*" :
        fmt.Println("Hasil:" , kali(angka1, angka2))

    case "/" :
        if angka2 == 0 {
            fmt.Println("Error: tidak bisa membagi dengan 0")
            return

            fmt.Println("Hasil: ", bagi(angka1,angka2))
        }

        default: 
        
        fmt.Println("Operator Tidak Valid")

    }

}