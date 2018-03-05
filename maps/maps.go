package maps

//Devuelve un map con llave tipo string y valores enteros
func GetMap() map[string]int {
	mapTest := make(map[string]int)
	mapTest["llave1"] = 0
	mapTest["llave2"] = 1
	mapTest["llave3"] = 2
	mapTest["llave4"] = 3
	mapTest["llave5"] = 4
	mapTest["llave6"] = 5
	//Elimina los valores de map
	delete(mapTest, "llave3")
	delete(mapTest, "llave4")
	return mapTest
}

func GetMap2() map[string]int {
	mapTest := map[string]int{
		"Juan":   0,
		"Pedro":  1,
		"Pablo":  2,
		"Johnny": 3,
	}
	return mapTest
}
