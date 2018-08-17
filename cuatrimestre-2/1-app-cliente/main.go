package main

/*
	Application name	Ejemplo
	API key				db72af200b346629a1bdd90e9f44d8e1
	Shared secret		a1828aa9a9263f18c0ab655fa64dcf52
	Registered to		emikohmann
*/

// dependencias
import (
	"fmt"
)

// constantes
const (

	// token de autenticacion
	apiKey = "db72af200b346629a1bdd90e9f44d8e1"

	// ubicacion de la api a consumir
	baseUrl = "http://ws.audioscrobbler.com/2.0/"
)

// funcion principal
func main() {
	fmt.Println("Ejemplo APP Cliente")
}