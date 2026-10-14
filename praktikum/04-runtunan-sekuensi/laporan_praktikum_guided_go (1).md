# <h1 align="center">Laporan Praktikum Guided - Bahasa Pemrograman Go</h1>
<p align="center">Akmal Ghani Laoda - 109092600005</p>

## Guided

### 1. `grade.go`

```go
package main

import (
    "bufio" // Digunakan untuk membaca masukan string yang memiliki spasi
    "fmt"
    "os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin (keyboard)
)

func main() {
    var nama string
    var nilai float64
    var grade string

    // Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
    scanner := bufio.NewScanner(os.Stdin)

    // Proses membaca masukan dari pengguna sampai tombol Enter ditekan
    scanner.Scan()

    // Mengambil teks yang baru saja dibaca dan menyimpannya ke variabel nama
    nama = scanner.Text()

    fmt.Scan(&nilai)

    if nilai >= 90 && nilai <= 100 {
        grade = "A"
    } else if nilai >= 80 && nilai < 90 {
        grade = "B"
    } else if nilai >= 70 && nilai < 80 {
        grade = "C"
    } else if nilai >= 60 && nilai < 70 {
        grade = "D"
    } else {
        grade = "F"
    }

    // %s adalah format penulisan (placeholder) untuk mencetak nilai bertipe string
    // %s pertama akan diisi oleh variabel 'nama', %s kedua diisi oleh 'grade'
    fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```

##### Output

![Screenshot Output Guided](guided/grade/image.png)

#### Deskripsi

Program `grade.go` digunakan untuk menentukan grade berdasarkan nilai yang dimasukkan oleh pengguna. Program membaca nama lengkap menggunakan `bufio.NewScanner(os.Stdin)` agar masukan nama yang mengandung spasi dapat dibaca sebagai satu baris. Setelah itu, program membaca nilai dalam bentuk `float64` menggunakan `fmt.Scan()`. Percabangan `if`, `else if`, dan `else` digunakan untuk menentukan grade: nilai 90 sampai 100 mendapat A, 80 sampai kurang dari 90 mendapat B, 70 sampai kurang dari 80 mendapat C, dan 60 sampai kurang dari 70 mendapat D. Nilai di luar rentang tersebut mendapat F. Terakhir, `fmt.Printf()` menampilkan nama dan grade yang diperoleh. Sebagai contoh, jika nama yang dimasukkan adalah `Akmal Ghani` dan nilainya `85`, keluaran yang dihasilkan adalah `Akmal Ghani mendapatkan nilai B`.

### 2. `penilaian.go`

```go
package main

import (
    "bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
    "fmt"
    "os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
    // Pendefinisian variabel secara eksplisit beserta tipe datanya
    var pilihan int
    var nama string
    var nilai float64
    var grade string

    // Menampilkan cetakan Menu ke layar
    fmt.Println("============= Menu =============")
    fmt.Println("1. Sistem penilaian")
    fmt.Println("0. Keluar")
    fmt.Print("Pilih opsi: ")

    // Membaca masukan pilihan (angka).
    // Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
    // ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
    fmt.Scanln(&pilihan)

    // Percabangan/Sekuensi bersyarat berdasarkan input pengguna
    if pilihan == 1 {
        fmt.Print("Masukkan nama siswa: ")

        // Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
        scanner := bufio.NewScanner(os.Stdin)

        // Proses membaca masukan dari pengguna sampai tombol Enter ditekan
        scanner.Scan()

        // Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
        nama = scanner.Text()

        fmt.Print("Masukkan nilai siswa: ")
        fmt.Scanln(&nilai)

        // Menentukan grade nilai
        if nilai >= 90 && nilai <= 100 {
            grade = "A"
        } else if nilai >= 80 && nilai < 90 {
            grade = "B"
        } else if nilai >= 70 && nilai < 80 {
            grade = "C"
        } else if nilai >= 60 && nilai < 70 {
            grade = "D"
        } else {
            grade = "F"
        }

        // %s adalah format (placeholder) untuk mencetak data bertipe string
        // %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
        fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

    } else if pilihan == 0 {
        // Dijalankan jika pengguna mengetik 0
        fmt.Println("Keluar dari program.")

    } else {
        // Dijalankan jika pengguna mengetik angka selain 1 dan 0
        fmt.Println("Pilihan tidak valid.")
    }
}
```

##### Output

![Screenshot Output Guided](guided/penilaian/image.png)

#### Deskripsi

