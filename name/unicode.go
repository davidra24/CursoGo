package name

import "fmt"

//GO maneja UTF-8 nativo
func GetUniCode() string {
	return "もしもし！"
}

func MakeUniCode(value string) {
	fmt.Println("Cadena UTF-8: ", value)
}
