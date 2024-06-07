package main

import "fmt"

func main() {
	var fruta string

	fmt.Print("Ingrese un valor alfanumérico: ")
	fmt.Scan(&fruta)

	switch fruta {
	case "manzana":
		fmt.Println("Ingresaste manzana.")
	case "banana":
		fmt.Println("Ingresaste banana")
	default:
		fmt.Println("La palabra que ingresaste", fruta)
	}

}
