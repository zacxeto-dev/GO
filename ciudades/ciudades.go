package main

import (
	"fmt"
	util "arsistema/utilidades"
	"strings"
	"slices"
)

func main()  {
	ciudades := []string{}  // Slice
	avisos := [6]string{
		"primera",
		"segunda",
		"tercera",
		"cuarta",
		"quinta",
		"sexta",
	} // Array
	var ciudad string
	const nroCiudades int = 6

	VerCiudades(ciudades, "Ingreso de Ciudades", 50)
	for i := range nroCiudades {
		ciudad = util.LeerTexto("Ingresa " + avisos[i] + " ciudad", "ciudad")
		ciudades = append(ciudades, ciudad)
	}

	VerCiudades(ciudades, "Lista de ciudades", 46)
	

	slices.Sort(ciudades)
	VerCiudades(ciudades, "Lista de ciudades ascendente", 46)
	

	slices.SortFunc(ciudades, func(a, b string) int {
		return strings.Compare(b, a)
	})
	VerCiudades(ciudades, "Lista de ciudades descendente", 46)
	
}

func VerCiudades(ciudades []string, titulo string, largo int)  {
	util.Encabezado(titulo, largo)
	for _, ciudad := range ciudades {
		fmt.Println("La ciudad es", strings.ToUpper(ciudad))
	}
}