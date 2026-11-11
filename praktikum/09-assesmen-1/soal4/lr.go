package main

import "fmt"

func main() {
	var L, R int
	fmt.Scan(&L, &R)

	jumlahValid := 0

	for bilangan := L; bilangan <= R; bilangan++ {
		temp := bilangan
		jumlahDigit := 0
		digitKanan := -1
		tidakMemenuhi := false

		for temp > 0 {
			digit := temp % 10
			jumlahDigit += digit

			if digitKanan != -1 && digit > digitKanan {
				tidakMemenuhi = true
			}

			digitKanan = digit
			temp /= 10
		}

		prima := true
		if jumlahDigit < 2 {
			prima = false
		} else {
			pembagi := 2
			for pembagi*pembagi <= jumlahDigit && prima {
				if jumlahDigit%pembagi == 0 {
					prima = false
				}
				pembagi++
			}
		}

		if !tidakMemenuhi && prima {
			jumlahValid++
		}
	}

	fmt.Println(jumlahValid)
}
