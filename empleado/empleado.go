package main

import (
	util "arsistema/utilidades"
	"fmt"
	"strings"
)

func main()  {
	empleado := map[string]any{
		"cedula" : "32460448",
		"nombre" : "benito",
		"apellido" : "camela",
		"edad" : 25,
		"sueldo" : 1000.22,
	}

	util.Linea(50)
	fmt.Printf("Cédula: %v \n", empleado["cedula"])
	nombrecompleto := strings.ToUpper(empleado["nombre"].(string)) + " " + strings.ToUpper(empleado["apellido"].(string))
	fmt.Printf("Nombre completo: %v \n", nombrecompleto)
	util.Linea(50)

	util.Encabezado("Datos del empleado", 46)
	for clave, valor := range empleado {
		fmt.Printf("%v: %v \n", strings.ToUpper(clave), verDatos(clave, valor))
	}
	util.Linea(46)

}

func verDatos(clave string, valor any) any {
	switch clave {
	case "nombre": return strings.ToUpper(valor.(string))
	case "apellido": return strings.ToUpper(valor.(string))
	default: return valor
		
	}
}