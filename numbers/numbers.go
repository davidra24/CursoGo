package numbers

import "fmt"

//Devuelve 3 numeros enteros
func GetIntegers() (int, int32, int64) {
	return 10, 2147000000, 90313131313131313
}

//Devuelve 2 numeros flotantes
func GetFloats() (float32, float64) {
	return float32(0.1), float64(0.1)
}

//Multiple return de numeros
func GetVariables() (int, int, int) {
	return 1, 2, 3
}

//Devuelve la suma de dos datos
func Sum(a int, b int) int {
	return a + b
}

func PrintNumbers(number int, a int, b int, c int) {
	fmt.Println(number, a, b, c)
}

func PrintValues(integer1 int, integer2 int32, integer3 int64) {
	fmt.Println(integer1, integer2, integer3)
}

func PrintFloats(float1 float32, float2 float64) {
	fmt.Println(float1, float2)
}
