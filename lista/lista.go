package main

/* estructura de datos

	conjuntos de datos

		arrays
		mapas
		slice
*/

import (
	"fmt"
	util "arsistema/utilidades"
	"slices"
)

func main()  {
	// Declarar un slice de edades
	edades := []int{21,19,11,18,30,20,45,60,12,11,40}

	util.Linea(50)
	fmt.Println("Cantidad de edades es", len(edades))
	util.Linea(50)
	fmt.Println("La tercera edad es", edades[2])
	util.Linea(50)
	fmt.Println("La quinta edad es", edades[4])
	util.Linea(50)
	fmt.Println("La ultima edad es", edades[len(edades) - 1])
	util.Linea(50)
	edades = append(edades, 14)
	edades = append(edades, 86)
	fmt.Println(edades)
	util.Linea(50)
	fmt.Println("Porción del slice edades:", edades[2:6])

	util.Encabezado("Leer la lista de edades", 50)
	for _, edad := range edades {
		fmt.Println("La edad es", edad)
	}

	util.Encabezado("Leer la lista de edades con Indices", 50)
	for i, edad := range edades {
		fmt.Printf("La edad %v es %v \n", i + 1, edad)
	}

	util.Encabezado("Lista de edades ordenadas ascendente", 50)
	slices.Sort(edades)
	for _, edad := range edades {
		fmt.Println("La edad es", edad)
	}

	util.Encabezado("Lista de edades ordenadas decendente", 50)
	slices.SortFunc(edades, func (a, b int) int {
		return b - a
	})
	for _, edad := range edades {
		fmt.Println("la edad es", edad)
	}

	util.Encabezado("Lista de mayores de edad", 50)
	var mayoresEdad = []int{}
	for _, edad := range edades {
		if edad >= 18 {
			mayoresEdad = append(mayoresEdad, edad)
		}
	}
	for _, edad := range mayoresEdad {
		fmt.Println("La edad es", edad)
	}

	util.Encabezado("Lista de meneros de edad", 50)
	var menoresEdad = []int{}
	for _, edad := range edades {
		if edad < 18 {
			menoresEdad = append(menoresEdad, edad)
		}
	}
	for _, edad := range menoresEdad {
		fmt.Println("La edad es", edad)
	}

}