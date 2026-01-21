package main

import (
	util "arsistema/utilidades"
	"fmt"
	"strings"

	
)

func main()  {
	var (
		cedula string = "32.460.448"
		nombre string = "juanito"
		apellido string = "alcachofa"
		edad int = 20
		altura float32 = 1.69
		matricula float64 = 1324.66
		sexo bool = false // false = masculino, true = femenino
		turno int = 2 // 1 = mañana, 2 = tarde y 3 = Noche
	)

	verDatos(cedula, nombre, apellido, edad, altura, matricula, sexo, turno)
}

func verDatos(datos ...any)  {
	util.Encabezado("Datos del alumno", 46)
	fmt.Println("Cédula :", datos[0])
	nombreCompleto := strings.ToUpper(datos[1].(string)) + " " + strings.ToUpper(datos[2].(string))
	fmt.Printf("Nombre Completo : %v \n", nombreCompleto )
	fmt.Printf("Edad : %v años \n", datos[3])
	fmt.Printf("Altura : %v mts \n", datos[4])
	fmt.Printf("Matricula : %v$ \n", datos[5])
	fmt.Println("Sexo :", ObtenerSexo(datos[6].(bool)))
	fmt.Println("Turno :", ObtenerTurno(datos[7].(int)))
	util.Linea(46)
}

func ObtenerSexo(sexo bool) string {
	if sexo {
		return "Femenino"
	} else {
		return "Masculino"
	}
}

func ObtenerTurno(turno int) string {
	 switch turno {
	 case 1: return "Mañana"
	 case 2: return  "Tarde"
	 case 3: return  "noche"
	 default: return "No tiene turno"
		
	 }
}