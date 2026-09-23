package main

import "fmt"

func main() {
	var nama string
	var mat, ing int

	// membaca input
	fmt.Scan(&nama)
	fmt.Scan(&mat)
	fmt.Scan(&ing)

	total := mat + ing
	rataRata := total / 2

	fmt.Println("Nama:", nama)
	fmt.Println("Total Nilai:", total)
	fmt.Println("Rata-rata Nilai:", rataRata)
}
