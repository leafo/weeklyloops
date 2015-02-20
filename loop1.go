package main

import (
	"math"

	gl "github.com/go-gl/gl"
	"github.com/leafo/weeklyloops/loops"
)

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "loop1"
	loop.Speed = 0.3

	v := loops.Vec2{0, 0.5}
	v.Print()

	x1 := v[0]
	y1 := v[1]

	v = v.RotateAngle(360.0 / 3)
	v.Print()

	x2 := v[0]
	y2 := v[1]

	v = v.RotateAngle(360.0 / 3)
	v.Print()

	x3 := v[0]
	y3 := v[1]

	triangle := []float32{x1, y1, x2, y2, x3, y3}

	loop.Draw = func(t float64, g *loops.Graphics) {
		s := loops.NewScaleMatrix(0.1, 0.1, 1)

		rad := float32(t * math.Pi * 2)

		i := 1
		row := 0
		col := 0

		for y := -0.8; y <= 0.8; y += 0.1 {
			col = 0
			for x := -0.8; x <= 0.8; x += 0.1 {
				rad = -rad
				realRad := rad * float32(row/3+1)

				g.SetMat(s.Rotate(realRad).Translate(float32(x), float32(y), 0))
				g.Draw(gl.TRIANGLE_STRIP, triangle)

				i += 1
				col += 1
			}

			row += 1
		}

	}

	loop.Run()
}
