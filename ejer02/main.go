package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool = true

func main() {
	num2, num3, num4, texto, status := 2, 5, 67, "Este es mi texto", false

	var moneda float64 = 0
	num2 = int(moneda)

	texto = strconv.Itoa(int(moneda))

	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(texto)
	fmt.Println(status)

	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(status)
}
