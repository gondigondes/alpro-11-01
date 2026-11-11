package main

import "fmt"

func main() {
	var nilai, sebelumnya int
	var pertama bool = true
	var selesai bool = false
	var jumlah, banyak int
	var maks, min int

	for !selesai {
		fmt.Scan(&nilai)

		if !pertama && nilai == sebelumnya {
			selesai = true
		} else {
			if pertama {
				maks = nilai
				min = nilai
				pertama = false
			} else {
				if nilai > maks {
					maks = nilai
				}
				if nilai < min {
					min = nilai
				}
			}

			jumlah += nilai
			banyak++
			sebelumnya = nilai
		}
	}

	rataRata := float64(jumlah) / float64(banyak)
	fmt.Printf("%d %d %.2f\n", maks, min, rataRata)
}
