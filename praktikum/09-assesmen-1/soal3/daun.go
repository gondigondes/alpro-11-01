package main

import "fmt"

func main() {
	var n, lebar, min int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&lebar)

		if i == 0 || lebar < min {
			min = lebar
		}
	}

	fmt.Println(min)
}
