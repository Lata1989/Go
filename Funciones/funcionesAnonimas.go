package main

import "fmt"

var unaFuncion = map[string]func(int, int) int{
	"suma":       func(a int, b int) int { return a + b },
	"resta":      func(a int, b int) int { return a - b },
	"multiplica": func(a int, b int) int { return a * b },
	"divide":     func(a int, b int) int { return a / b },
}

func presentarResultado(operacion string, a int, b int) {
	f := unaFuncion[operacion]
	//resultado := f(a, b) // Llama a la función con los argumentos a y b
	fmt.Println("El resultado de ", a, " "+operacion+" ", b, " = ", f(a, b))
}

func main() {
	presentarResultado("suma", 9, 5)
	presentarResultado("resta", 9, 5)
	presentarResultado("multiplica", 9, 5)
	presentarResultado("divide", 9, 5)
}
