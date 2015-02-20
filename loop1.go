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

				g.SetMat(s.Scale(float32(realScale), float32(realScale), 1).Rotate(realRad).Translate(float32(x), float32(y), 0))
				g.Draw(gl.TRIANGLE_STRIP, triangle)

				i += 1
				col += 1
			}

			row += 1
		}

	}

	loop.Run()
}
