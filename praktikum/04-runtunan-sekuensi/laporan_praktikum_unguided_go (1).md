# <h1 align="center">Laporan Praktikum Unguided - Bahasa Pemrograman Go</h1>
<p align="center">Akmal Ghani Laoda - 109092600005</p>

## Unguided

### 1. `grade.go` (Modifikasi Menggunakan `switch`)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var nama string
    var nilai float64
    var grade string

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    nama = scanner.Text()

    fmt.Scan(&nilai)

    switch {
    case nilai >= 90 && nilai <= 100:
        grade = "A"
    case nilai >= 80 && nilai < 90:
        grade = "B"
    case nilai >= 70 && nilai < 80:
        grade = "C"
    case nilai >= 60 && nilai < 70:
        grade = "D"
    default:
        grade = "F"
    }

    fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```

##### Output

![Screenshot Output Unguided](unguided/grade/image.png)

#### Deskripsi

Program `grade.go` merupakan modifikasi dari program penentuan grade sebelumnya dengan mengganti struktur percabangan `if-else` menjadi `switch` tanpa ekspresi. Program membaca nama lengkap melalui `bufio.NewScanner(os.Stdin)`, kemudian membaca nilai menggunakan `fmt.Scan()`. Struktur `switch` memeriksa rentang nilai secara berurutan. Nilai 90 sampai 100 mendapat grade A, nilai 80 sampai kurang dari 90 mendapat B, nilai 70 sampai kurang dari 80 mendapat C, dan nilai 60 sampai kurang dari 70 mendapat D. Jika tidak memenuhi kondisi tersebut, grade yang diberikan adalah F. Terakhir, program menampilkan nama dan grade menggunakan `fmt.Printf()`. Contohnya, jika masukan nama adalah `Lee Haechan` dan nilai `80`, keluaran yang dihasilkan adalah `Lee Haechan mendapatkan nilai B`.

### 2. `pajak.go`

```go
package main

import "fmt"

func main() {
    var penghasilan float64
    var pajak float64

    fmt.Scan(&penghasilan)

    if penghasilan <= 50 {
        pajak = 0.05 * penghasilan
    } else if penghasilan <= 100 {
        pajak = (0.05 * 50) + (0.10 * (penghasilan - 50))
    } else if penghasilan <= 200 {
        pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * (penghasilan - 100))
    } else {
        pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * 100) + (0.20 * (penghasilan - 200))
    }

    fmt.Println(pajak)
}
```

##### Output

![Screenshot Output Unguided](unguided/pajak/image.png)

#### Deskripsi

Program `pajak.go` digunakan untuk menghitung total pajak penghasilan berdasarkan jumlah penghasilan yang dimasukkan dalam satuan juta rupiah. Program menerima input bertipe `float64` menggunakan `fmt.Scan()`, kemudian menentukan pajak dengan percabangan `if-else if-else`. Jika penghasilan tidak lebih dari 50 juta, pajak dihitung sebesar 5% dari seluruh penghasilan. Jika penghasilan lebih dari 50 juta hingga 100 juta, pajak dihitung sebesar 5% untuk 50 juta pertama dan 10% untuk sisanya. Untuk penghasilan lebih dari 100 juta hingga 200 juta, perhitungan mencakup 5% dari 50 juta pertama, 10% dari 50 juta berikutnya, dan 15% dari sisa penghasilan di atas 100 juta. Jika penghasilan lebih dari 200 juta, pajak juga mencakup 20% dari bagian penghasilan di atas 200 juta. Hasil perhitungan disimpan dalam variabel `pajak` dan ditampilkan menggunakan `fmt.Println()`. Sebagai contoh, masukan `75` menghasilkan keluaran `5`, yang berarti pajak sebesar 5 juta rupiah.

## Kesimpulan

Berdasarkan dua program unguided yang telah dibuat, dapat dipahami penerapan struktur kontrol dalam bahasa Go untuk menyelesaikan permasalahan penentuan grade dan perhitungan pajak penghasilan. Program `grade.go` memodifikasi percabangan sebelumnya dengan menggunakan `switch` tanpa ekspresi untuk menentukan grade berdasarkan rentang nilai. Sementara itu, program `pajak.go` menggunakan percabangan `if-else if-else` untuk menghitung pajak secara bertingkat sesuai kategori penghasilan. Melalui praktikum ini, penggunaan variabel, input dan output, operator perbandingan, operasi aritmatika, serta struktur kontrol percabangan dalam bahasa Go dapat dipahami dan diterapkan.

## Referensi

1. Donovan, A. A. A., & Kernighan, B. W. (2015). *The Go Programming Language*. Boston: Addison-Wesley Professional. Diakses melalui https://www.gopl.io/
2. The Go Authors. (2026). *The Go Programming Language Specification*. Diakses melalui https://go.dev/ref/spec
3. The Go Authors. (2026). *A Tour of Go*. Diakses melalui https://go.dev/tour/
