package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hola mundo!!")

	/*
		Hay 3 formas de declarar variables
	*/

	var nombrePersona string = "Latita"

	fmt.Println("Hola gente de Go! Soy " + nombrePersona)

	var apellidoPersona = "de la Iglesia"

	fmt.Println("Mi apellido es " + apellidoPersona)

	segundoNombre := "Alejandro"

	fmt.Println("Mi segundo nombre es " + segundoNombre)

	/*
		Parte numerica
	*/
	var asoActual int16 = 2024

	fmt.Println("El año actual es ", asoActual)

	var asoReducido int8 = 24

	fmt.Println("El año abreviado es ", asoReducido)

	edad := 35

	fmt.Println("Mi edad es ", edad)

	/*
		Arreglos
		Tiene dos formas
	*/

	var listaVerdura = [4]string{}

	fmt.Println(listaVerdura)

	var listaFruta = [4]string{"Pera", "Manzana", "Naranja", "Melon"}

	fmt.Println(listaFruta)
	fmt.Println(listaFruta[0])
	fmt.Println(listaFruta[1])
	fmt.Println(listaFruta[2])
	fmt.Println(listaFruta[3])

	listaPaises := [3]string{}
	listaPaises[0] = "Argentina"

	fmt.Println(listaPaises)

	/*
		Slices es identico a un arreglo pero con tamaño definido
	*/

	listaColores := []string{}

	listaColores = append(listaColores, "Negro")
	listaColores = append(listaColores, "Azul")
	listaColores = append(listaColores, "Rojo")
	listaColores = append(listaColores, "Violeta")
	listaColores = append(listaColores, "Indigo")
	listaColores = append(listaColores, "Turquesa")

	fmt.Println("Los colores de la lista son: ", listaColores)

	listaColores2 := listaColores[1:4]

	fmt.Println(listaColores2)

	listaColores2 = listaColores[:3]

	fmt.Println(listaColores2)

	listaColores2 = listaColores[4:]

	fmt.Println(listaColores2)
}
