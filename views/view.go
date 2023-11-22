package views

import (
	"math/rand"
	"time"

	"parking/models"
	"parking/scenes"

	pixel "github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

const (
	speed       = 0.5
	carWidth    = 40.0
	carHeight   = 80.0
	spaceHeight = 10.0
	carDistance = 10.0
)

func CalcTime(car *models.Car) {
	time.Sleep(time.Duration(rand.Intn(15)+5) * time.Second)
	car.State = models.StateExiting
}

func GenerateCar(p *models.Parking) {
	rand.Seed(time.Now().UnixNano())

	for {
		car := &models.Car{PosX: -carWidth - carDistance, PosY: carHeight + spaceHeight, Dir: 1, State: models.StateEntering}
		pos := p.Enter(car)

		if pos != -1 {
			go CalcTime(car)
		}
		time.Sleep(time.Millisecond * 1500)
	}
}

func Run(win *opengl.Window, p *models.Parking) {

	go GenerateCar(p)

	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0))

		imd := imdraw.New(nil)
		scenes.RenderParkingLot(imd, p)

		p.Mut.Lock()
		for i, car := range p.Occupied {
			if car != nil {
				if car.State == models.StateEntering {
					car.PosX += speed * car.Dir
					car.PosY = carHeight + spaceHeight*2.2
					if car.PosX >= 90*float64(i) {
						car.State = models.StateParked
					}
					imd.Color = pixel.RGB(1, 1, 1)
					imd.Push(pixel.V(car.PosX, car.PosY))
					imd.Push(pixel.V(car.PosX+carWidth, car.PosY+carHeight))
					imd.Rectangle(0)
				} else if car.State == models.StateParked {
					car.PosX = 90*float64(i) + 25
					car.PosY = carHeight + spaceHeight - 80
					imd.Color = pixel.RGB(1, 1, 1)
					imd.Push(pixel.V(car.PosX, car.PosY))
					imd.Push(pixel.V(car.PosX+carWidth, car.PosY+carHeight))
					imd.Rectangle(0)
				} else if car.State == models.StateExiting {
					car.PosX -= speed * car.Dir
					car.PosY = carHeight + spaceHeight*2.2
					imd.Color = pixel.RGB(1, 0, 0)
					imd.Push(pixel.V(car.PosX, car.PosY))
					imd.Push(pixel.V(car.PosX+carWidth, car.PosY+carHeight))
					imd.Rectangle(0)

					if car.PosX <= -carWidth-carDistance {
						p.Occupied[i] = nil
						p.Boxes[car.Box] = false
					}
				}
			}
		}
		p.Mut.Unlock()

		imd.Draw(win)
		win.Update()
	}
}
