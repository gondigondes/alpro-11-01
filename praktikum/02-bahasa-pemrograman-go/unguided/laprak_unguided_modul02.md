# <h1 align="center">Laporan Praktikum Unguided - Bahasa Pemrograman Go</h1>
<p align="center">Akmal Ghani Laoda - 109092600005</p>



## Unguided

### 1. `cacaguang.go`

```go
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
```

##### Output

![Screenshot Output Unguided](unguided/cacaguang/output.png)

#### Deskripsi

Program `cacaguang.go` digunakan untuk menghitung jumlah lembar uang pecahan Rp10.000, Rp5.000, dan Rp1.000 berdasarkan nominal yang dimasukkan oleh pengguna. Program menerima input nominal menggunakan `fmt.Scan()`. Selanjutnya, jumlah lembar Rp10.000 dihitung menggunakan operasi pembagian, sedangkan operasi modulus digunakan untuk mendapatkan sisa nominal. Sisa tersebut kemudian digunakan untuk menghitung jumlah lembar Rp5.000 dan Rp1.000. Hasil akhir berupa jumlah masing-masing pecahan ditampilkan menggunakan `println()`.

### 2. `kalkulator.go`

```go
package main

import "fmt"

func main() {
	var a, b int

	// Membaca input
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Menampilkan output
	fmt.Print("Hasil Penjumlahan: ")
	fmt.Println(a + b)
	fmt.Print("Hasil Pengurangan: ")
	fmt.Println(a - b)
	fmt.Print("Hasil Perkalian: ")
	fmt.Println(a * b)
	fmt.Print("Hasil Pembagian: ")
	fmt.Println(a / b)
	fmt.Print("Hasil Sisa Hasil Bagi: ")
	fmt.Println(a % b)
}
```

##### Output

![Screenshot Output Unguided](unguided/kalkulator/output.png)

#### Deskripsi

Program `kalkulator.go` digunakan untuk melakukan operasi aritmatika terhadap dua bilangan yang dimasukkan oleh pengguna. Program menerima nilai `a` dan `b`, kemudian melakukan operasi penjumlahan, pengurangan, perkalian, pembagian, dan sisa hasil bagi. Setiap hasil ditampilkan menggunakan `fmt.Print()` dan `fmt.Println()` dengan keterangan masing-masing operasi. Program ini menerapkan penggunaan variabel bertipe `int`, input menggunakan `fmt.Scan()`, serta operator aritmatika dasar dalam bahasa Go.

## Kesimpulan

Berdasarkan dua program unguided yang telah dibuat, dapat dipahami penerapan dasar bahasa Go dalam menyelesaikan permasalahan sederhana menggunakan operasi aritmatika. Program `cacaguang.go` menerapkan operasi pembagian dan modulus untuk menentukan jumlah pecahan uang, sedangkan program `kalkulator.go` menerapkan berbagai operator aritmatika untuk mengolah dua nilai. Dari praktikum ini dapat dipahami penggunaan variabel, input, proses perhitungan, dan output dalam program Go secara lebih mandiri.

## Referensi

1. Donovan, A. A. A., & Kernighan, B. W. (2015). *The Go Programming Language*. Boston: Addison-Wesley Professional. Diakses melalui https://www.gopl.io/
2. The Go Authors. (2026). *The Go Programming Language Specification*. Diakses melalui https://go.dev/ref/spec
3. The Go Authors. (2026). *A Tour of Go*. Diakses melalui https://go.dev/tour/