Program `penilaian.go` merupakan program penilaian siswa yang menyediakan menu sederhana. Pada awal program, pengguna memilih opsi 1 untuk menjalankan sistem penilaian atau opsi 0 untuk keluar. Jika memilih opsi 1, program meminta nama siswa dan nilai siswa. Nama dibaca menggunakan `bufio.NewScanner(os.Stdin)` agar nama lengkap yang mengandung spasi dapat diterima, sedangkan nilai dibaca menggunakan `fmt.Scanln()`. Selanjutnya, percabangan `if`, `else if`, dan `else` menentukan grade dengan rentang nilai yang sama seperti program `grade.go`, yaitu A untuk 90–100, B untuk 80–kurang dari 90, C untuk 70–kurang dari 80, D untuk 60–kurang dari 70, dan F untuk nilai lainnya. Hasil penilaian ditampilkan menggunakan `fmt.Printf()`. Jika pengguna memilih 0, program menampilkan pesan bahwa program keluar. Jika pengguna memasukkan pilihan selain 0 dan 1, program menampilkan pesan `Pilihan tidak valid.`

### 3. `klasifikasi.go`

```go
package main

import (
    "fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
    // Pendefinisian variabel secara eksplisit beserta tipe datanya
    var usia int
    var gaji int
    var keterangan string

    // Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
    // fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
    fmt.Scan(&usia)
    fmt.Scan(&gaji)

    // Menggunakan switch tanpa ekspresi.
    // Cara kerjanya sama persis seperti deretan if - else if.
    // Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
    switch {
    case usia < 18:
        keterangan = "Masih sekolah"
    case usia >= 18 && usia <= 25 && gaji >= 50:
        keterangan = "Muda sukses"
    case usia >= 18 && usia <= 25 && gaji < 50:
        keterangan = "Masih belajar hidup"
    case usia >= 26 && usia <= 40 && gaji >= 100:
        keterangan = "Pekerja mapan"
    case usia >= 26 && usia <= 40 && gaji < 100:
        keterangan = "Perlu perbaikan karier"
    case usia > 40 && gaji >= 150:
        keterangan = "Profesional berpengalaman"
    case usia > 40 && gaji < 150:
        keterangan = "Perlu evaluasi finansial"
    }

    // Menampilkan hasil klasifikasi ke layar
    fmt.Println(keterangan)
}
```

##### Output

![Screenshot Output Guided](guided/klasifikasi/image.png)

#### Deskripsi

Program `klasifikasi.go` digunakan untuk mengelompokkan seseorang berdasarkan usia dan jumlah gaji yang dimasukkan. Program membaca usia dan gaji sebagai bilangan bulat menggunakan `fmt.Scan()`, kemudian menggunakan `switch` tanpa ekspresi untuk memeriksa kondisi secara berurutan. Jika usia kurang dari 18 tahun, keterangan yang ditampilkan adalah `Masih sekolah`. Untuk usia 18–25 tahun, gaji minimal 50 menghasilkan keterangan `Muda sukses`, sedangkan gaji di bawah 50 menghasilkan `Masih belajar hidup`. Untuk usia 26–40 tahun, gaji minimal 100 menghasilkan `Pekerja mapan`, sedangkan gaji di bawah 100 menghasilkan `Perlu perbaikan karier`. Jika usia lebih dari 40 tahun, gaji minimal 150 menghasilkan `Profesional berpengalaman`, sedangkan gaji di bawah 150 menghasilkan `Perlu evaluasi finansial`. Hasil klasifikasi ditampilkan menggunakan `fmt.Println()`. Program mengasumsikan masukan usia dan gaji berada pada rentang yang sesuai dengan kondisi yang telah ditentukan.

## Kesimpulan

Berdasarkan tiga program guided yang telah dibuat, dapat dipahami penerapan percabangan dalam bahasa Go untuk menyelesaikan permasalahan penentuan nilai, sistem penilaian siswa, dan klasifikasi berdasarkan usia serta gaji. Program `grade.go` menggunakan percabangan `if`, `else if`, dan `else` untuk menentukan grade dari nilai yang dimasukkan. Program `penilaian.go` menerapkan percabangan untuk menjalankan menu, memproses penilaian siswa, atau keluar dari program. Sementara itu, program `klasifikasi.go` menggunakan `switch` tanpa ekspresi untuk menentukan keterangan berdasarkan kombinasi kondisi usia dan gaji. Melalui praktikum ini, penggunaan variabel, input dan output, operator perbandingan, operator logika, serta struktur kontrol percabangan dalam bahasa Go dapat dipahami dan diterapkan.

## Referensi

1. Donovan, A. A. A., & Kernighan, B. W. (2015). *The Go Programming Language*. Boston: Addison-Wesley Professional. Diakses melalui https://www.gopl.io/
2. The Go Authors. (2026). *The Go Programming Language Specification*. Diakses melalui https://go.dev/ref/spec
3. The Go Authors. (2026). *A Tour of Go*. Diakses melalui https://go.dev/tour/
