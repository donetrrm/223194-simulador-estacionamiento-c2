package models

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

const (
	puntoDeEntrada = 185.00
	velocidad      = 10
)

type Carro struct {
	ubicacion floatgeom.Rect2
	entity    *entities.Entity
	mutex     sync.Mutex
}

func NewCarro(context *scene.Context) *Carro {
	ubicacion := floatgeom.NewRect2(445, -20, 465, 0)

	carRender, err := render.LoadSprite("assets/images/car.png")
	if err != nil {
		log.Fatal(err)
	}

	entity := entities.New(context, entities.WithRect(ubicacion), entities.WithRenderable(carRender), entities.WithDrawLayers([]int{1}))

	return &Carro{
		ubicacion: ubicacion,
		entity:    entity,
	}
}

func (c *Carro) Enqueue(admin *CarAdministrador) {

	for c.GetY() < 145 {
		if !c.carrosCollision("down", admin.GetCarros()) {
			c.ShiftY(1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}

}

func (c *Carro) PuertaEntrada(admin *CarAdministrador) {
	for c.GetY() < puntoDeEntrada {
		if !c.carrosCollision("down", admin.GetCarros()) {
			c.ShiftY(1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (c *Carro) PuertaSalida(admin *CarAdministrador) {
	for c.GetY() > 145 {
		if !c.carrosCollision("up", admin.GetCarros()) {
			c.ShiftY(-1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (c *Carro) Estacionamiento(cajon *CajonEstacionamiento, admin *CarAdministrador) {
	for i := 0; i < len(*cajon.GetDireccionesParaEstancionar()); i++ {
		direcciones := *cajon.GetDireccionesParaEstancionar()
		//fmt.Println("Punto de destino: " + fmt.Sprintf("%f", direcciones[i].Ubicacion))
		if direcciones[i].Direccion == "right" {
			for c.GetX() < direcciones[i].Ubicacion {
				if !c.carrosCollision("right", admin.GetCarros()) {
					c.ShiftX(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[i].Direccion == "down" {
			for c.GetY() < direcciones[i].Ubicacion {
				if !c.carrosCollision("down", admin.GetCarros()) {
					c.ShiftY(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[i].Direccion == "left" {
			for c.GetX() > direcciones[i].Ubicacion {
				if !c.carrosCollision("left", admin.GetCarros()) {
					c.ShiftX(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[i].Direccion == "up" {
			for c.GetY() > direcciones[i].Ubicacion {
				if !c.carrosCollision("up", admin.GetCarros()) {
					c.ShiftY(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		}
	}
}

func (c *Carro) DejarCajon(cajon *CajonEstacionamiento, admin *CarAdministrador) {
	for i := 0; i < len(*cajon.GetDirecccionesParaSalir()); i++ {
		direcciones := *cajon.GetDirecccionesParaSalir()
		if direcciones[i].Direccion == "left" {

			for c.GetX() > direcciones[i].Ubicacion {
				if !c.carrosCollision("left", admin.GetCarros()) {
					c.ShiftX(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[i].Direccion == "right" {
			for c.GetX() < direcciones[i].Ubicacion {
				if !c.carrosCollision("right", admin.GetCarros()) {
					c.ShiftX(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[i].Direccion == "up" {
			for c.GetY() > direcciones[i].Ubicacion {
				if !c.carrosCollision("up", admin.GetCarros()) {
					c.ShiftY(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[i].Direccion == "down" {
			for c.GetY() < direcciones[i].Ubicacion {
				if !c.carrosCollision("down", admin.GetCarros()) {
					c.ShiftY(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		}
	}
}

func (c *Carro) DejarLugar(admin *CarAdministrador) {
	spotX := c.GetX()
	for c.GetX() > spotX-30 {
		if !c.carrosCollision("left", admin.GetCarros()) {
			c.ShiftX(-1)
			time.Sleep(velocidad * time.Millisecond)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func (c *Carro) Irse(admin *CarAdministrador) {
	for c.GetY() > -20 {
		if !c.carrosCollision("up", admin.GetCarros()) {
			c.ShiftY(-1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (c *Carro) ShiftY(dy float64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entity.ShiftY(dy)
}

func (c *Carro) ShiftX(dx float64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entity.ShiftX(dx)
}

func (c *Carro) GetX() float64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.entity.X()
}

func (c *Carro) GetY() float64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.entity.Y()
}

func (c *Carro) Eliminar() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entity.Destroy()
}

func (c *Carro) carrosCollision(direccion string, carros []*Carro) bool {
	const distanciaMinima = 40.0

	for _, carro := range carros {
		switch direccion {
		case "left":
			if c.estaDentroDeLaDistancia(carro, distanciaMinima, true) && c.GetY() == carro.GetY() && c.GetX() > carro.GetX() {
				return true
			}
		case "right":
			if c.estaDentroDeLaDistancia(carro, distanciaMinima, true) && c.GetY() == carro.GetY() && c.GetX() < carro.GetX() {
				return true
			}
		case "up":
			if c.estaDentroDeLaDistancia(carro, distanciaMinima, false) && c.GetX() == carro.GetX() && c.GetY() > carro.GetY() {
				return true
			}
		case "down":
			if c.estaDentroDeLaDistancia(carro, distanciaMinima, false) && c.GetX() == carro.GetX() && c.GetY() < carro.GetY() {
				return true
			}
		}
	}
	return false
}

func (c *Carro) estaDentroDeLaDistancia(carro *Carro, distancia float64, horizontal bool) bool {
	if horizontal {
		return math.Abs(c.GetX()-carro.GetX()) < distancia
	}
	return math.Abs(c.GetY()-carro.GetY()) < distancia
}
