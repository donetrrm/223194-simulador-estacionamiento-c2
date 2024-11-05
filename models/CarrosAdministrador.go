package models

import "sync"

type CarAdministrador struct {
	Carros []*Carro
	Mutex  sync.Mutex
}

func NewCarAdministrador() *CarAdministrador {
	return &CarAdministrador{
		Carros: make([]*Carro, 0),
	}
}

func (ca *CarAdministrador) AgregarCarro(carro *Carro) {
	ca.Mutex.Lock()
	defer ca.Mutex.Unlock()
	ca.Carros = append(ca.Carros, carro)
}

func (ca *CarAdministrador) EliminarCarroScene(carro *Carro) {
	ca.Mutex.Lock()
	defer ca.Mutex.Unlock()
	for i, c := range ca.Carros {
		if c == carro {
			ca.Carros = append(ca.Carros[:i], ca.Carros[i+1:]...)
			break
		}
	}
}

func (ca *CarAdministrador) GetCarros() []*Carro {
	ca.Mutex.Lock()
	defer ca.Mutex.Unlock()
	carrosCopia := make([]*Carro, len(ca.Carros))
	copy(carrosCopia, ca.Carros)
	return carrosCopia
}
