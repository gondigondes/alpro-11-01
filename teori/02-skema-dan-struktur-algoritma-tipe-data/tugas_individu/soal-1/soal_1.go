package main

import "fmt"

func main() {
	var x, y float64

	// Membaca masukan x dan y
	fmt.Scan(&x, &y)

	// Menghitung hasil persamaan f(x, y)
	hasil := 5*(x*x) - 2*x*y + (y*y*y)/(x+1)

	// Mencetak hasil
	fmt.Println(hasil)
}

/*
==================================================
LEMBAR JAWABAN

Persoalan Komputasi
• Data yang sudah diketahui: Tidak ada
• Data yang dibaca: Dua bilangan bulat x dan y
• Proses: Membaca x dan y, lalu menghitung f(x, y) = 5x² - 2xy + y³ / (x + 1)
• Informasi yang dicetak: Sebuah bilangan yang menyatakan nilai dari f(x, y)

Pseudocode
program Matematika_E
kamus
    x, y : integer
    hasil : real
algoritma
    read(x, y)
    hasil <- 5 * (x * x) - 2 * x * y + (y * y * y) / (x + 1)
    write(hasil)
endprogram
==================================================
*/