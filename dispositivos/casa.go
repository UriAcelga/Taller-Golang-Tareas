package dispositivos

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func GestionCasa() error {
	electrodomesticos := []*dispositivo{
		New("Lavarropas"),
		New("Heladera"),
		New("Aire Acondicionado"),
		New("Cafetera"),
		New("Alarma"),
		New("Luces exteriores"),
		New("Luces interiores"),
	}

MenuDispositivos:
	for {
		opcion := menuDispositivo()
		fmt.Printf("\n")
		switch opcion {
		case 1:
			err := mostrarDispositivos(electrodomesticos)
			if err != nil {
				fmt.Println(err)
			}
			_, err = fmt.Scanln(&opcion)
			if err != nil {
				return err
			}

			if opcion >= 0 && opcion < len(electrodomesticos) {
				err = electrodomesticos[opcion].Encender()
				if err != nil {
					color.Red(err.Error())
				} else {
					color.Yellow(electrodomesticos[opcion].Nombre + " Encendido\n")
				}
			}

		case 2:
			err := mostrarDispositivos(electrodomesticos)
			if err != nil {
				fmt.Println(err)
			}
			_, err = fmt.Scanln(&opcion)
			if err != nil {
				return err
			}

			if opcion >= 0 && opcion < len(electrodomesticos) {
				err = electrodomesticos[opcion].Apagar()
				if err != nil {
					color.Red(err.Error())
				} else {
					color.Yellow(electrodomesticos[opcion].Nombre + " Apagado\n")
				}
			}
		case 3:
			for _, v := range electrodomesticos {
				v.EstadoActual()
			}
		case 0:
			break MenuDispositivos

		}
	}
	return nil
}

func menuDispositivo() int {
	fmt.Println("\n◄----------------►")
	fmt.Println("► 1. Encender dispositivo")
	fmt.Println("► 2. Apagar dispositivo")
	fmt.Println("► 3. Ver estado")
	fmt.Println("► 0. Salir")
	fmt.Println("◄----------------►")
	fmt.Printf("► Opción: ")
	var opcion int
	_, err := fmt.Scanln(&opcion)
	if err != nil {
		fmt.Println("Scanf retorno con error: " + err.Error())
	}
	return opcion

}

func mostrarDispositivos(d []*dispositivo) error {
	if len(d) == 0 {
		return errors.New("no hay dispositivos")
	}
	fmt.Println("◄- Elige un dispositivo -►")
	fmt.Println("◄----------------►")
	for i, v := range d {
		fmt.Printf("► %d. %s\n", i, v.Nombre)
	}
	fmt.Println("◄----------------►")
	fmt.Printf("► Opción: ")
	return nil
}
