package main

import (
	"log"
	"math"

	gl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/leafo/weeklyloops/loops"
	"github.com/leafo/weeklyloops/loops/ply"
)

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "loop4"
	loop.Speed = 0.2

	obj, _ := ply.NewObjectFromFile("cube.ply")

	if obj == nil {
		log.Fatal("failed to load model")
	}

	verts := obj.PackF32("x", "y", "z", "nx", "ny", "nz")
	indexes := obj.PackIndexesB()

	perspective := loops.NewPerspectiveMatrix(60, 1, 0.1, 100)

	loop.Draw = func(t float64, g *loops.Graphics) {
		loop.Enable3d()
		g.SetViewMat(perspective)

		drawCube := func(t float64, scale float64) {
			m := loops.NewIdentityMat4().
				Rotate3d(float32(t*math.Pi*2), 0, 1, 0).
				Translate(2, 2, float32(-3+-0.5*math.Sin(t*math.Pi*2*3)*scale)).
				Rotate(float32(t * math.Pi * 2))

			g.SetMat(m)

			g.Draw3d(gl.TRIANGLES, verts, indexes)

		}

		drawCube(t+math.Sin((t+0.75)*math.Pi*2)*0.05, 1)
		drawCube(t+0.25+math.Sin((t+0.5)*math.Pi*2)*0.05, -1)
		drawCube(t+0.5+math.Sin((t+0.25)*math.Pi*2)*0.05, 1)
		drawCube(t+0.75+math.Sin(t*math.Pi*2)*0.05, -1)
	}

	loop.Run()
}
