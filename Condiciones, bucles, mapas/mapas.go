package main

import "fmt"

func mapa() {

	/*
		Mapas: Go no garantiza el orden
	*/
	var miMapa = map[string]string{
		"Colombia":  "Bogotá",
		"Venezuela": "Caracas",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago de Chile",
	}

	fmt.Println("El mapa de paises es: ", miMapa)

	fmt.Println("La capital de Venezuela es", miMapa["Venezuela"])

	miMapa["Argentina"] = "Buenos Aires"

	fmt.Println("La capital de Argentina es", miMapa["Argentina"])

	delete(miMapa, "Venezuela")

	fmt.Println(miMapa)
}
