package name

import "fmt"

// Si empieza en mayuscula una funcion es publica
//Si empieza en minuscula una funcion es privada
func GetName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}
func GetLastName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu apellido: ")
	fmt.Scanf("%s", &name)
	return name
}
