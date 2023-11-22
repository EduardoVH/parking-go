package main

import (
	"parking/models"
	"parking/views"

	pixel "github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

func main() {
	opengl.Run(func() {
		cfg := opengl.WindowConfig{
			Title:  "Parking",
			Bounds: pixel.R(0, 0, 1800, 300),
		}
		win, err := opengl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		e := models.NewParking(20)
		views.Run(win, e)
	})
}
