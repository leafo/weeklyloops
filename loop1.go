package main

import (
	"image/color"
	"math"

	gl "github.com/go-gl/gl"
	"github.com/leafo/weeklyloops/loops"
)

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "loop1"
	loop.Speed = 0.3

	a := loops.Vec2{0, 0.5}
	b := a.RotateAngle(360.0 / 3)
	c := b.RotateAngle(360.0 / 3)

	triangle := []float32{
		a[0], a[1],
		b[0], b[1],
		c[0], c[1],
	}

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.SetColor(color.RGBA{20, 20, 20, 255})
		g.SetMat(loops.NewIdentityMat4())
		g.DrawRect(-1, -1, 2, 2)

		t = loops.Smoothstep(t)
		s := loops.NewScaleMatrix(0.1, 0.1, 1)

		rad := float32(t * math.Pi * 2)

		i := 1
		row := 0
		col := 0

		g.SetColor(color.RGBA{200, 200, 200, 255})
		for y := -0.8; y <= 0.8; y += 0.1 {
			col = 0
			for x := -0.8; x <= 0.8; x += 0.1 {

				rad = -rad
				realRad := rad * float32(row/3+1)
				realScale := 1.0
				if i%2 == 0 {
					realScale = realScale / 1.5
				}

				m := s.Scale(float32(realScale), float32(realScale), 1).
					Rotate(realRad).
					Translate(float32(x), float32(y), 0)

				g.SetMat(m)
				g.Draw(gl.TRIANGLE_STRIP, triangle)

				i += 1
				col += 1
			}

			row += 1
		}

	}

	loop.Run()
}
