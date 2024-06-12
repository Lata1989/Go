package main

import (
	"fmt"
	"strings"
)

func main() {

	// Ciclo for

	acumuladorImpar := 0
	acumuladorPar := 0
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			acumuladorPar = acumuladorPar + i
		} else {
			acumuladorImpar = acumuladorImpar + i
		}
	}

	fmt.Println("La suma de los impares entre 0 y 100 es igual a", acumuladorImpar)
	fmt.Println("La suma de los pares entre 0 y 100 es igual a", acumuladorPar)

	// For Each

	var miMapa = map[string]string{
		"Colombia":  "Bogotá",
		"Venezuela": "Caracas",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago de Chile",
	}

	for k, v := range miMapa {
		fmt.Println("La capital de " + k + " es " + v + ".")
	}

	// Do While
	fruta := ""
	for {
		fmt.Print("Ingrese una fruta: ")
		fmt.Scan(&fruta)
		fruta = strings.ToLower(fruta)
		fmt.Print("\n")
		if fruta == "naranja" {
			fmt.Println("El usuario escribió naranja.")
			break
		}
	}

	// While
	fruta = ""
	fmt.Print("Ingrese una fruta: ")
	fmt.Scan(&fruta)
	fruta = strings.ToLower(fruta)
	fmt.Print("\n")
	for {
		if fruta == "naranja" {
			fmt.Println("El usuario escribió naranja.")
			break
		}
	}

	// Ultima linea de codigo del main
}
