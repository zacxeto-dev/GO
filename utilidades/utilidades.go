package utilidades

import (
	"fmt"
	"strings"
)

func Centrar(texto string, largo int) string {
	relleno := (largo - len(texto)) / 2
	textoCentrado := strings.Repeat(" ", relleno) + texto
	return textoCentrado
}

func Linea(largo int)  {
	fmt.Println(strings.Repeat("═", largo))
}

func Encabezado(titulo string, largo int)  {
	Linea(largo)
	fmt.Println(Centrar(strings.ToUpper(titulo), largo))
	Linea(largo)
}

func LeerNumeroEntero(msj string) int {
	var numero int

   for {
	fmt.Print(msj + " ")

	_, err := fmt.Scanln(&numero)

	if err != nil {
		fmt.Println("Debes escribir un numero entero...")

		var limpiar string
		fmt.Scanln(&limpiar)
		continue
	}

	return numero

}

}

func LeerNumeroDecimales(msj string) float64 {
	var numero float64

	for {
		fmt.Print(msj + ": ")
		_, err := fmt.Scanln(&numero)

		if err != nil {
			fmt.Println("Debe escribir un numero")

			var limpiar string
			fmt.Scanln(&limpiar)
			continue
		}

		return numero
	}
}


func LeerTexto(msj string, aviso string) string {
	var texto string
	for {
		fmt.Print(msj + ": ")
		fmt.Scanln(&texto)

		if texto != "" {
			return texto
		
		} else {
			fmt.Printf("Debe Escribir %v... \n", aviso)
			continue
		}
	}
}

