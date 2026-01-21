package main

import (
	util "arsistema/utilidades"
	"fmt"
	"math"
	"strings"
)

func main()  {
	var (
		peso float64
		altura float64
		imc float64
	)
	util.Encabezado("Calculadora indice de masa corporal", 50)
	peso = util.LeerNumeroDecimales("Ingresa tu peso (Kg)")
	altura = util.LeerNumeroDecimales("Ingresa tu altura (Mts)")

	imc = calcularIMC(peso, altura)

	util.Linea(50)
	msj := "El indice de masa corporal es"
	fmt.Printf("%v %.2f (%v) \n", msj, imc, strings.ToUpper(EstatusIMC(imc)))
	util.Linea(50)
}

func calcularIMC(peso, altura float64) float64 {
	return  peso / math.Pow(altura, 2)
}

func EstatusIMC(imc float64) string {
	if imc < 18.5 {
		return "bajo peso"
	} else if imc >= 18.5 && imc < 24.99 {
		return "peso normal"
	} else if imc >= 24.99 && imc < 29.99 {
		return "sobrepeso"
	}else {
		return "obesidad"
	}
}