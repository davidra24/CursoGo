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
	fmt.Printf(helloWorld, name, lastname)
	fmt.Print("Hola mundo")
	fmt.Println(number, a, b, c)
	fmt.Println(integer1, integer2, integer3)
	fmt.Println(float1, float2)
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
