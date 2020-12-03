package main

import "fmt"

func hola(nombre string) string {
	if len(nombre) == 0 {
		return "Hola que tal"
	} else {
		return fmt.Sprintf("Hola %s", nombre)
	}
}
