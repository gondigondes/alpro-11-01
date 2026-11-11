package main

import "fmt"

func main() {
	var tipe string
	var motor, mobil, truk int
	var selesai bool = false

	for !selesai {
		fmt.Scan(&tipe)

		if tipe == "M" {
			motor++
		} else if tipe == "K" {
			mobil++
		} else if tipe == "T" {
			truk++
		} else {
			selesai = true
		}
	}

	fmt.Println("Tipe M =", motor)
	fmt.Println("Tipe K =", mobil)
	fmt.Println("Tipe T =", truk)
}
