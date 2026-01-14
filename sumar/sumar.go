package main

/* principios de dise;o 
 1- no te repitas
 2-no usar valores literales (estaticos)
 3-modular

   abrir terminal
    crear modulo: go mod init arsistema
	actualizar modulo: go mod tidy  */

import (
	"fmt"
	util "arsistema/utilidades"
)

func main ()  {
	var (
		numero1 int = 9
		numero2 int = 8
		suma int
	)

	util.Encabezado("sumar numeros", 36)

	numero1 = util.LeerNumeroEntero("Ingresa primer numero:")
	numero2 = util.LeerNumeroEntero("Ingresa segundo numero:")

	suma = numero1 + numero2

	util.Linea(36)
	fmt.Println("la suma es", suma)
	util.Linea(36)
}

