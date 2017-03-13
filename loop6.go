package main

import (
	"image/color"
	"math"

	gl "github.com/go-gl/gl/v3.3-core/gl"
	"github.com/leafo/weeklyloops/loops"
	"github.com/leafo/weeklyloops/loops/ply"
)

var background = color.RGBA{128, 140, 255, 255}

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "loop6"
	loop.Speed = 0.2

	perspective := loops.NewPerspectiveMatrix(200, 1, 0.1, 100)

	obj, _ := ply.NewObjectFromFile("icosphere.ply")

	verts := obj.PackF32("x", "y", "z", "nx", "ny", "nz")
	indexes := obj.PackIndexesB()

	pal := loops.SmoothPalette{
		[]color.RGBA{
			color.RGBA{114, 229, 254, 255},
			color.RGBA{255, 73, 217, 255},
			color.RGBA{252, 178, 118, 255},
		},
	}

	loop.Update = func(dt float64) {
	}

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.Disable3d()
		g.SetColor(background)
		g.SetViewMat(loops.NewIdentityMat4())
		g.SetMat(loops.NewIdentityMat4())
		g.DrawRect(-1, -1, 2, 2)

		g.Enable3d()
		g.SetViewMat(perspective)

		for k := 0; k < (len(indexes) / 3); k++ {
			s := t * 2

			if s > 1 {
				s = 2 - s
			}

			s = loops.Smoothstep(s)
			layers := 6
			layer := 0

			for f := 0.4; f <= 1.0; f += 0.1 {
				index := int32(indexes[k*3]) * 6
				nx := verts[index+3]
				ny := verts[index+4]
				nz := verts[index+5]

				m := loops.NewIdentityMat4().
					Scale(float32(f), float32(f), float32(f)).
					Translate(nx*2*(float32(s*f)+0.1), ny*2*(float32(s*f)+0.1), nz*2*(float32(s*f)+0.1)).
					Rotate3d(float32(s*math.Pi*2), nx, ny, nz).
					Rotate3d(float32(t*math.Pi*2), 0, 1, 0).
					Translate(0, 0, -10)

				g.SetMat(m)
				g.SetColor(pal.Color(float64(layer) / float64(layers)))
				g.Draw3d(gl.TRIANGLES, verts, indexes[k*3:(k+1)*3])
				layer++
			}
		}

	}

	loop.Run()
}
