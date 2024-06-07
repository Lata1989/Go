package main

import "fmt"

func condicionales() {
	var edad int = 70

	if edad < 18 {
		fmt.Println("No puede trabajar, es menor.")
	} else if edad >= 18 && edad < 65 {
		fmt.Println("Esta en edad económicamente activa.")
	} else {
		fmt.Println("Esta persona está jubilada.")
	}

	/*
		Operadores condicionales en Go

		Lógicos:

		&& And
		|| Or
		! Not

		Aritmeticos:

		Igual a: ==
		Distinto de: !=
		Mayor que: >
		Menor que: <
		Mayor o igual que: >=
		Menor o igual que: <=

	*/
}
