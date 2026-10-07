package main

import "fmt"

func main() {
    var celsius float64
    fmt.Scan(&celsius)

    kelvin := celsius + 273

    fmt.Println(kelvin)
}
