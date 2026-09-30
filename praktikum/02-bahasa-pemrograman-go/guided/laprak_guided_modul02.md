# <h1 align="center">Laporan Praktikum Modul 02 - Bahasa Pemrograman Go</h1>
<p align="center">Akmal Ghani Laoda - 109092600005</p>

## Dasar Teori

### A. Bahasa Pemrograman Go

Go merupakan bahasa pemrograman yang dikembangkan oleh Google dan dirancang untuk menghasilkan program yang sederhana, efisien, serta mudah dipelihara. Go memiliki sintaks yang relatif sederhana dan mendukung pemrograman terstruktur maupun konkuren. Program Go umumnya terdiri dari package, deklarasi fungsi, variabel, tipe data, serta perintah yang dijalankan secara berurutan.

Dalam praktikum ini, bahasa Go digunakan untuk mempelajari dasar pembuatan program, menerima input dari pengguna, melakukan operasi aritmatika, menggunakan variabel, serta menampilkan hasil ke layar.

### B. Package dan Struktur Program di Go

#### 1. Package `main` dan `func main()`

Setiap program Go yang dapat dijalankan sebagai program utama menggunakan `package main`. Fungsi `main()` merupakan titik awal eksekusi program. Perintah yang berada di dalam fungsi tersebut akan dijalankan ketika program dieksekusi.

Contoh struktur dasar program Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

Package `fmt` digunakan untuk berbagai operasi input dan output, seperti membaca masukan menggunakan `fmt.Scan()` dan menampilkan keluaran menggunakan `fmt.Println()`.

#### 2. Tipe Data dan Deklarasi Variabel

Variabel digunakan untuk menyimpan data yang diperlukan selama program berjalan. Go menyediakan berbagai tipe data, di antaranya `int` untuk bilangan bulat, `float64` untuk bilangan pecahan, dan `string` untuk teks.

Deklarasi variabel dapat dilakukan menggunakan kata kunci `var`, misalnya:

```go
var nama string
var nilai int
var suhu float64
```

Go juga menyediakan deklarasi singkat menggunakan operator `:=`, contohnya:

```go
pi := 3.14
```

#### 3. Input dan Output

Input dari pengguna dapat dibaca menggunakan `fmt.Scan()`. Agar nilai masukan disimpan ke dalam variabel, alamat variabel diberikan menggunakan operator `&`.

Contoh:

```go
var r float64
fmt.Scan(&r)
```

Sementara itu, output dapat ditampilkan menggunakan `fmt.Println()` atau `fmt.Print()`.

#### 4. Operasi Aritmatika

Go mendukung operasi aritmatika seperti penjumlahan (`+`), pengurangan (`-`), perkalian (`*`), dan pembagian (`/`). Operasi tersebut dapat digunakan untuk menghitung nilai berdasarkan data yang dimasukkan pengguna.

Pada tipe data `int`, pembagian menghasilkan bilangan bulat. Sedangkan pada tipe data pecahan seperti `float64`, hasil pembagian dapat berupa bilangan desimal.

## Guided

### 1. `lingkaran.go`

```go
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
```

#### Deskripsi

Program `lingkaran.go` digunakan untuk menghitung luas lingkaran berdasarkan jari-jari yang dimasukkan oleh pengguna. Program terlebih dahulu membuat variabel `r` dengan tipe data `float64`, kemudian membaca nilai jari-jari menggunakan `fmt.Scan()`. Nilai `pi` ditetapkan sebesar `3.14`, lalu luas dihitung menggunakan rumus `pi × r × r`. Hasil perhitungan kemudian ditampilkan menggunakan `fmt.Println()`. Program berhasil menghasilkan nilai luas sesuai dengan jari-jari yang diberikan.

### 2. `skor.go`

```go
package main

import "fmt"

func main() {
    var nama string
    var skor_Emteka, skor_BahasaInggris int

    // Membaca input
    fmt.Println("Masukkan nama:")
    fmt.Scan(&nama)

    fmt.Println("Masukkan skor Emteka:")
    fmt.Scan(&skor_Emteka)

    fmt.Println("Masukkan skor Bahasa Inggris:")
    fmt.Scan(&skor_BahasaInggris)

    // Menghitung total dan rata-rata (Pembagian bilangan bulat)
    total := skor_Emteka + skor_BahasaInggris
    rata_rata := total / 2

    // Menampilkan output
    fmt.Println(nama)
    fmt.Println(total)
    fmt.Println(rata_rata)

    // Menampilkan nama dengan keterangan
    fmt.Println("Nama :", nama)
    fmt.Println("Total Skor :", total)
    fmt.Println("Rata-rata Skor :", rata_rata)
}
```

