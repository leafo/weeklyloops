package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/leafo/weeklyloops/loops"
)

func main() {
	loop := loops.NewLoopWindow()

	loop.Title = "loop2"
	loop.Speed = 0.3

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.SetMat(loops.NewIdentityMat4())
		g.DrawColored(gl.TRIANGLE_STRIP, []float32{
			-0.5, -0.5, 1, 0, 0, 1,
			0.5, -0.5, 0, 1, 0, 1,
			0, 0, 0, 0, 1, 1,
		})
	}

	loop.Run()
}
