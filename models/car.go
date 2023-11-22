package models

const (
	carWidth  = 25.0
	carHeight = 25.0
)

type AutoState int

const (
	StateEntering AutoState = iota
	StateParked
	StateExiting
)

type Car struct {
	PosX   float64
	PosY   float64
	Dir    float64
	Box    int
	Moving bool
	State  AutoState
}
