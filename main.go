package main

//Si no se utilzia una libreria no va a compilar
import (
	"./flow"
	"./name"
	"./numbers"
	"./structs"
)

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {

	//Uso del package numbers
	number := numbers.Sum(50, 50)
	a, b, c := numbers.GetVariables()
	integer1, integer2, integer3 := numbers.GetIntegers()
	float1, float2 := numbers.GetFloats()
	numbers.PrintNumbers(number, a, b, c)
	numbers.PrintValues(integer1, integer2, integer3)
	numbers.PrintFloats(float1, float2)

	//uso del package name
	firstName := name.GetName()
	lastname := name.GetLastName()
	stringUTF8 := name.GetUniCode()
	name.MakeUniCode(stringUTF8)
	name.Saludo(helloWorld, firstName, lastname)
	name.Strings2()

	//Uso del package structs
	structs.GetArray()
	structs.GetArray2()
	structs.GetSlice()

	//Uso del package flow
	flow.IfTest()
	flow.ForTest()
	flow.SwitchExample()
}
