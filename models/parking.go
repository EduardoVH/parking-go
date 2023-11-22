package models

import (
	"sync"
)

type Parking struct {
	Spaces   int
	Mut      sync.Mutex
	Occupied []*Car
	Waiting  []*Car
	Boxes    []bool
}

func NewParking(capability int) *Parking {
	return &Parking{
		Spaces:   capability,
		Occupied: make([]*Car, capability),
		Boxes:    make([]bool, capability),
	}
}

func (p *Parking) Enter(car *Car) int {
	p.Mut.Lock()
	defer p.Mut.Unlock()
	for i, spot := range p.Occupied {
		if spot == nil {
			p.Occupied[i] = car
			car.Moving = true
			for j := 0; j < p.Spaces; j++ {
				if !p.Boxes[j] {
					p.Boxes[j] = true
					car.Box = j
					return i
				}
			}
			return -1
		}
	}
	return -1
}
