package main

import "fmt"

func main() {
	var r float64

	// Menerima masukan jari-jari dari pengguna
	fmt.Scan(&r)

	// Assignment variabel pi
	pi := 3.14

	// Menghitung luas lingkaran
	luas := pi * r * r

	// Mencetak hasil keluaran menggunakan fmt.Println
	fmt.Println(luas)
}