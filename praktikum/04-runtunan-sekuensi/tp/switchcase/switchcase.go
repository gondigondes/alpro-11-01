package main

import "fmt"

func main() {
    var nilai string
    fmt.Scan(&nilai)

    switch nilai {
    case "A":
        fmt.Println("Sangat Baik")
    case "B":
        fmt.Println("Baik")
    case "C":
        fmt.Println("Cukup")
    case "D":
        fmt.Println("Kurang")
    default:
        fmt.Println("Nilai tidak valid")
    }
}