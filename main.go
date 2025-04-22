package main

import (
	"fmt"
	"taller/dispositivos"
	"taller/puntajes"
)

func main() {
menuPrincipal:
	for {
		opcion := mostrarMenu()
		switch opcion {
		case 1:
			puntajes.EvaluarPuntajes()
		case 2:
			dispositivos.GestionCasa()
		case 0:
			fmt.Println("Saliendo del programa.. Éxitos usando GO!")
			break menuPrincipal
		}
	}

}

func mostrarMenu() int {
	var opcion int
	fmt.Println("\n◄====================►")
	fmt.Println("► 1. Puntajes")
	fmt.Println("► 2. Dispositivos")
	fmt.Println("► 0. Finalizar")
	fmt.Println("◄====================►")
	fmt.Printf("► Opción: ")
	_, err := fmt.Scanln(&opcion)
	if err != nil {
		fmt.Println("Scanf retorno con error: " + err.Error())
	}
	return opcion
}
