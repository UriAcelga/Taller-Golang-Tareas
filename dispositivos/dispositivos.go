package dispositivos

import (
	"errors"
	"math/rand"

	"github.com/fatih/color"
)

type dispositivo struct {
	Nombre string
	Estado bool
}

func New(nombre string) *dispositivo {
	return &dispositivo{
		Nombre: nombre,
		Estado: (rand.Int() % 2) == 1,
	}
}

func (d *dispositivo) Encender() error {
	if d.Estado {
		return errors.New(d.Nombre + " ya estaba encendido")
	}
	d.Estado = true
	return nil
}

func (d *dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New(d.Nombre + " ya estaba apagado")
	}

	d.Estado = false
	return nil
}

func (d *dispositivo) EstadoActual() string {
	if d.Estado {
		color.Green(d.Nombre + " Encendido")
		return d.Nombre + " Encendido"
	} else {
		color.Black(d.Nombre + " Apagado")
		return d.Nombre + " Apagado"
	}
}

type Controlable interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}
