package main

import (
	"math"

	gl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/leafo/weeklyloops/loops"
	"github.com/leafo/weeklyloops/loops/ply"
)

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "loop5"
	loop.Speed = 0.2

	obj, _ := ply.NewObjectFromFile("icosphere.ply")

	perspective := loops.NewPerspectiveMatrix(120, 1, 0.1, 100)

	verts := obj.PackF32("x", "y", "z", "nx", "ny", "nz")
	indexes := obj.PackIndexesB()

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.Enable3d()
		g.EnableWireframe()
		g.SetViewMat(perspective)

		m := loops.NewIdentityMat4().
			Rotate3d(float32(t*math.Pi*2), 0, 1, 0).
			Translate(0, 0, -3)

		g.SetMat(m)
		g.Draw3d(gl.TRIANGLES, verts, indexes)
	}

	loop.Run()
}
