package main

import "fmt"

func main() {
	var x float64

	// Membaca masukan x
	fmt.Scan(&x)

	// Menghitung hasil persamaan f(x)
	hasil := (x*x + 2*x + 1) / (x - 3)

	// Mencetak hasil
	fmt.Println(hasil)
}

/*
==================================================
LEMBAR JAWABAN

Persoalan Komputasi
• Data yang sudah diketahui: Tidak ada
• Data yang dibaca: Sebuah bilangan berkoma x
• Proses: Membaca x, lalu menghitung f(x) = (x² + 2x + 1) / (x - 3)
• Informasi yang dicetak: Sebuah bilangan yang menyatakan nilai dari f(x)

Pseudocode
program Matematika_F
kamus
    x, hasil : real
algoritma
    read(x)
    hasil <- (x * x + 2 * x + 1) / (x - 3)
    write(hasil)
endprogram
==================================================
*/