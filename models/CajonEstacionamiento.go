package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

type CajonEstacionamiento struct {
	ubicacionDelCajon         *floatgeom.Rect2
	direccionesParaEstacionar *[]CajonEstacionamientoDireccion
	direccionesParaSalir      *[]CajonEstacionamientoDireccion
	numeroDeCajon             int
	disponible                bool
}

func NewCajonEstacionamiento(x, y, x2, y2 float64, column, number int) *CajonEstacionamiento {
	direccionesEstacionar := getDireccionParaEstacionar(x, y, column)
	direccionesSalir := getDireccionesParaSalir()
	ubicacion := floatgeom.NewRect2(x, y, x2, y2)

	return &CajonEstacionamiento{
		ubicacionDelCajon:         &ubicacion,
		direccionesParaEstacionar: direccionesEstacionar,
		direccionesParaSalir:      direccionesSalir,
		numeroDeCajon:             number,
		disponible:                true,
	}
}

func getDireccionParaEstacionar(posX, posY float64, columna int) *[]CajonEstacionamientoDireccion {
	var direcciones []CajonEstacionamientoDireccion

	if columna == 1 {
		direcciones = append(direcciones, *newDireccionCajon("left", 445))
	} else if columna == 2 {
		direcciones = append(direcciones, *newDireccionCajon("left", 355))
	} else if columna == 3 {
		direcciones = append(direcciones, *newDireccionCajon("left", 265))
	} else if columna == 4 {
		direcciones = append(direcciones, *newDireccionCajon("left", 175))
	} else if columna == 5 {
		direcciones = append(direcciones, *newDireccionCajon("left", 85))
	}

	direcciones = append(direcciones, *newDireccionCajon("down", posY+5))
	direcciones = append(direcciones, *newDireccionCajon("left", posX+5))

	return &direcciones
}

func getDireccionesParaSalir() *[]CajonEstacionamientoDireccion {
	var direcciones []CajonEstacionamientoDireccion

	direcciones = append(direcciones, *newDireccionCajon("down", 380))
	direcciones = append(direcciones, *newDireccionCajon("right", 475))
	direcciones = append(direcciones, *newDireccionCajon("up", 185))

	return &direcciones
}

func (p *CajonEstacionamiento) GetUbicacionCajon() *floatgeom.Rect2 {
	return p.ubicacionDelCajon
}

func (p *CajonEstacionamiento) GetNumeroDeCajon() int {
	return p.numeroDeCajon
}

func (p *CajonEstacionamiento) GetDireccionesParaEstancionar() *[]CajonEstacionamientoDireccion {
	return p.direccionesParaEstacionar
}

func (p *CajonEstacionamiento) GetDirecccionesParaSalir() *[]CajonEstacionamientoDireccion {
	return p.direccionesParaSalir
}

func (p *CajonEstacionamiento) GetEstaCajonDisponible() bool {

	return p.disponible
}

func (p *CajonEstacionamiento) SetEstaCajonDisponible(estaDisponible bool) {
	p.disponible = estaDisponible
}
