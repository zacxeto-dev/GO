package main

import (
	util "arsistema/utilidades"
	"fmt"
	"math/rand/v2"
	"strings"
)

func main()  {
	var (
		secreto int = rand.IntN(20) + 1
	    numero int 
	 	alias string
	)

	// println("El numero secreto:", secreto)
	util.Encabezado("Juego de adivinar el Numero", 50)
	alias = util.LeerTexto("Ingresa tu alias", "alias")
	// fmt.Println("El alias es", alias)
	util.Linea(50)
	jugar(secreto, numero, alias) 

}

func jugar(secreto, numero int, alias string) {
	var msj string
	intentos := 1
	alias = strings.ToUpper(alias)

	for numero != secreto {
		if intentos == 5 {
			msj = "Perdiste por manco el numero secreto era"
			util.Linea(50)
			fmt.Printf("%v, %v %v \n", alias, msj, secreto)
			util.Linea(50)
		break
		}

		numero = util.LeerNumeroEntero("Ingresa tu numero")
		intentos++

		if numero < secreto {
			msj = "Tu Numero es menor que el numero Secreto"
			fmt.Printf("%v, %v \n", alias, msj)
		} else if numero > secreto{
			msj = "Tu numero es mayor que el numero Secreto"
			fmt.Printf("%v, %v \n", alias, msj)
		} else {
			msj = "YOU WIN..."
			util.Linea(50)
			fmt.Printf("%v, %v \n", alias, msj)
			util.Linea(50)
		}
	}
}