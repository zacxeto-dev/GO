package main

import (
	"fmt"
	util "arsistema/utilidades"
	"strconv"
)

func main()  {
	var (
		numero int
		inicio int
		final int
		titulo string
	)

	util.Encabezado("Gnerador de tablas de multiplicar", 50)
	numero = util.LeerNumeroEntero("Ingresa tu numero :")
	titulo = "Tabla de Multiplicar del " + strconv.Itoa(numero)
	util.Encabezado(titulo, 50)
	generarTabla(numero)

	util.Encabezado("Generar rango de tablas", 50)
	inicio = util.LeerNumeroEntero("Ingresa numero de inicio")
	final = util.LeerNumeroEntero("Ingresa numero final")
	titulo = "Tablas de multiplicar del" + strconv.Itoa(inicio) + " al " + strconv.Itoa(final)
	util.Encabezado(titulo, 50)
	generarRangoTablas(inicio, final)
}

func generarRangoTablas(inicio, final int) {
	for i := inicio; i <= final; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%v x %v = %v \n", i, j, producto(i, j))
		}
		util.Linea(50)
	}
}

func generarTabla(numero int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v \n", numero, i, producto(numero, i))
	}
}

func producto(a, b int) int {
	return  a * b
}