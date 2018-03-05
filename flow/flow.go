package flow

import (
	"fmt"
)

func IfTest() {
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

func ForTest() {
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

//Switch
func SwitchExample() {
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
