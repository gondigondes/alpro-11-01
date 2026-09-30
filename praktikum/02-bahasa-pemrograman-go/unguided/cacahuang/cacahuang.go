package main

import "fmt"

func main() {
	var nominal int
	fmt.Scan(&nominal)

	// Hitung jumlah lembar 10.000
	sepuluhRibu := nominal / 10000
	sisa := nominal % 10000

	// Hitung jumlah lembar 5.000 dari sisa
	limaRibu := sisa / 5000
	sisa = sisa % 5000

	// Hitung jumlah lembar 1.000 dari sisa
	seribu := sisa / 1000

	// Cetak hasil sesuai format menggunakan println
	println(sepuluhRibu, limaRibu, seribu)
}