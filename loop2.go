package main

import "github.com/leafo/weeklyloops/loops"

func main() {
	loop := loops.NewLoopWindow()

	loop.Title = "loop2"
	loop.Speed = 0.3

	rect := loops.NewRectangle(1, 0.5)

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.SetMat(loops.NewIdentityMat4())
		g.DrawShape(rect)
	}

	loop.Run()
}
