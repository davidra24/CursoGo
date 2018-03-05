package name

import (
	"fmt"
	"strings"
)

func Saludo(helloWorld string, firstName string, lastname string) {
	//Operaciones y multiples Strings
	fmt.Printf(helloWorld, firstName, lastname)
	fmt.Print("Hola mundo")
	//Retorna el ASCII del char en posicion 0
	println("Hola"[0])
	//Retorna el char de la posicion 0
	println(string("Hola"[0]))
	//Cantidad de letras y/o carateres de un string
	println("Hola tiene: ", len("Hola"), " carácteres")
}

//Operaciones con cadenas
func Strings2() {
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
