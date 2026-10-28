package main

import "fmt"

func main() {
	kapasitasMaksimum := 50
	jumlahKendaraan := 0

	fmt.Println("=== SISTEM PENCATATAN PARKIR KAMPUS ===")

	for jumlahKendaraan < kapasitasMaksimum {
		jumlahKendaraan++
		fmt.Printf("Kendaraan ke-%d masuk dan dicatat oleh petugas.\n", jumlahKendaraan)
	}

	fmt.Println("\nArea parkir penuh! Petugas memasang papan 'Parkir Penuh'.")
}