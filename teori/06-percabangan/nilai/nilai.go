package main

import "fmt"

func main() {
	var nilai int
	fmt.Scan(&nilai)

	if nilai < 50 {
		fmt.Println("Nilai Anda: E")
	} else if nilai >= 50 && nilai < 70 {
		fmt.Println("Nilai Anda: C")
	} else if nilai >= 70 && nilai < 85 {
		fmt.Println("Nilai Anda: B")
	} else if nilai >= 85 && nilai <= 100 {
		fmt.Println("Nilai Anda: A")
	} else {
		fmt.Println("Nilai tidak valid (harus antara 0 - 100)")
	}
}


