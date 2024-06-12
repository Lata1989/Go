package main

import "fmt"

func funciones() {
	fmt.Println("El total de la suma de los numeros es ", suma(7, 8))
	presentaResultadoSuma(13, 14)
	presentaResultadoSuma(1, 14)
	presentaResultadoSuma(33, 54)
	presentaResultado("sumar", 5, 9)

}

func suma(num1 int, num2 int) int {
	//total := num1 + num2
	return num1 + num2
}

func presentaResultadoSuma(a int, b int) {
	fmt.Println("El resultado de sumar ", a, " + ", b, " = ", suma(a, b))
}

func presentaResultado(funcion string, a int, b int) {
	restar := func(a int, b int) int {
		return a - b
	}
	/*
		sumar := func(a int, b int) int {
			return a + b
		}
	*/

	resultado := 0

	if funcion == "sumar" {
		resultado = suma(a, b)
	} else if funcion == "restar" {
		resultado = restar(a, b)
	} else {
		resultado = 0
	}

	fmt.Println("El resultado de la operacion es: ", resultado)
}
