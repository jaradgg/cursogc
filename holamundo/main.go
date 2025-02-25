package main

import (
	"fmt"
)

/* 
comentario de bloque 
*/

//formato de linea


// declarar e inicializar variables
var (
	language = "Go"	
	company = "Google"
	year = 2009
)

// constantes
const pi = 3.1416

func main() {

	var nombre string //declarar variable
	nombre = "Jorge"
	apellido := "Della Gaspera" //declarar e inicializar variable en el cuerpo de una función
	fmt.Println("Hola Mundo, ", nombre +" "+ apellido)

	fmt.Println("Lenguaje", language)
	fmt.Println("Compañia", company)
	fmt.Println("Año", year)
	fmt.Println("Pi", pi)

	// mostrar datos en una línea
	fmt.Printf("Lenguaje %s creado por %s en %d\n", language, company, year)

	fmt.Println("Hola Mundo")
}
