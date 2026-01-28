package main

import (
	util "arsistema/utilidades"
	"fmt"
)

func main()  {
	var (
		altura float64
		sumatoria float64
		maximo float64
		minimo float64
		promedio float64
	)

	const nroPersonas int = 6

	avisos := [6]string{
		"primera",
		"segunda",
		"tercera",
		"cuarta",
		"quinta",
		"sexta",
	}

	util.Encabezado("Ingresos de altura", 46)
	for i := range nroPersonas {
		altura = util.LeerNumeroDecimales("Ingresa " + avisos[i] + " altura")
		sumatoria += altura

		if i == 0 {
			maximo = altura
			minimo = altura
		}

		if altura > maximo {
			maximo = altura
		}

		if altura < minimo {
			minimo = altura
		}
	}

	promedio = sumatoria / float64(nroPersonas)

	util.Linea(46)
	fmt.Printf("El promedio de altura es %.2f mts \n", promedio)
	util.Linea(46)
	fmt.Printf("La maxima altura es %.2f mts \n", maximo)
	util.Linea(46)
	fmt.Printf("La minima altura es %.2f \n", minimo)
	util.Linea(46)
}