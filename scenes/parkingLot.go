package scenes

import (
	"parking/models"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

const (
	windowWidth   = 1000.0
	spaces        = 10
	spacesSize    = windowWidth / float64(spaces+1.1)
	spaceHeight   = 100.0
	boxSpaceWidth = 70
)

func RenderParkingLot(imd *imdraw.IMDraw, p *models.Parking) {
	for i := 0; i < p.Spaces; i++ {
		if p.Boxes[i] {
			imd.Color = pixel.RGB(0, 0, 1)
		} else {
			imd.Color = pixel.RGB(1, 1, 1)
		}
		xStart := spacesSize*float64(i) + spacesSize/2
		centerY := spaceHeight / 2

		imd.Push(pixel.V(xStart, centerY+spaceHeight/2))
		imd.Push(pixel.V(xStart, centerY-spaceHeight/2))
		imd.Rectangle(boxSpaceWidth)

		imd.Push(pixel.V(xStart, centerY+spaceHeight/2))
		imd.Push(pixel.V(xStart, centerY-spaceHeight/2))
		imd.Rectangle(0)
	}
}
