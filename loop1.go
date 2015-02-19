package main

import (
	"math"

	gl "github.com/go-gl/gl"
	"github.com/leafo/weeklyloops/loops"
)

func main() {
	loop := loops.NewLoopWindow()

	var elapsed float64

	loop.Load = func() {
	}

	loop.Update = func(dt float64) {
		elapsed += dt
	}

	loop.Draw = func(g *loops.Graphics) {
		t := loops.NewTranslateMatrix(float32(math.Sin(elapsed)), 0, 0)
		r := loops.NewRotate2DMatrix(float32(elapsed))

		g.SetMat(r.Mul(t))
		g.Draw(gl.TRIANGLE_STRIP, []float32{-0.5, -0.5, 0.5, -0.5, 0, 0.5})
		g.Draw(gl.TRIANGLE_STRIP, []float32{0, 0, 1, 1, 1, 0})
	}

	loop.Run()
}
