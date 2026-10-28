package main

import "fmt"

func main() {
	lamp_status := "merah"

	if lamp_status == "merah" {
		fmt.Println("berhenti")
	} else if lamp_status == "kuning" {
		fmt.Println("hati-hati")
	} else if lamp_status == "hijau" {
		fmt.Println("jalan")
	} else {
		fmt.Println("lampu tidak valid")
	}
}