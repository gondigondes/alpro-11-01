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