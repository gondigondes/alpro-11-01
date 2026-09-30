package main

import "fmt"

func main() {
	var p, l int

	// Membaca masukan p dan l
	fmt.Scan(&p, &l)

	// Menghitung luas dan keliling
	luas := p * l
	keliling := 2 * (p + l)

	// Mencetak luas dan keliling dipisahkan spasi
	fmt.Println(luas, keliling)
}

/*
==================================================
LEMBAR JAWABAN

Persoalan Komputasi
• Data yang sudah diketahui: Tidak ada
• Data yang dibaca: Dua bilangan bulat p dan l
• Proses: Membaca p dan l, menghitung luas = p * l dan keliling = 2 * (p + l)
• Informasi yang dicetak: Dua bilangan yang menyatakan luas dan keliling persegi panjang

Pseudocode
program Luas_Keliling
kamus
    p, l, luas, keliling : integer
algoritma
    read(p, l)
    luas <- p * l
    keliling <- 2 * (p + l)
    write(luas, keliling)
endprogram
==================================================
*/