package main

import (
	"fmt"

	"./maps"
)

type Estructura struct {
	nombre string
	slug   string
	skills []string
}

//Una estructura puede "heredar" de otra
type EstructuraHerencia struct {
	Estructura
}

func main() {
	m0 := maps.GetMap()
	fmt.Println(m0)
	m1 := maps.GetMap2()
	fmt.Println(m1)

	estructura := Estructura{nombre: "Go", slug: "golang.com", skills: []string{"1", "2"}}
	fmt.Println(estructura)

	estructura1 := new(Estructura)
	estructura1.nombre = "GO1"
	estructura1.slug = "Golang1.com"
	estructura1.skills = []string{"1", "2", "3"}
}