#### Deskripsi

Program `skor.go` digunakan untuk menerima nama dan dua nilai, yaitu skor Emteka dan skor Bahasa Inggris. Data nama disimpan dalam variabel bertipe `string`, sedangkan kedua skor menggunakan tipe `int`. Setelah input diterima, program menghitung total skor dengan menjumlahkan kedua nilai dan menghitung rata-rata dengan membagi total skor dengan dua. Hasil berupa nama, total skor, dan rata-rata kemudian ditampilkan ke layar. Program juga menunjukkan penggunaan variabel, input, operasi aritmatika, dan output pada bahasa Go.

## Unguided

### 1. `suhu.go`

```go
package main

import "fmt"

func main() {
    var celsius float64

    // Menerima masukan suhu dalam Celsius
    fmt.Scan(&celsius)

    // Menghitung konversi suhu sesuai rumus
    reamur := celsius * 4.0 / 5.0
    fahrenheit := (celsius * 9.0 / 5.0) + 32.0
    kelvin := celsius + 273.15

    // Mencetak tiga nilai keluaran dipisahkan oleh spasi
    fmt.Println(reamur, fahrenheit, kelvin)
}
```

##### Output

![Screenshot Output Unguided](unguided/suhu/output.png)

#### Deskripsi

Program `suhu.go` digunakan untuk mengonversi suhu dari Celsius ke tiga satuan suhu, yaitu Reamur, Fahrenheit, dan Kelvin. Program menerima nilai Celsius dari pengguna menggunakan `fmt.Scan()`. Selanjutnya, program menghitung masing-masing hasil menggunakan rumus konversi yang sesuai. Hasil konversi kemudian ditampilkan secara berurutan menggunakan `fmt.Println()`. Dari program ini dapat dipahami penggunaan tipe data `float64`, operasi aritmatika, input pengguna, dan output dalam bahasa Go.

### 2. `tukar.go`

```go
package main

import "fmt"

func main() {
    var a, b int

    fmt.Print("Masukkan nilai a:")
    fmt.Scan(&a)

    fmt.Print("Masukkan nilai b:")
    fmt.Scan(&b)

    a, b = b, a

    fmt.Println(a)
    fmt.Println(b)
}
```

##### Output

![Screenshot Output Unguided](unguided/tukar/output.png)

#### Deskripsi

Program `tukar.go` digunakan untuk menukar nilai dari dua variabel, yaitu `a` dan `b`. Program terlebih dahulu meminta pengguna memasukkan kedua nilai tersebut. Setelah itu, pertukaran dilakukan menggunakan sintaks `a, b = b, a` tanpa membutuhkan variabel sementara. Hasil pertukaran kemudian ditampilkan menggunakan `fmt.Println()`. Program ini menunjukkan salah satu fitur bahasa Go yang memungkinkan beberapa variabel menerima nilai secara bersamaan.

## Kesimpulan

Berdasarkan praktikum yang telah dilakukan, bahasa pemrograman Go memiliki struktur program yang sederhana dan mudah dipahami. Melalui beberapa program yang dibuat, dapat dipahami penggunaan `package main`, fungsi `main()`, deklarasi variabel, tipe data, input menggunakan `fmt.Scan()`, serta output menggunakan `fmt.Println()` dan `fmt.Print()`. Praktikum juga memberikan pemahaman mengenai operasi aritmatika, perhitungan luas, konversi suhu, pengolahan nilai, serta pertukaran nilai antarvariabel. Dengan memahami dasar-dasar tersebut, pembuatan program sederhana menggunakan bahasa Go dapat dilakukan dengan lebih terstruktur.

## Referensi

1. Donovan, A. A. A., & Kernighan, B. W. (2015). *The Go Programming Language*. Boston: Addison-Wesley Professional. Diakses melalui https://www.gopl.io/
2. The Go Authors. (2026). *The Go Programming Language Specification*. Diakses melalui https://go.dev/ref/spec
3. The Go Authors. (2026). *A Tour of Go*. Diakses melalui https://go.dev/tour/
