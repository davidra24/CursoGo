package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	lastname := "<Modificar con mi apellido>"
	name := getName()
	number := sum(50, 50)
	a, b, c := getVariables()
	integer1, integer2, integer3 := getIntegers()
	float1, float2 := getFloats()
	stringUTF8 := getUnitCode()
	fmt.Printf(helloWorld, name, lastname)
	fmt.Print("Hola mundo")
	fmt.Println(number, a, b, c)
	fmt.Println(integer1, integer2, integer3)
	fmt.Println(float1, float2)
	fmt.Println("Cadena UTF-8: ", stringUTF8)
	//Retorna el ASCII del char en posicion 0
	println("Hola"[0])
	//Retorna el char de la posicion 0
	println(string("Hola"[0]))
	//Cantidad de letras y/o carateres de un string
	println("Hola tiene: ", len("Hola"), " carácteres")
}

//GO maneja UTF-8 nativo
func getUnitCode() string {
	return "もしもし！"
}

func getName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getIntegers() (int, int32, int64) {
	return 10, 2147000000, 90313131313131313
}

func getFloats() (float32, float64) {
	return float32(0.1), float64(0.1)
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

func sum(a int, b int) int {
	return a + b
}
