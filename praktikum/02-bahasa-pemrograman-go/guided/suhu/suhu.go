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