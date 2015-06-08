package main

import (
	"image/color"
	"log"
	"math"

	gl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/leafo/weeklyloops/loops"
	"github.com/leafo/weeklyloops/loops/ply"
)

var background = color.RGBA{28, 30, 22, 255}
var foreground = color.RGBA{215, 241, 220, 255}

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "loop4"
	loop.Speed = 0.2
	loop.Width = 400
	loop.Height = 400

	obj, _ := ply.NewObjectFromFile("cube.ply")

	if obj == nil {
		log.Fatal("failed to load model")
	}

	verts := obj.PackF32("x", "y", "z", "nx", "ny", "nz")
	indexes := obj.PackIndexesB()

	identity := loops.NewIdentityMat4()

	loop.Draw = func(t float64, g *loops.Graphics) {
		perspective := loops.NewPerspectiveMatrix(90+30*math.Sin(t*math.Pi*2), 1, 0.1, 100)

		g.Disable3d()
		// g.DisableWireframe()

		g.SetViewMat(identity)
		g.SetMat(identity)
		g.SetColor(background)
		g.DrawRect(-1, -1, 2, 2)

		g.Enable3d()
		// g.EnableWireframe()

		g.SetViewMat(perspective)
		g.SetColor(foreground)

		drawCube := func(t, scale, dist float64) {
			m := loops.NewIdentityMat4().
				Rotate3d(float32(t*math.Pi*2), 0, 1, 0).
				Translate(2, 2, float32(dist+-0.2*math.Sin(t*math.Pi*2*3)*scale)).
				Rotate(float32(t * math.Pi * 2))

			g.SetMat(m)

			g.Draw3d(gl.TRIANGLES, verts, indexes)

		}

		drawGroup := func(t, dist float64) {
			drawCube(t+math.Sin((t+0.75)*math.Pi*2)*0.05, 1, dist)
			drawCube(t+0.25+math.Sin((t+0.5)*math.Pi*2)*0.05, -1, dist)
			drawCube(t+0.5+math.Sin((t+0.25)*math.Pi*2)*0.05, 1, dist)
			drawCube(t+0.75+math.Sin(t*math.Pi*2)*0.05, -1, dist)
		}

		{

			g.SetColor(color.RGBA{242, 250, 250, 255})
			drawGroup(t+.7, -48*2)

			g.SetColor(color.RGBA{215, 241, 220, 255})
			drawGroup(t+.1, -48)

			g.SetColor(color.RGBA{209, 224, 183, 255})
			drawGroup(t+.5, -24)

			g.SetColor(color.RGBA{210, 187, 160, 255})
			drawGroup(t+.8, -12)

			g.SetColor(color.RGBA{198, 137, 155, 255})
			drawGroup(t+.3, -6)

			g.SetColor(color.RGBA{174, 111, 184, 255})

			drawGroup(t+0.6, -4)

			g.SetColor(color.RGBA{94, 92, 170, 255})
			drawGroup(t, -2)
		}
	}

	loop.Run()
}
