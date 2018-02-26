package main

//fmt imprime en pantalla
import "fmt"

func main() {
	var name string
	fmt.Print("Escribe tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hola ", name)
}
