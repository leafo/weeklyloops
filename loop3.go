package main

import (
	"image/color"

	"github.com/leafo/weeklyloops/loops"
)

func main() {
	loop := loops.NewLoopWindow()

	loop.Title = "loop2"
	loop.Speed = 0.2

	circle := loops.NewCircle(1.0, 10)
	rect := loops.NewRectangle(0.5, 0.5)

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.SetColor(color.RGBA{200, 200, 200, 255})
		g.DrawShape(circle)
		g.DrawShape(rect)
	}

	loop.Run()
}
