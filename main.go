package main

//fmt imprime en pantalla
import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	//declaracion de variables
	var name string
	//:= crea una variable; = asigna la variable
	lastname := "<Modificar con mi apellido>"
	var number = 100
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Print("Escribe tu nombre: ")
	fmt.Printf(helloWorld, name, lastname)
	fmt.Println(number, a, b, c)
}
