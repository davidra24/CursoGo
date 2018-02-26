package main

//fmt imprime en pantalla
import "fmt"

func main() {
	var name string
	fmt.Print("Escribe tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hola %s, bienvenido al mundo de GO \n", name)
}
