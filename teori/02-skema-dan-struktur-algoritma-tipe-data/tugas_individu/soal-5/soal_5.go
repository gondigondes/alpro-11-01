package main

import "fmt"

func main() {
	var harga, persen float64

	// Membaca masukan harga dan persen diskon
	fmt.Scan(&harga, &persen)

	// Menghitung besarnya potongan harga dan harga akhir
	potongan := harga * persen / 100.0
	hargaAkhir := harga - potongan

	// Mencetak harga akhir
	fmt.Println(hargaAkhir)
}

/*
==================================================
LEMBAR JAWABAN

Persoalan Komputasi
• Data yang sudah diketahui: Tidak ada
• Data yang dibaca: Dua bilangan bulat harga dan persen
• Proses: Membaca harga dan persen, menghitung potongan = harga * persen / 100, lalu harga_akhir = harga - potongan
• Informasi yang dicetak: Sebuah bilangan yang menyatakan harga akhir setelah diskon

Pseudocode
program Diskon_Belanja
kamus
    harga, persen : integer
    potongan, harga_akhir : real
algoritma
    read(harga, persen)
    potongan <- harga * persen / 100.0
    harga_akhir <- harga - potongan
    write(harga_akhir)
endprogram
==================================================
*/