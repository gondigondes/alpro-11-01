package main

import "fmt"

func main() {
	var c float64

	// Membaca masukan c (Celsius)
	fmt.Scan(&c)

	// Menghitung konversi ke Fahrenheit
	f := c*9.0/5.0 + 32.0

	// Mencetak hasil Fahrenheit
	fmt.Println(f)
}

/*
==================================================
LEMBAR JAWABAN

Persoalan Komputasi
• Data yang sudah diketahui: Tidak ada
• Data yang dibaca: Sebuah bilangan c
• Proses: Membaca c, lalu menghitung f = c * 9 / 5 + 32
• Informasi yang dicetak: Sebuah bilangan yang menyatakan suhu dalam derajat Fahrenheit

Pseudocode
program Konversi_Suhu
kamus
    c, f : real
algoritma
    read(c)
    f <- c * 9.0 / 5.0 + 32.0
    write(f)
endprogram
==================================================
*/