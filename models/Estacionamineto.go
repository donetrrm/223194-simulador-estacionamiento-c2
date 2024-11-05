package models

import (
	"sync"
)

type CarQueue struct {
	cars []Carro
}

func NewCarQueue() *CarQueue {
	return &CarQueue{
		cars: make([]Carro, 0),
	}
}

type Estacionamiento struct {
	cajones                     []*CajonEstacionamiento
	queueCars                   *CarQueue
	mutex                       sync.Mutex
	cajonesDisponiblesCondition *sync.Cond
}

func NewCajonEstacionamineto(cajones []*CajonEstacionamiento) *Estacionamiento {
	p := &Estacionamiento{
		cajones:   cajones,
		queueCars: NewCarQueue(),
	}
	p.cajonesDisponiblesCondition = sync.NewCond(&p.mutex)
	return p
}

func (p *Estacionamiento) GetCajones() []*CajonEstacionamiento {
	return p.cajones
}

func (p *Estacionamiento) GetCajonEstacionamientoDisponible() *CajonEstacionamiento {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for {
		for _, cajon := range p.cajones {
			if cajon.GetEstaCajonDisponible() {
				cajon.SetEstaCajonDisponible(false)
				return cajon
			}
		}
		p.cajonesDisponiblesCondition.Wait()
	}
}

func (p *Estacionamiento) LiberarCajon(cajon *CajonEstacionamiento) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	cajon.SetEstaCajonDisponible(true)
	p.cajonesDisponiblesCondition.Signal()
}

func (p *Estacionamiento) GetQueueCars() *CarQueue {
	return p.queueCars
}
