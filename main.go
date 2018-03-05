package main

//Si no se utilzia una libreria no va a compilar
import (
	"fmt"
	"strings"
)

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
	//Imprime array
	getArray()
	getArray2()
	getSlide()
	ifTest()
	forTest()
	strings2()
	switchExample()
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

func getArray() {
	var arr1 [2]string
	arr1[0] = "Espacio 0 de array"
	arr1[1] = "Espacio 1 de array"
	fmt.Println(arr1)
}

func getArray2() {
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
}

func getSlide() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}

func ifTest() {
	var numb int
	fmt.Print("Ingresa un numero del 1 al 10: ")
	fmt.Scanf("%d", &numb)
	if numb%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}
	if numb2 := numbr2(); numb2 == 3 {
		fmt.Println("Entró al condicional")
	}
}

func numbr2() int {
	return 3
}

func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("FOR: ", i)
	}
	//FOR.. while(?)
	h := 100
	for h > 0 {
		h -= 10
		fmt.Println("FOR: Solo cun una condicion de H > 0:", h)
	}
	//Ciclo infinito
	s := 10
	for {
		fmt.Println("Bucle infinito")
		if s == 0 {
			break
		} else {
			s--
		}

	}
}

//Operaciones con cadenas
func strings2() {
	var text = "Todo en mayusculas"
	fmt.Println(strings.ToUpper(text))
	text = "Todo en minusculas"
	fmt.Println(strings.ToLower(text))
	text = "Hello world, Hello GoLang, Hello amigo"
	fmt.Println(text)
	//variable, texto cambio, texto nuevo, -1 cambia todas las coincidencias
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	//Split function
	fmt.Println(strings.Split(text, ","))
	fmt.Println(strings.Split(text, " "))
	fmt.Println(strings.Split(text, ", "))
}

//Switch
func switchExample() {
	var number = 0
	fmt.Print("SWITCH: Ingresa un numero del 1 al 10: ")
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("El numero es 1")
	default:
		fmt.Println("El numero no es 1")
	}
	switch {
	case number%2 == 0:
		fmt.Println("El numero es par")
	default:
		fmt.Println("El numero es impar")
	}
	//Recomendacion de Google: Si un If pasa más de dos condiciones se debe usar el switch
}
