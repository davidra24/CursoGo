package structs

import "fmt"

func GetArray() {
	var arr1 [2]string
	arr1[0] = "Espacio 0 de array"
	arr1[1] = "Espacio 1 de array"
	fmt.Println(arr1)
}

func GetArray2() {
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
}

func GetSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}
