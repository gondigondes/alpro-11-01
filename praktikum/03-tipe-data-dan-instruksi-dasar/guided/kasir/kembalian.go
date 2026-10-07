package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    sepuluhRibu := x / 10000
    x = x % 10000

    limaRibu := x / 5000
    x = x % 5000

    seribu := x / 1000

    fmt.Println(sepuluhRibu, limaRibu, seribu)
}
